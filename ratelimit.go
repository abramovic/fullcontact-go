package fullcontact

import (
	"time"
)

// RateLimit gets updated after each HTTP response from FullContact's API
type RateLimit struct {
	Limit     int64     `json:"limit"`
	Remaining int64     `json:"remaining"`
	Reset     int64     `json:"reset"`
	Updated   time.Time `json:"updated"`
}

// RateLimit returns the current rate limit information.
func (c *Client) RateLimit() RateLimit {
	return *c.limit
}
