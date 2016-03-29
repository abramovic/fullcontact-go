package fullcontact

import (
	"testing"
)

// TODO: need to get a public api key for testing
func TestPersonEmail(t *testing.T) {
	client, _ := NewClient("abc-123")

	_, err := client.Person.Email("john@gmail.com", nil)
	if err == nil {
		t.Errorf("Person API: Email - %s", "an error should have occured")
	}
	// _, err = client.Person.Twitter("facebook", nil)
	// if err != nil {
	// 	t.Errorf("Person API: Twitter - %s", err.Error())
	// }
	// _, err = client.Person.Phone("555-123-4567", nil)
	// if err != nil {
	// 	t.Errorf("Person API: Phone - %s", err.Error())
	// }
}
