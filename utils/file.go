package utils

import (
	"os"

	"github.com/isa0-gh/httpshare/models"
)

func GetFiles(path string) (models.DirectoryEntries, error) {
	var directoryEntries models.DirectoryEntries
	entries, err := os.ReadDir(path)
	if err != nil {
		return directoryEntries, err
	}
	for _, entry := range entries {
		if entry.IsDir() {
			directoryEntries.Directories = append(directoryEntries.Directories, entry.Name())
		} else {
			directoryEntries.Files = append(directoryEntries.Files, entry.Name())
		}
	}

	return directoryEntries, nil
}
