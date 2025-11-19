package models

import "time"

type FileEntry struct {
	Name        string
	IsDir       bool
	IsImage     bool
	IsVideo     bool
	IsAudio     bool
	Size        int64
	ModTime     time.Time
	Permissions string
}

type DirectoryEntries struct {
	Path    string
	Entries []FileEntry
}

// ShareLink represents a shareable link
type ShareLink struct {
	ID           string     `json:"id"`
	FilePath     string     `json:"file_path"`
	ExpiresAt    *time.Time `json:"expires_at,omitempty"`
	Downloads    int        `json:"downloads"`
	MaxDownloads int        `json:"max_downloads,omitempty"`
	CreatedAt    time.Time  `json:"created_at"`
}

// FileComment represents a comment on a file
type FileComment struct {
	ID        string    `json:"id"`
	FilePath  string    `json:"file_path"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// Webhook represents a webhook configuration
type Webhook struct {
	ID        string    `json:"id"`
	URL       string    `json:"url"`
	Events    []string  `json:"events"` // upload, delete, rename
	Active    bool      `json:"active"`
	CreatedAt time.Time `json:"created_at"`
}

// Theme settings
type Theme struct {
	Mode         string            `json:"mode"` // light, dark, auto
	CustomColors map[string]string `json:"custom_colors,omitempty"`
}
