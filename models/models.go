package models

import "time"

type FileEntry struct {
	Name        string
	IsDir       bool
	IsImage     bool
	Size        int64
	ModTime     time.Time
	Permissions string
}

type DirectoryEntries struct {
	Path    string
	Entries []FileEntry
}
