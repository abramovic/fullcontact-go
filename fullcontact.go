package fullcontact

import (
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

var (
	// MaxIdleConnections for shared http client
	MaxIdleConnections = 10
	// RequestTimeout for shared http client
	RequestTimeout = 60

	domain   = "api.fullcontact.com"
	protocol = "https"
)

// Client for fullcontact-go
type Client struct {
	apikey     string
	httpClient *http.Client
	limit      *RateLimit

	Account *AccountAPI
	Person  *PersonAPI
	Company *CompanyAPI
}

var httpClient = &http.Client{
	Transport: &http.Transport{
		MaxIdleConnsPerHost: MaxIdleConnections,
	},
	Timeout: time.Duration(RequestTimeout) * time.Second,
}

func (c *Client) configure() {
	c.limit = &RateLimit{}
	c.Account = &AccountAPI{c}
	c.Person = &PersonAPI{c}
	c.Company = &CompanyAPI{c}
}

// NewClient returns a client for fullcontact-go
func NewClient(apikey string) (client *Client, err error) {
	if apikey == "" {
		apikey = os.Getenv("FULLCONTACT_API")
	}
	client = &Client{apikey: apikey}
	client.configure()
	return client, nil
}

func (c *Client) do(r *http.Request) (*http.Response, error) {
	resp, err := httpClient.Do(r)
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
		return nil, errStatus403
	}
	if resp.StatusCode == 404 {
		return nil, errStatus404
	}
	if resp.StatusCode == 405 {
		return nil, errStatus405
	}
	if resp.StatusCode == 422 {
		return nil, errStatus422
	}
	if resp.StatusCode == 500 {
		return nil, errStatus500
	}
	return resp, nil
}

func (c *Client) get(search, value, endPoint string, webhook *Webhook) (*http.Request, error) {
	u := &url.URL{
		Scheme: protocol,
		Host:   domain,
		Path:   fmt.Sprintf("v2/%s.json", endPoint),
	}
	q := u.Query()
	q.Set(search, value)
	if webhook != nil && webhook.URL != "" {
		q.Set("webhookUrl", url.QueryEscape(webhook.URL))
		if webhook.ID != "" {
			q.Set("webhookId", url.QueryEscape(webhook.ID))
		}
	}
	u.RawQuery = q.Encode()
	r, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}
	r.Header.Set("X-FullContact-APIKey", c.apikey)
	return r, nil
}
