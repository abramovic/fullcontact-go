package fullcontact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"
)

// PersonAPI queries the FullContact Person API
type PersonAPI struct {
	shared *Client
}

// Email returns a person based off of an email addres.
func (c *PersonAPI) Email(value string, webhook *Webhook) (*PersonResponse, error) {
	if value == "" {
		return nil, fmt.Errorf("%s. Missing lookup value", errLibrary)
	}
	r, err := c.shared.get("email", value, "person", webhook)
	if err != nil {
		return nil, err
	}
	return c.search(r)
}

// Twitter returns a person based off of an twitter handle.
func (c *PersonAPI) Twitter(value string, webhook *Webhook) (*PersonResponse, error) {
	if value == "" {
		return nil, fmt.Errorf("%s. Missing lookup value", errLibrary)
	}
	r, err := c.shared.get("twitter", value, "person", webhook)
	if err != nil {
		return nil, err
	}
	return c.search(r)
}

var regexPhone = regexp.MustCompile(`^[1-9]\d{2}-\d{3}-\d{4}`)

// Phone returns a person based off of a phone number.
func (c *PersonAPI) Phone(value string, webhook *Webhook) (*PersonResponse, error) {
	if !regexPhone.Match([]byte(value)) {
		return nil, fmt.Errorf("Invalid Phone: %s. Does not match 123-456-7890", value)
	}
	r, err := c.shared.get("phone", value, "person", webhook)
	if err != nil {
		return nil, err
	}
	return c.search(r)
}

func (c *PersonAPI) search(r *http.Request) (*PersonResponse, error) {
	resp, err := c.shared.do(r)
	if err != nil {
		return nil, err
	}
	var response PersonResponse
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
