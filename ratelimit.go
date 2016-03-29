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
	if c.limit != nil {
		return *c.limit
	}
	account, err := c.Account.Get()
	if err != nil {
		return RateLimit{}
	}
	c.limit = &RateLimit{}
	for _, metric := range account.Metrics {
		if metric.PlanLevel > 0 {
			c.limit.Limit = metric.PlanLevel
			c.limit.Remaining = metric.Remaining
		}
	}
	return *c.limit
}
