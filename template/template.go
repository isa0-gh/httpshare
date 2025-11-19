package template

import (
	"bytes"
	_ "embed"
	"html/template"

	"gitlab.com/isa0/httpshare/models"
	"gitlab.com/isa0/httpshare/utils"
)

//go:embed index.html
var indexHTML string

//go:embed tailwind.js
var tailwindJS string

func canPreview(filename string) bool {
	return utils.CanPreview(filename)
}

func Render(data models.DirectoryEntries) (string, error) {
	funcMap := template.FuncMap{
		"formatSize": utils.FormatSize,
		"canPreview": canPreview,
		"isVideo":    utils.IsVideo,
		"isAudio":    utils.IsAudio,
		"isOffice":   utils.IsOfficeDoc,
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
