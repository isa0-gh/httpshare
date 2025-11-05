package main

import (
	"fmt"
	"log"

	"github.com/isa0-gh/httpshare/config"
	"github.com/isa0-gh/httpshare/routes"
	"github.com/labstack/echo/v4"
)

func main() {
	// Load configuration
	cfg := config.Load()

	// Create Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Setup routes
	routes.Setup(e)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("🚀 httpshare starting on http://%s", addr)
	log.Printf("📁 Serving files from: %s", ".")
	log.Printf("💡 Press Ctrl+C to stop")
	
	if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}
}
