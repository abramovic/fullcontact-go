# FullContact Client in Go

[![Go Report Card](https://goreportcard.com/badge/github.com/Abramovic/fullcontact-go)](https://goreportcard.com/report/github.com/Abramovic/fullcontact-go)

This is a Go library built to support [FullContact](http://fullcontact.com)'s REST API.

Most of FullContact's functionality is supported with the exception of batching (coming soon) and Card Reader API.

Feel free to create an issue or send me a pull request if you have any questions, bugs, or suggestions for this library.

- [Go Docs](https://godoc.org/github.com/Abramovic/fullcontact-go)
- [Examples (WIP)](https://github.com/Abramovic/fullcontact-go/tree/master/examples)

## Initialize Client

In order to access the FullContact API you will need to [sign up](https://www.fullcontact.com/developer/) for an account.

FullContact allows developers to pass an API key either in the url query string or as an HTTP request header. For added security we only pass the API key as an HTTP header (X-FullContact-APIKey').

If your application needs to pass the API key in the query string then please create an issue or send me a pull request.

```go
import "github.com/Abramovic/fullcontact-go"

client, err := fullcontact.NewClient("your-api-key")
```

#### Rate Limits

FullContact tracks rate limits on a 60 second basis. If your API is subject to a 10/second rate limit, they will allow you 600 requests per 60 second window.

We will do our best job at keeping your [rate limits](https://www.fullcontact.com/developer/docs/#rate-limiting) up to date. After every request to the FullContact API we will update your rate limit information.

```go
/*
  All of these values are int64
*/
client.RateLimit().Limit // rate limit ceiling for your request
client.RateLimit().Remaining // number of requests left in the 60 second window.
client.RateLimit().Reset // number of UTC epoch seconds remaining until the 60 second window resets
```

## Person API

```go
  client.Person.Email(string, *fullcontact.Webhook) (*fullcontact.PersonResponse, error)
  client.Person.Phone(string, *fullcontact.Webhook) (*fullcontact.PersonResponse, error)
  client.Person.Twitter(string, *fullcontact.Webhook) (*fullcontact.PersonResponse, error)
```
You can query a person by email, phone, and twitter to lookup information about a specific person.

```go
// get person by email address without webhook
person, err := client.Person.Email("john.smith@gmail.com", nil)
if err != nil {
  log.Fatal(err)
}

// get person by phone number without webhook
person, err := client.Person.Phone("555-123-7890", nil)
if err != nil {
  log.Fatal(err)
}

// get person by twitter with a webhook
person, err := client.Person.Twitter("@github", webhook)
if err != nil {
  log.Fatal(err)
}
```

#### Webhook

Webhooks require the callback url that you'd like the data to be posted back to (ie. https://mydomain.com/callback/listener). This client will automatically url encode this value for you.

When you include the pointer to your webhook, an HTTP POST request will be triggered to the URL you've specified. The payload of the response POSTed to the webhook URL is by default formatted as a URL-encoded form with the contents of the "result" form/post parameters being a URL-encoded JSON document.

```go
// Webhook without an ID
webhook, err := fullcontact.NewWebhook("https://mydomain.com/callback/listener")
```

##### Create Webhook with ID

You can include an optional Webhook ID that will be passed back to you in FullContact's response.

```go
// Webhook with an ID (this will be passed back to you)
webhook, err := fullcontact.NewWebhook("https://mydomain.com/callback/listener", "ID-123")
```

Webhook Example Response (URL-Encoded Form)

```html
https://mydomain.com/callback/listener?&webhookId=myID&result=%22status%22%3A200%2C%22contactInfo%22%3A
%7B%22familyName%22%3A%22Lorang%22%2C%22fullName%22%3A%22Bart+Lorang%22%2C%22givenName%22%3A%22Bart%22
  /*** results truncated ***/
}
```

##### FullContact API

NOTE: FullContact will make 3 attempts to deliver the payload, waiting a minimum of 2 seconds between each attempt. If a 200 response is not received within 3 attempts, the request will be dropped.

BLACKLISTING: If FullContact can’t deliver at least 10 consecutive, successful messages to a specified URI over a 5 minute period, we will temporarily impose a 5 minute blacklisting of the URI. After the 5 minutes has elapsed, the blacklist will be automatically removed and FullContact will resume webhook delivery attempts.

## Company API

```go
  client.Company.Domain(string, *fullcontact.Webhook) (*fullcontact.CompanyResponse, error)
```

You can query a company by their domain.

```go
// get company without webhook
company, err := client.Company.Domain("facebook.com", nil)
if err != nil {
  log.Fatal(err)
}
```
