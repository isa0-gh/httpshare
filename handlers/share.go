package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/isa0/httpshare/models"
	"gitlab.com/isa0/httpshare/utils"
)

// CreateShareLink creates a shareable link
func CreateShareLink(c echo.Context) error {
	filePath := c.FormValue("path")
	expiresIn := c.FormValue("expiresIn") // in hours

	if filePath == "" {
		return c.JSON(400, map[string]string{"error": "Path is required"})
	}

	// Generate random ID
	b := make([]byte, 16)
	rand.Read(b)
	id := hex.EncodeToString(b)

	link := &models.ShareLink{
		ID:        id,
		FilePath:  filePath,
		Downloads: 0,
		CreatedAt: time.Now(),
	}

	// Set expiration if provided
	if expiresIn != "" {
		var hours int
		if _, err := time.ParseDuration(expiresIn + "h"); err == nil {
			expiresAt := time.Now().Add(time.Duration(hours) * time.Hour)
			link.ExpiresAt = &expiresAt
		}
	}

	utils.AddShareLink(link)

	return c.JSON(200, map[string]interface{}{
		"id":   id,
		"url":  "/share/" + id,
		"link": link,
	})
}

// GetShareLink retrieves a file via share link
func GetShareLink(c echo.Context) error {
	id := c.Param("id")

	link := utils.GetShareLink(id)
	if link == nil {
		return c.JSON(404, map[string]string{"error": "Share link not found"})
	}

	// Check expiration
	if link.ExpiresAt != nil && time.Now().After(*link.ExpiresAt) {
		utils.DeleteShareLink(id)
		return c.JSON(410, map[string]string{"error": "Share link expired"})
	}

	// Check max downloads
	if link.MaxDownloads > 0 && link.Downloads >= link.MaxDownloads {
		utils.DeleteShareLink(id)
		return c.JSON(410, map[string]string{"error": "Download limit reached"})
	}

	// Increment download counter
	link.Downloads++
	utils.AddShareLink(link)

	fullPath := utils.UrlToFilePath(cfg.Directory, link.FilePath)
	return c.File(fullPath)
}

// ListShareLinks lists all share links
func ListShareLinks(c echo.Context) error {
	links := utils.GetAllShareLinks()
	return c.JSON(200, links)
}

// DeleteShareLink deletes a share link
func DeleteShareLinkHandler(c echo.Context) error {
	id := c.Param("id")
	utils.DeleteShareLink(id)
	return c.JSON(200, map[string]string{"message": "Share link deleted"})
}
