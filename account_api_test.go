package fullcontact

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestAccountGet(t *testing.T) {

	client, _ := NewClient("")

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}
	account, err := client.Account.Get()
	if err != nil {
		t.Errorf("Account API: %s", err.Error())
	}
	if account.Status != 200 {
		t.Errorf("Expected Account Status 200. Got %v", account.Status)
		return
	}

	if len(account.Metrics) == 0 {
		t.Errorf("Expected Account Metrics. Got 0.")
		return
	}

}

func TestAccountget(t *testing.T) {

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, accountJSON)
		}),
	)
	defer server.Close()

	r, _ := http.NewRequest("GET", server.URL, nil)

	client, _ := NewClient("")

	account, err := client.Account.get(r)
	if err != nil {
		t.Errorf("Account API: %s", err.Error())
		return
	}

	if account.Status != 200 {
		t.Errorf("Expected Account Status 200. Got %v", account.Status)
		return
	}

	if len(account.Metrics) == 0 {
		t.Errorf("Expected Account Metrics. Got 0.")
		return
	}

}

var accountJSON = `{
  "status" : 200,
  "periodStart" : "2017-06-01T00:00:00 -0600",
  "periodEnd" : "2017-06-30T23:59:59 -0600",
  "plan" : "Free (500 matches / mo)",
  "planBasePrice" : 0.0,
  "planOveragePrice" : 0.0,
  "applicationId" : "12345",
  "metrics" : [ {
    "metricName" : "Success - Person  (\"200\")",
    "metricId" : "200",
    "planLevel" : 500,
    "usage" : 0,
    "remaining" : 500
  }, {
    "metricName" : "Success - Disposable Email (\"200\")",
    "metricId" : "200_disposable",
    "planLevel" : 10,
    "usage" : 0,
    "remaining" : 10
  }, {
    "metricName" : "Name/Location/Icon/Stats",
    "metricId" : "200_free",
    "planLevel" : 100000,
    "usage" : 0,
    "remaining" : 100000
  }, {
    "metricName" : "Queued for Search - Person (\"202\")",
    "metricId" : "202",
    "usage" : 0
  }, {
    "metricName" : "Success - Card Reader (\"202\") Low",
    "metricId" : "202_CardShark",
    "planLevel" : 50,
    "usage" : 0,
    "remaining" : 50
  }, {
    "metricName" : "Success - Card Reader (\"202\") High",
    "metricId" : "202_CardShark_High",
    "planLevel" : 0,
    "usage" : 0
  }, {
    "metricName" : "Success - Card Reader (\"202\") Medium",
    "metricId" : "202_CardShark_Medium",
    "planLevel" : 0,
    "usage" : 0
  }, {
    "metricName" : "Bad Request (\"400\")",
    "metricId" : "400",
    "usage" : 0
  }, {
    "metricName" : "Unauthorized - Name/Location/Icon/Stats (\"403\")",
    "metricId" : "403_free",
    "usage" : 0
  }, {
    "metricName" : "No Results Found - Person (\"404\")",
    "metricId" : "404",
    "usage" : 0
  }, {
    "metricName" : "No Results - Name/Location/Icon/Stats (\"404\")",
    "metricId" : "404_free",
    "usage" : 0
  }, {
    "metricName" : "Invalid - Name/Location/Icon/Stats (\"422\")",
    "metricId" : "422_free",
    "usage" : 0
  }, {
    "metricName" : "Error - Name/Location/Icon/Stats (\"500\")",
    "metricId" : "500_free",
    "usage" : 0
  }, {
    "metricName" : "/v2/cardReader/GET",
    "metricId" : "_v2_cardShark_GET",
    "usage" : 0
  }, {
    "metricName" : "/v2/cardReader/POST",
    "metricId" : "_v2_cardShark_POST",
    "usage" : 0
  }, {
    "metricName" : "/v2/company",
    "metricId" : "_v2_company_lookup",
    "usage" : 0
  }, {
    "metricName" : "/v2/person",
    "metricId" : "_v2_person",
    "usage" : 0
  }, {
    "metricName" : "Total API Requests",
    "metricId" : "api_requests",
    "usage" : 0
  }, {
    "metricName" : "Success - Company (\"200\")",
    "metricId" : "company_200",
    "planLevel" : 500,
    "usage" : 0,
    "remaining" : 500
  }, {
    "metricName" : "Company API Total Requests",
    "metricId" : "company_rate_limit",
    "planLevel" : 60,
    "usage" : 0
  }, {
    "metricName" : "Feature - Company Search",
    "metricId" : "feature_company_search",
    "planLevel" : 500,
    "usage" : 0,
    "remaining" : 500
  }, {
    "metricName" : "Person API Total Requests",
    "metricId" : "person_rate_limit",
    "planLevel" : 60,
    "usage" : 0
  } ]
}`
