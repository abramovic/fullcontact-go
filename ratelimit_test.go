package fullcontact

import (
	"testing"
)

func TestRateLimit(t *testing.T) {
	client, _ := NewClient("abc-123")

	if client.RateLimit().Limit < 0 {
		t.Errorf("RateLimit: - %s", "negative limit")
	}
}
