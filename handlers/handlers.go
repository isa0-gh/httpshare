package handlers

import (
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
	"gitlab.com/isa0/httpshare/config"
	"gitlab.com/isa0/httpshare/template"
	"gitlab.com/isa0/httpshare/utils"
)

var cfg *config.Config = config.Cfg

func Logger() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			err := next(c)
			req := c.Request()
			res := c.Response()

			// Custom log format
			if req.RequestURI != "/tailwind.js" && req.RequestURI != "/favicon.ico" {
				now := time.Now()
				logMsg := fmt.Sprintf("%s [%s] %s %s | Status-Code: %d | User-Agent: %s",
					now.Format("2006/01/02 15:04:05"),
					req.Method,
					req.RequestURI,
					c.RealIP(),
					res.Status,
					req.UserAgent(),
				)
				if cfg.LogFile != "" {
					file, err := os.OpenFile(cfg.LogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
					if err != nil {
						log.Info("Log writing error:", err.Error())
					}

					defer file.Close()
					_, err = file.WriteString(logMsg + "\n")
					if err != nil {
						log.Info("Log writing error:", err.Error())
					}
				}
				fmt.Println(logMsg)
			}

			return err
		}
	}
}

// UploadFile handles file upload
func UploadFile(c echo.Context) error {
	dirPath := utils.UrlToFilePath(cfg.Directory, c.FormValue("path"))

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
	if err := utils.DeleteFile(utils.UrlToFilePath(cfg.Directory, path)); err != nil {
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

	if err := utils.RenameFile(utils.UrlToFilePath(cfg.Directory, oldPath), newName); err != nil {
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

	if err := utils.CreateDirectory(utils.UrlToFilePath(cfg.Directory, path)); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Directory created successfully"})
}

// SearchFiles handles file search
func SearchFiles(c echo.Context) error {
	query := c.QueryParam("q")
	basePath := utils.UrlToFilePath(cfg.Directory, c.QueryParam("path"))

	results, err := utils.SearchFiles(basePath, query)
	if err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, results)
}

// BrowseFiles handles file browsing and rendering
func BrowseFiles(c echo.Context) error {
	decoded, _ := url.QueryUnescape(c.Param("*"))
	path := utils.UrlToFilePath(cfg.Directory, decoded)
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
