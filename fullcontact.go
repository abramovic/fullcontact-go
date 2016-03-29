package fullcontact

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	defaultEndpoint = "https://api.fullcontact.com/v2"
)

var (
	MaxIdleConnections = 10
	RequestTimeout     = 60

	errUnknown = fmt.Errorf("Error unknown")
	status400  = fmt.Errorf("Your request was malformed.")
	status403  = fmt.Errorf("Your API key is invalid, missing, or has exceeded its quota. **Paid plans will not receive a 403 response when they exceed their alotted matches. They will only receive a 403 for exceeding rate limit quotas.")
	status404  = fmt.Errorf("This person was searched in the past 24 hours and nothing was found.")
	status405  = fmt.Errorf("You have queried the API with an unsupported HTTP method. Retry your query with either GET or POST.")
	status410  = fmt.Errorf("This resource cannot be found. You will receive this status code if you attempt to query our deprecated V1 endpoints.")
	status422  = fmt.Errorf("Invalid or missing API query parameter.")
	status500  = fmt.Errorf("There was an unexpected error on our server. If you see this please contact support@fullcontact.com.")
	status503  = fmt.Errorf("There is a transient downstream error condition. We include a 'Retry-After' header dictating when to attempt the call again.")
)

type Client struct {
	apikey     string
	httpClient *http.Client
}

func NewClient(apikey string) (client *Client, err error) {
	client = &Client{apikey: apikey}
	err = client.createHTTPClient()
	if err != nil {
		return nil, err
	}
	return client, nil
}

func (c *Client) configure() (err error) {
	err = c.createHTTPClient()
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) createHTTPClient() error {
	c.httpClient = &http.Client{
		Transport: &http.Transport{
			MaxIdleConnsPerHost: MaxIdleConnections,
		},
		Timeout: time.Duration(RequestTimeout) * time.Second,
	}
	return nil
}

type Webhook struct {
	ID  string `json:"webhookId"`
	URL string `json:"webhookUrl"`
}

func NewWebhook(url string, opts ...string) (webhook *Webhook, err error) {
	id := ""
	if url == "" {
		return nil, fmt.Errorf("Missing Required URL")
	}
	if len(opts) == 1 {
		id = opts[0]
	}
	return &Webhook{ID: id, URL: url}, nil
}

func (c *Client) Person(search, value string, webhook *Webhook) (*PersonResponse, error) {
	search = strings.ToLower(search)
	switch search {
	case "email":
		search = "email"
	case "phone":
		search = "phone"
	case "twitter":
		search = "twitter"
	default:
		return nil, fmt.Errorf("Unknown search type: %s", search)
	}
	if value == "" {
		return nil, fmt.Errorf("Missing lookup value")
	}

	fullURL := fmt.Sprintf("%s/%s?%s=%s", defaultEndpoint, "person.json", search, url.QueryEscape(value))
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

	resp, err := c.httpClient.Do(r)
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
	return &response, nil
}
