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

	// Trigger webhooks
	utils.TriggerWebhooks("upload", dstPath, map[string]interface{}{
		"filename": file.Filename,
		"size":     file.Size,
	})

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

	// Trigger webhooks
	utils.TriggerWebhooks("delete", path, nil)

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

	// Trigger webhooks
	utils.TriggerWebhooks("rename", oldPath, map[string]interface{}{"newName": newName})

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

// DownloadZip handles zip download of files/folders
func DownloadZip(c echo.Context) error {
	path := c.QueryParam("path")
	if path == "" {
		return c.JSON(400, map[string]string{"error": "Path is required"})
	}

	fullPath := utils.UrlToFilePath(cfg.Directory, path)
	info, err := os.Stat(fullPath)
	if err != nil {
		return c.JSON(404, map[string]string{"error": "File not found"})
	}

	// Create temp zip file
	tempZip := filepath.Join(os.TempDir(), fmt.Sprintf("%s_%d.zip", filepath.Base(fullPath), time.Now().Unix()))
	defer os.Remove(tempZip)

	if info.IsDir() {
		if err := utils.ZipDirectory(fullPath, tempZip); err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to create zip"})
		}
	} else {
		if err := utils.ZipFile(fullPath, tempZip); err != nil {
			return c.JSON(500, map[string]string{"error": "Failed to create zip"})
		}
	}

	return c.Attachment(tempZip, filepath.Base(fullPath)+".zip")
}

// BulkDownloadZip handles zip download of multiple files/folders
func BulkDownloadZip(c echo.Context) error {
	var paths []string
	if err := c.Bind(&paths); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	if len(paths) == 0 {
		return c.JSON(400, map[string]string{"error": "No files selected"})
	}

	// Create temp zip file
	tempZip := filepath.Join(os.TempDir(), fmt.Sprintf("bulk_%d.zip", time.Now().Unix()))
	defer os.Remove(tempZip)

	// Convert paths to full paths
	fullPaths := make([]string, len(paths))
	for i, p := range paths {
		fullPaths[i] = utils.UrlToFilePath(cfg.Directory, p)
	}

	if err := utils.ZipMultipleFiles(fullPaths, tempZip); err != nil {
		return c.JSON(500, map[string]string{"error": "Failed to create zip: " + err.Error()})
	}

	return c.Attachment(tempZip, fmt.Sprintf("files_%d.zip", time.Now().Unix()))
}

// CopyFile handles file/folder copying
func CopyFileHandler(c echo.Context) error {
	srcPath := c.FormValue("srcPath")
	dstPath := c.FormValue("dstPath")

	if srcPath == "" || dstPath == "" {
		return c.JSON(400, map[string]string{"error": "Missing parameters"})
	}

	fullSrc := utils.UrlToFilePath(cfg.Directory, srcPath)
	fullDst := utils.UrlToFilePath(cfg.Directory, dstPath)

	if err := utils.CopyFile(fullSrc, fullDst); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	return c.JSON(200, map[string]string{"message": "Copied successfully"})
}

// MoveFile handles file/folder moving
func MoveFileHandler(c echo.Context) error {
	srcPath := c.FormValue("srcPath")
	dstPath := c.FormValue("dstPath")

	if srcPath == "" || dstPath == "" {
		return c.JSON(400, map[string]string{"error": "Missing parameters"})
	}

	fullSrc := utils.UrlToFilePath(cfg.Directory, srcPath)
	fullDst := utils.UrlToFilePath(cfg.Directory, dstPath)

	if err := utils.MoveFile(fullSrc, fullDst); err != nil {
		return c.JSON(500, map[string]string{"error": err.Error()})
	}

	utils.TriggerWebhooks("move", srcPath, map[string]interface{}{"destination": dstPath})

	return c.JSON(200, map[string]string{"message": "Moved successfully"})
}
