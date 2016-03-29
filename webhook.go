package fullcontact

import (
	"fmt"
)

// Webhook requires a URL with an optional ID
type Webhook struct {
	ID  string `json:"webhookId"`
	URL string `json:"webhookUrl"`
}

// NewWebhook returns a new webhook pointer
func NewWebhook(url string, opts ...string) (webhook *Webhook, err error) {
	id := ""
	if url == "" {
		return nil, fmt.Errorf("%s. Missing URL for Webhook", errLibrary)
	}
	if len(opts) > 0 {
		id = opts[0]
	}
	return &Webhook{ID: id, URL: url}, nil
}
