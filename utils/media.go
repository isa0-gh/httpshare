package utils

import (
	"path/filepath"
	"strings"
)

var imageExtensions = []string{".jpg", ".jpeg", ".png", ".gif", ".bmp", ".webp", ".svg"}
var videoExtensions = []string{".mp4", ".webm", ".avi", ".mov", ".mkv", ".flv"}
var audioExtensions = []string{".mp3", ".wav", ".ogg", ".flac", ".m4a", ".aac"}
var previewExtensions = []string{".txt", ".md", ".log", ".json", ".xml", ".csv", ".pdf"}
var officeExtensions = []string{".docx", ".xlsx", ".pptx", ".doc", ".xls", ".ppt"}

func IsImage(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, e := range imageExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

func IsVideo(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, e := range videoExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

func IsAudio(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, e := range audioExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

func CanPreview(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, e := range append(append(append(previewExtensions, videoExtensions...), audioExtensions...), imageExtensions...) {
		if ext == e {
			return true
		}
	}
	for _, e := range officeExtensions {
		if ext == e {
			return true
		}
	}
	return false
}

func IsOfficeDoc(filename string) bool {
	ext := strings.ToLower(filepath.Ext(filename))
	for _, e := range officeExtensions {
		if ext == e {
			return true
		}
	}
	return false
}
