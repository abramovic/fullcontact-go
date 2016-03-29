package fullcontact

import (
	"fmt"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

var (
	// FullContact API
	defaultEndpoint = "https://api.fullcontact.com/v2"
	// MaxIdleConnections for shared http client
	MaxIdleConnections = 10
	// RequestTimeout for shared http client
	RequestTimeout = 60
)

// Client for fullcontact-go
type Client struct {
	apikey     string
	httpClient *http.Client
	limit      *RateLimit

	Account *AccountAPI
	Email   *EmailAPI
	Person  *PersonAPI
	Company *CompanyAPI
}

func (c *Client) configure() (err error) {
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	c.limit = &RateLimit{}
	c.Account = &AccountAPI{c}
	c.Email = &EmailAPI{c}
	c.Person = &PersonAPI{c}
	c.Company = &CompanyAPI{c}
	return nil
}

// NewClient returns a client for fullcontact-go
func NewClient(apikey string) (client *Client, err error) {
	client = &Client{apikey: apikey}
	err = client.configure()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	resp, err := c.httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("%s. HTTP Response: %s", errUnknown, err.Error())
	}
	c.limit.Limit, _ = strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Limit"), 0, 64)
	c.limit.Remaining, _ = strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Remaining"), 0, 64)
	c.limit.Reset, _ = strconv.ParseInt(resp.Header.Get("X-Rate-Limit-Reset"), 0, 64)
	c.limit.Updated = time.Now().UTC()
	if resp.StatusCode == 400 {
		return nil, errStatus400
	}
	if resp.StatusCode == 403 {
		return nil, fmt.Errorf("%s: Rate Limit Hit. Please wait %d seconds before trying again", libraryName, c.limit.Reset)
	}
	if resp.StatusCode == 404 {
		return nil, errStatus404
	}
	if resp.StatusCode == 405 {
		return nil, errStatus405
	}
	if resp.StatusCode == 410 {
		return nil, fmt.Errorf("%s. Please contact github.com/Abramovic to update the library", errStatus410)
	}
	if resp.StatusCode == 422 {
		return nil, errStatus422
	}
	if resp.StatusCode == 500 {
		return nil, errStatus500
	}
	if resp.StatusCode == 503 {
		// FullContact does not document the specific header (is it X-Retry-After?), only that it contains 'Retry-After'
		for header, _ := range resp.Header {
			if strings.Contains(header, "Retry-After") {
				return nil, fmt.Errorf("%s. Retry After: %s", errStatus503, resp.Header.Get(header))
			}
		}
		return nil, errStatus503
	}
	return resp, nil
}

func (c *Client) get(search, value, endPoint string, webhook *Webhook) (*http.Request, error) {
	fullURL := fmt.Sprintf("%s/%s?%s=%s", defaultEndpoint, fmt.Sprintf("%s.json", endPoint), search, url.QueryEscape(value))
	if webhook != nil && webhook.URL != "" {
		fullURL = fmt.Sprintf("%s&webhookUrl=%s", fullURL, url.QueryEscape(webhook.URL))
		if webhook.ID != "" {
			fullURL = fmt.Sprintf("%s&webhookId=%s", fullURL, url.QueryEscape(webhook.ID))
		}
	}
	r, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("X-FullContact-APIKey", c.apikey)
	return r, nil
}
