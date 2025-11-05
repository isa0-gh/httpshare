package handlers

import (
	"github.com/isa0-gh/httpshare/template"
	"github.com/labstack/echo/v4"
)

// ServeTailwind serves the embedded Tailwind CSS
func ServeTailwind(c echo.Context) error {
	return c.String(200, template.GetTailwind())
}
