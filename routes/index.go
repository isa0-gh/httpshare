package routes

import (
	_ "embed"
	"net/url"

	"github.com/isa0-gh/httpshare/args"
	"github.com/isa0-gh/httpshare/template"
	"github.com/isa0-gh/httpshare/utils"
	"github.com/labstack/echo/v4"
)

//go:embed tailwind.js
var tailwind string

func Index(c echo.Context) error {
	decoded, _ := url.QueryUnescape(c.Param("*"))
	path := utils.UrlToFilePath(*args.Dir, decoded)
	if path == "" {
		path = "."
	}
	if c.QueryParam("download") == "true" {
		return c.File(path)
	}
	data, err := utils.GetFiles(path)
	if err != nil {
		return c.String(500, err.Error())
	}
	output, err := template.Render(data)
	if err != nil {
		return c.String(500, err.Error())
	}

	return c.HTML(200, output)
}

func Tailwind(c echo.Context) error {
	return c.String(200, tailwind)
}
