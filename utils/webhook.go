package utils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"time"
)

type WebhookPayload struct {
	Event     string                 `json:"event"`
	FilePath  string                 `json:"file_path"`
	Timestamp time.Time              `json:"timestamp"`
	Details   map[string]interface{} `json:"details,omitempty"`
}

func TriggerWebhooks(event, filePath string, details map[string]interface{}) {
	webhooks := GetAllWebhooks()

	payload := WebhookPayload{
		Event:     event,
		FilePath:  filePath,
		Timestamp: time.Now(),
		Details:   details,
	}

	for _, webhook := range webhooks {
		if !webhook.Active {
			continue
		}

		// Check if webhook is subscribed to this event
		subscribed := false
		for _, e := range webhook.Events {
			if e == event || e == "*" {
				subscribed = true
				break
			}
		}

		if !subscribed {
			continue
		}

		// Send webhook in goroutine
		go func(url string) {
			data, _ := json.Marshal(payload)
			http.Post(url, "application/json", bytes.NewBuffer(data))
		}(webhook.URL)
	}
}
