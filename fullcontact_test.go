package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestNewClient(t *testing.T) {
	client, err := NewClient("abc-123")
	if err != nil {
		t.Errorf("New Client: %s", err.Error())
	}
	if client == nil {
		t.Errorf("New Client: missing client")
	}
}

func TestNewClientMissingApiKey(t *testing.T) {
	client, err := NewClient("")
	if err == nil {
		t.Errorf("New Client: unknown error")
	}
	if client != nil {
		t.Errorf("New Client: client should be nil")
	}
}

func TestClientGet(t *testing.T) {
	client, _ := NewClient("abc-123")
	_, err := client.get("email", "john@gmail.com", "person", nil)
	if err != nil {
		t.Errorf("Client Get: %s", err.Error())
	}
}
