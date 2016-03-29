package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestEmailGet(t *testing.T) {
	client, _ := NewClient("abc-123")

	_, err := client.Email.Get("")
	if err == nil {
		t.Errorf("Email API: %s", "an error should have occured")
	}
}
