package utils

import (
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/isa0-gh/httpshare/models"
)

func isImage(filename string) bool {
	ext := path.Ext(filename)
	for _, extType := range []string{".png", ".bmp", ".webp", ".gif", ".jpg", ".jpeg", ".svg", ".tiff"} {
		if ext == extType {
			return true
		}
	}
	return false
}

func GetFiles(path string) (models.DirectoryEntries, error) {
	var directoryEntries models.DirectoryEntries
	directoryEntries.Path = path
	entries, err := os.ReadDir(path)
	if err != nil {
		return directoryEntries, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			directoryEntries.Directories = append(directoryEntries.Directories, entry.Name())
		} else {
			if isImage(entry.Name()) {
				directoryEntries.Images = append(directoryEntries.Images, entry.Name())
			} else {
				directoryEntries.Files = append(directoryEntries.Files, entry.Name())
			}
		}
	}
	return directoryEntries, nil
}

func UrlToFilePath(basePath string, url string) string {
	parts := strings.Split(url, "/")
	return filepath.Join(basePath, filepath.Join(parts...))
}
