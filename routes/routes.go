package routes

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gitlab.com/isa0/httpshare/handlers"
)

// Setup configures all application routes
func Setup(e *echo.Echo) {
	// Middleware
	e.Use(handlers.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Static assets
	e.GET("/tailwind.js", handlers.ServeTailwind)

	// API routes
	api := e.Group("/api")
	{
		api.POST("/upload", handlers.UploadFile)
		api.DELETE("/delete", handlers.DeleteFile)
		api.POST("/rename", handlers.RenameFile)
		api.POST("/mkdir", handlers.CreateDirectory)
		api.GET("/search", handlers.SearchFiles)
	}

	// File browsing (catch-all route)
	e.GET("/*", handlers.BrowseFiles)
}
