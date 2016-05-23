package fullcontact

import (
	"encoding/json"
	"fmt"
)

// PersonAPI queries the FullContact Person API
type PersonAPI struct {
	*Client
}

// Email returns a person based off of an email addres.
func (c *PersonAPI) Email(value string, webhook *Webhook) (*PersonResponse, error) {
	return c.search("email", value, webhook)
}

// Twitter returns a person based off of an twitter handle.
func (c *PersonAPI) Twitter(value string, webhook *Webhook) (*PersonResponse, error) {
	return c.search("twitter", value, webhook)
}

// Phone returns a person based off of a phone number.
func (c *PersonAPI) Phone(value string, webhook *Webhook) (*PersonResponse, error) {
	return c.search("phone", value, webhook)
}

func (c *PersonAPI) search(search, value string, webhook *Webhook) (*PersonResponse, error) {
	if value == "" {
		return nil, fmt.Errorf("%s. Missing lookup value", errLibrary)
	}
	r, err := c.get(search, value, "person", webhook)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}

	var response PersonResponse
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	if search == "email" {
		response.Email = value
	}
	return &response, nil
}
