package fullcontact

import (
	"encoding/json"
)

// AccountAPI queries the FullContact Account API
type AccountAPI struct {
	*Client
}

// Get returns the current account information.
func (c *AccountAPI) Get() (*AccountResponse, error) {
	r, err := c.get("", "", "stats", nil)
	if err != nil {
		return nil, err
	}
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, err
	}
	var response AccountResponse
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
