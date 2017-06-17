package fullcontact

// AccountResponse is the response from FullContact's Account API
/*
{
  "status": {"type":"number"},  // common status codes can be found in diagram below
  "periodStart": {"type":"string"},
  "periodEnd": {"type":"string"},
  "plan": {"type":"string"},
  "planBasePrice": {"type":"number"},
  "planOveragePrice": {"type":"number"},
  "applicationId": {"type":"string"},
  "metrics":
  [
    {
      "metricName": {"type":"string"},  // usage metrics can be found by logging into your account portal
      "metricId": {"type":"string"},
      "planLevel": {"type":"number"},
      "usage": {"type":"number"},
      "remaining": {"type":"number"},
      "overage": {"type":"number"}
    }
  ]
}
*/
type AccountResponse struct {
	ID               string          `json:"applicationId"`
	Status           int             `json:"status"`
	PeriodStart      string          `json:"periodStart"`
	PeriodEnd        string          `json:"periodEnd"`
	PlanName         string          `json:"plan"`
	PlanBasePrice    float64         `json:"planBasePrice"`
	PlanOveragePrice float64         `json:"planOveragePrice"`
	Metrics          []AccountMetric `json:"metrics"`
}

// AccountMetric is a sub-model of AccountResponse
type AccountMetric struct {
	ID        string `json:"metricId"`
	Name      string `json:"metricName"`
	PlanLevel int64  `json:"planLevel"`
	Usage     int64  `json:"usage"`
	Remaining int64  `json:"remaining"`
	Overage   int64  `json:"overage"`
}
