package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/isa0/httpshare/models"
	"gitlab.com/isa0/httpshare/utils"
)

// AddComment adds a comment to a file
func AddComment(c echo.Context) error {
	filePath := c.FormValue("path")
	author := c.FormValue("author")
	content := c.FormValue("content")

	if filePath == "" || content == "" {
		return c.JSON(400, map[string]string{"error": "Missing parameters"})
	}

	if author == "" {
		author = "Anonymous"
	}

	// Generate random ID
	b := make([]byte, 8)
	rand.Read(b)
	id := hex.EncodeToString(b)

	comment := models.FileComment{
		ID:        id,
		FilePath:  filePath,
		Author:    author,
		Content:   content,
		CreatedAt: time.Now(),
	}

	utils.AddComment(comment)

	return c.JSON(200, comment)
}

// GetComments retrieves comments for a file
func GetComments(c echo.Context) error {
	filePath := c.QueryParam("path")
	if filePath == "" {
		return c.JSON(400, map[string]string{"error": "Path is required"})
	}

	comments := utils.GetComments(filePath)
	return c.JSON(200, comments)
}
