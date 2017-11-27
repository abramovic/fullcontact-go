package fullcontact

import (
	"net/http"
	"os"
	"testing"
)

// TODO: need to get a public api key for testing
func TestNewClient(t *testing.T) {
	client, err := NewClient("")
	if err != nil {
		t.Errorf("New Client: %s", err.Error())
		return
	}

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}

	req404, _ := http.NewRequest("GET", "https://api.fullcontact.com/v3", nil)

	_, err404 := client.do(req404)
	if err404 != ErrStatus404 {
		t.Errorf("Expected 404 from FullContact")
		return
	}

	req403, _ := http.NewRequest("GET", "https://api.fullcontact.com/v2/company/lookup.json?domain=github.com&apiKey=abd123", nil)

	_, err403 := client.do(req403)
	if err403 != ErrStatus403 {
		t.Errorf("Expected 403 from FullContact")
		return
	}

	req405, _ := http.NewRequest("DELETE", "https://api.fullcontact.com/v2/company/lookup.json?domain=github.com&apiKey=abd123", nil)
	_, err405 := client.do(req405)
	if err405 != ErrStatus405 {
		t.Errorf("Expected 405 from FullContact. Only GET and POST methods are allowed")
		return
	}

}

func TestClientGet(t *testing.T) {
	client, _ := NewClient("")
	_, err := client.get("email", "john@gmail.com", "person", nil)
	if err != nil {
		t.Errorf("Client Get: %s", err.Error())
		return
	}
}
