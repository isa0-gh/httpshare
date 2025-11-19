package utils

import (
	"encoding/json"
	"os"
	"path/filepath"
	"sync"

	"gitlab.com/isa0/httpshare/models"
)

var (
	dataDir      = ".httpshare_data"
	shareLinks   = make(map[string]*models.ShareLink)
	comments     = make(map[string][]models.FileComment)
	webhooks     = make(map[string]*models.Webhook)
	storageMutex sync.RWMutex
)

func init() {
	os.MkdirAll(dataDir, 0755)
	loadData()
}

func loadData() {
	// Load share links
	if data, err := os.ReadFile(filepath.Join(dataDir, "shares.json")); err == nil {
		json.Unmarshal(data, &shareLinks)
	}
	// Load comments
	if data, err := os.ReadFile(filepath.Join(dataDir, "comments.json")); err == nil {
		json.Unmarshal(data, &comments)
	}
	// Load webhooks
	if data, err := os.ReadFile(filepath.Join(dataDir, "webhooks.json")); err == nil {
		json.Unmarshal(data, &webhooks)
	}
}

func saveData() {
	// Save share links
	if data, err := json.Marshal(shareLinks); err == nil {
		os.WriteFile(filepath.Join(dataDir, "shares.json"), data, 0644)
	}
	// Save comments
	if data, err := json.Marshal(comments); err == nil {
		os.WriteFile(filepath.Join(dataDir, "comments.json"), data, 0644)
	}
	// Save webhooks
	if data, err := json.Marshal(webhooks); err == nil {
		os.WriteFile(filepath.Join(dataDir, "webhooks.json"), data, 0644)
	}
}

// Share link functions
func AddShareLink(link *models.ShareLink) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	shareLinks[link.ID] = link
	saveData()
}

func GetShareLink(id string) *models.ShareLink {
	storageMutex.RLock()
	defer storageMutex.RUnlock()
	return shareLinks[id]
}

func GetAllShareLinks() []*models.ShareLink {
	storageMutex.RLock()
	defer storageMutex.RUnlock()
	links := make([]*models.ShareLink, 0, len(shareLinks))
	for _, link := range shareLinks {
		links = append(links, link)
	}
	return links
}

func DeleteShareLink(id string) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	delete(shareLinks, id)
	saveData()
}

// Comment functions
func AddComment(comment models.FileComment) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	comments[comment.FilePath] = append(comments[comment.FilePath], comment)
	saveData()
}

func GetComments(filePath string) []models.FileComment {
	storageMutex.RLock()
	defer storageMutex.RUnlock()
	return comments[filePath]
}

// Webhook functions
func AddWebhook(webhook *models.Webhook) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	webhooks[webhook.ID] = webhook
	saveData()
}

func GetWebhook(id string) *models.Webhook {
	storageMutex.RLock()
	defer storageMutex.RUnlock()
	return webhooks[id]
}

func GetAllWebhooks() []*models.Webhook {
	storageMutex.RLock()
	defer storageMutex.RUnlock()
	hooks := make([]*models.Webhook, 0, len(webhooks))
	for _, hook := range webhooks {
		hooks = append(hooks, hook)
	}
	return hooks
}

func DeleteWebhook(id string) {
	storageMutex.Lock()
	defer storageMutex.Unlock()
	delete(webhooks, id)
	saveData()
}
