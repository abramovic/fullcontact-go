package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestWebhook(t *testing.T) {
	_, err := NewWebhook("")
	if err == nil {
		t.Errorf("NewWebhook: - %s", "an error should have occured")
		return
	}
	webhook, err := NewWebhook("localhost/endpoint", "hookID", "callback")
	if webhook.ID != "hookID" {
		t.Errorf("NewWebhook: - %s", "hook id was not properly set")
	}
}
