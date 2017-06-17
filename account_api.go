package fullcontact

import (
	"encoding/json"
	"net/http"
)

// AccountAPI queries the FullContact Account API
type AccountAPI struct {
	shared *Client
}

func (c *AccountAPI) get(r *http.Request) (*AccountResponse, error) {
	resp, err := httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	var response AccountResponse
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

// Get returns the current account information.
func (c *AccountAPI) Get() (*AccountResponse, error) {
	r, err := c.shared.get("", "", "stats", nil)
	if err != nil {
		return nil, err
	}
	return c.get(r)
}
