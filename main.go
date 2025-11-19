package main

import (
	"fmt"
	"log"

	"gitlab.com/isa0/httpshare/config"

	"github.com/labstack/echo/v4"
	"gitlab.com/isa0/httpshare/routes"
)

func main() {
	// Load configuration
	var cfg *config.Config = config.Cfg

	// Create Echo instance
	e := echo.New()
	e.HideBanner = true
	e.HidePort = true

	// Setup routes
	routes.Setup(e)

	// Start server
	addr := fmt.Sprintf("%s:%d", cfg.Host, cfg.Port)
	log.Printf("🚀 httpshare starting on http://%s", addr)
	log.Printf("📁 Serving files from: %s", cfg.Directory)
	log.Printf("💡 Press Ctrl+C to stop")

	if err := e.Start(fmt.Sprintf(":%d", cfg.Port)); err != nil {
		log.Fatalf("❌ Failed to start server: %v", err)
	}

}
