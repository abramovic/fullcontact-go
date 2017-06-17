package fullcontact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

// CompanyAPI queries the FullContact Company API
type CompanyAPI struct {
	shared *Client
}

func (c *CompanyAPI) domain(r *http.Request) (*CompanyResponse, error) {
	resp, err := c.shared.do(r)
	if err != nil {
		return nil, err
	}
	var response CompanyResponse
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&response)
	if err != nil {
		return nil, err
	}
	return &response, nil
}

func (c *CompanyAPI) getDomain(value string, webhook *Webhook) (*http.Request, error) {
	if !strings.Contains(value, ".") { // Instead of a regex let's at least check if there's at least a period in the domain
		return nil, fmt.Errorf("%s. Invalid domain: %s", errLibrary, value)
	}
	r, err := c.shared.get("domain", value, "company/lookup", webhook)
	if err != nil {
		return nil, err
	}
	return r, nil
}

// Domain returns information about a company using their domain name.
func (c *CompanyAPI) Domain(value string, webhook *Webhook) (*CompanyResponse, error) {
	r, err := c.getDomain(value, webhook)
	if err != nil {
		return nil, err
	}
	return c.domain(r)
}
