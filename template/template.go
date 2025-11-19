package template

import (
	"bytes"
	_ "embed"
	"html/template"
	"path"
	"strings"

	"gitlab.com/isa0/httpshare/models"
	"gitlab.com/isa0/httpshare/utils"
)

//go:embed index.html
var indexHTML string

//go:embed tailwind.js
var tailwindJS string

func canPreview(filename string) bool {
	ext := strings.ToLower(path.Ext(filename))
	previewableExts := []string{".txt", ".md", ".log", ".json", ".xml", ".csv", ".pdf", ".mp4", ".webm", ".mp3", ".wav"}
	for _, e := range previewableExts {
		if ext == e {
			return true
		}
	}
	return false
}

func Render(data models.DirectoryEntries) (string, error) {
	funcMap := template.FuncMap{
		"formatSize": utils.FormatSize,
		"canPreview": canPreview,
	}
	tmpl := template.Must(template.New("index").Funcs(funcMap).Parse(indexHTML))
	var buf bytes.Buffer

	if err := tmpl.Execute(&buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

// GetTailwind returns the embedded Tailwind CSS
func GetTailwind() string {
	return tailwindJS
}
