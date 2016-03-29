package fullcontact

import (
	"encoding/json"
	"fmt"
	"strings"
)

// CompanyAPI queries the FullContact Company API
type CompanyAPI struct {
	*Client
}

// Domain returns information about a company using their domain name.
func (c *CompanyAPI) Domain(value string, webhook *Webhook) (*CompanyResponse, error) {
	if !strings.Contains(value, ".") { // Instead of a regex let's at least check if there's at least a period in the domain
		return nil, fmt.Errorf("%s. Invalid domain: %s", errLibrary, value)
	}
	r, err := c.get("domain", value, "company", webhook)
	if err != nil {
		return nil, err
	}
	resp, err := c.do(r)
	if err != nil {
		return nil, err
	}
	var response CompanyResponse
	decoder := json.NewDecoder(resp.Body)
	defer resp.Body.Close()
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}
