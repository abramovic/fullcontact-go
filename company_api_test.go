package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestCompanyDomain(t *testing.T) {
	client, _ := NewClient("abc-123")
	
	_, err := client.Company.Domain("localhost", nil)
	if err == nil {
		t.Errorf("Company API: %s", "an error should have occured")
	}
}
