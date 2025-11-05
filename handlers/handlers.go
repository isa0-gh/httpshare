package handlers

import (
	"io"
	"net/url"
	"os"
	"path/filepath"

	"github.com/isa0-gh/httpshare/template"
	"github.com/isa0-gh/httpshare/utils"
	"github.com/labstack/echo/v4"
)

// UploadFile handles file upload
func UploadFile(c echo.Context) error {
	dirPath := c.FormValue("path")
	if dirPath == "" {
		dirPath = "."
	}

	file, err := c.FormFile("file")
	if err != nil {
		return c.JSON(400, map[string]string{"error": "No file uploaded"})
	}

	src, err := file.Open()
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Cannot open file"})
	}
	defer src.Close()

	dstPath := filepath.Join(dirPath, file.Filename)
	dst, err := os.Create(dstPath)
	if err != nil {
		return c.JSON(500, map[string]string{"error": "Cannot create file"})
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return c.JSON(500, map[string]string{"error": "Cannot save file"})
	}

	return c.JSON(200, map[string]string{"message": "File uploaded successfully"})
}

// DeleteFile handles file/folder deletion
func DeleteFile(c echo.Context) error {
	path := c.QueryParam("path")
	if path == "" {
		return c.JSON(400, map[string]string{"error": "Path is required"})
	}

	if err := utils.DeleteFile(path); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Deleted successfully"})
}

// RenameFile handles file/folder renaming
func RenameFile(c echo.Context) error {
	oldPath := c.FormValue("oldPath")
	newName := c.FormValue("newName")

	if oldPath == "" || newName == "" {
		return c.JSON(400, map[string]string{"error": "Missing parameters"})
	}

	if err := utils.RenameFile(oldPath, newName); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Renamed successfully"})
}

// CreateDirectory handles directory creation
func CreateDirectory(c echo.Context) error {
	path := c.FormValue("path")
	if path == "" {
		return c.JSON(400, map[string]string{"error": "Path is required"})
	}

	if err := utils.CreateDirectory(path); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Directory created successfully"})
}

// SearchFiles handles file search
func SearchFiles(c echo.Context) error {
	query := c.QueryParam("q")
	basePath := c.QueryParam("path")
	if basePath == "" {
		basePath = "."
	}

	results, err := utils.SearchFiles(basePath, query)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, results)
}

// BrowseFiles handles file browsing and rendering
func BrowseFiles(c echo.Context) error {
	decoded, _ := url.QueryUnescape(c.Param("*"))
	path := utils.UrlToFilePath(decoded)
	if path == "" {
		path = "."
	}

	// Download mode
	if c.QueryParam("download") == "true" {
		return c.File(path)
	}

	// Get files
	data, err := utils.GetFiles(path)
	if err != nil {
		return c.String(500, err.Error())
	}

	// Apply sorting
	sortBy := c.QueryParam("sort")
	order := c.QueryParam("order")
	if sortBy != "" {
		if order == "" {
			order = "asc"
		}
		data.Entries = utils.SortEntries(data.Entries, sortBy, order)
	}

	// Render template
	output, err := template.Render(data)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.HTML(200, output)
}
