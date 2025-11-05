package utils

import (
	"fmt"
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
		info, err := entry.Info()
		if err != nil {
			continue
		}
		
		fileEntry := models.FileEntry{
			Name:        entry.Name(),
			IsDir:       entry.IsDir(),
			IsImage:     isImage(entry.Name()),
			Size:        info.Size(),
			ModTime:     info.ModTime(),
			Permissions: info.Mode().String(),
		}
		directoryEntries.Entries = append(directoryEntries.Entries, fileEntry)
	}
	return directoryEntries, nil
}

func UrlToFilePath(basePath string, url string) string {
	parts := strings.Split(url, "/")
	return filepath.Join(basePath, filepath.Join(parts...))
}

// SearchFiles searches for files/folders matching the query
func SearchFiles(basePath, query string) ([]models.FileEntry, error) {
	var results []models.FileEntry
	query = strings.ToLower(query)
	
	err := filepath.Walk(basePath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if strings.Contains(strings.ToLower(info.Name()), query) {
			relPath, _ := filepath.Rel(basePath, path)
			results = append(results, models.FileEntry{
				Name:        relPath,
				IsDir:       info.IsDir(),
				IsImage:     isImage(info.Name()),
				Size:        info.Size(),
				ModTime:     info.ModTime(),
				Permissions: info.Mode().String(),
			})
		}
		return nil
	})
	return results, err
}

// SortEntries sorts file entries by field and direction
func SortEntries(entries []models.FileEntry, sortBy, order string) []models.FileEntry {
	if len(entries) == 0 {
		return entries
	}
	
	// Simple bubble sort
	for i := 0; i < len(entries)-1; i++ {
		for j := 0; j < len(entries)-i-1; j++ {
			swap := false
			switch sortBy {
			case "name":
				if order == "asc" {
					swap = entries[j].Name > entries[j+1].Name
				} else {
					swap = entries[j].Name < entries[j+1].Name
				}
			case "size":
				if order == "asc" {
					swap = entries[j].Size > entries[j+1].Size
				} else {
					swap = entries[j].Size < entries[j+1].Size
				}
			case "date":
				if order == "asc" {
					swap = entries[j].ModTime.After(entries[j+1].ModTime)
				} else {
					swap = entries[j].ModTime.Before(entries[j+1].ModTime)
				}
			}
			if swap {
				entries[j], entries[j+1] = entries[j+1], entries[j]
			}
		}
	}
	return entries
}

// DeleteFile deletes a file or directory
func DeleteFile(path string) error {
	return os.RemoveAll(path)
}

// RenameFile renames a file or directory
func RenameFile(oldPath, newName string) error {
	dir := filepath.Dir(oldPath)
	newPath := filepath.Join(dir, newName)
	return os.Rename(oldPath, newPath)
}

// CreateDirectory creates a new directory
func CreateDirectory(path string) error {
	return os.MkdirAll(path, 0755)
}

// FormatSize formats bytes into human-readable size
func FormatSize(bytes int64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	units := []string{"B", "KB", "MB", "GB", "TB"}
	return fmt.Sprintf("%.1f %s", float64(bytes)/float64(div), units[exp])
}
