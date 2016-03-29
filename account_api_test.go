package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestAccountGet(t *testing.T) {
	client, _ := NewClient("abc-123")

	account, err := client.Account.Get()
	if err != nil {
		t.Errorf("Account API: %s", err.Error())
	}
	if account == nil {
		t.Errorf("Account API: %s", "missing response")
	}
}
