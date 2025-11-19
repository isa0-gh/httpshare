package handlers

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/labstack/echo/v4"
	"gitlab.com/isa0/httpshare/models"
	"gitlab.com/isa0/httpshare/utils"
)

// CreateWebhook creates a new webhook
func CreateWebhook(c echo.Context) error {
	var webhook models.Webhook
	if err := c.Bind(&webhook); err != nil {
		return c.JSON(400, map[string]string{"error": "Invalid request"})
	}

	if webhook.URL == "" {
		return c.JSON(400, map[string]string{"error": "URL is required"})
	}

	// Generate random ID
	b := make([]byte, 8)
	rand.Read(b)
	webhook.ID = hex.EncodeToString(b)
	webhook.CreatedAt = time.Now()
	webhook.Active = true

	utils.AddWebhook(&webhook)

	return c.JSON(200, webhook)
}

// ListWebhooks lists all webhooks
func ListWebhooks(c echo.Context) error {
	webhooks := utils.GetAllWebhooks()
	return c.JSON(200, webhooks)
}

// DeleteWebhook deletes a webhook
func DeleteWebhookHandler(c echo.Context) error {
	id := c.Param("id")
	utils.DeleteWebhook(id)
	return c.JSON(200, map[string]string{"message": "Webhook deleted"})
}

// ToggleWebhook toggles webhook active status
func ToggleWebhook(c echo.Context) error {
	id := c.Param("id")
	webhook := utils.GetWebhook(id)

	if webhook == nil {
		return c.JSON(404, map[string]string{"error": "Webhook not found"})
	}

	webhook.Active = !webhook.Active
	utils.AddWebhook(webhook)

	return c.JSON(200, webhook)
}
