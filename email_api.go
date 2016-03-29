package fullcontact

import (
	"encoding/json"
	"fmt"
	"strings"
)

// EmailAPI queries the FullContact Email API
type EmailAPI struct {
	*Client
}

// Get returns information about a specific email address
func (c *EmailAPI) Get(email string) (*EmailResponse, error) {
	if !strings.Contains(email, ".") {
		return nil, fmt.Errorf("%s. Invalid email: %s", errLibrary, email)
	}
	r, err := c.get("email", email, "email/disposable", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	var response EmailResponse
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
