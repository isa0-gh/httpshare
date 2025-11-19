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
		// File operations
		api.POST("/upload", handlers.UploadFile)
		api.DELETE("/delete", handlers.DeleteFile)
		api.POST("/rename", handlers.RenameFile)
		api.POST("/mkdir", handlers.CreateDirectory)
		api.GET("/search", handlers.SearchFiles)
		api.GET("/download-zip", handlers.DownloadZip)
		api.POST("/bulk-download-zip", handlers.BulkDownloadZip)
		api.POST("/copy", handlers.CopyFileHandler)
		api.POST("/move", handlers.MoveFileHandler)

		// Share links
		api.POST("/share", handlers.CreateShareLink)
		api.GET("/shares", handlers.ListShareLinks)
		api.DELETE("/share/:id", handlers.DeleteShareLinkHandler)

		// Comments
		api.POST("/comment", handlers.AddComment)
		api.GET("/comments", handlers.GetComments)

		// Webhooks
		api.POST("/webhook", handlers.CreateWebhook)
		api.GET("/webhooks", handlers.ListWebhooks)
		api.DELETE("/webhook/:id", handlers.DeleteWebhookHandler)
		api.POST("/webhook/:id/toggle", handlers.ToggleWebhook)
	}

	// Share link access
	e.GET("/share/:id", handlers.GetShareLink)

	// File browsing (catch-all route)
	e.GET("/*", handlers.BrowseFiles)
}
