package fullcontact

import (
	"testing"
)

func TestRateLimit(t *testing.T) {
	client, _ := NewClient("")

	if client.RateLimit().Limit < 0 {
		t.Errorf("RateLimit: - %s", "negative limit")
		return
	}

}
