package main

import (
	_ "embed"
	"fmt"

	"github.com/isa0-gh/httpshare/args"
	"github.com/isa0-gh/httpshare/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	args.Init()

	e.HideBanner = true

	e.GET("/tailwind.js", routes.Tailwind)
	e.GET("/*", routes.Index)

	fmt.Printf("Listening on http://localhost:%d/\nDirectory:%s\n", *args.Port, *args.Dir)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", *args.Port)))
}
