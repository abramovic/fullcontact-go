# FullContact Client in Go

## Initialize Client

```go

client, err := fullcontact.NewClient("your-api-key")
if err != nil {
  log.Fatal(err)
}

```


## Lookup Person

```go
func (c *Client) Person(lookup string, value string, webhook *fullcontact.Webhook) (*fullcontact.PersonResponse, error)
```

You can look up a person by:

-email
-phone
-twitter

```go
person, err := client.Person("email", "john.smith@gmail.com", nil)
if err != nil {
  log.Fatal(err)
}

```

#### Webhook

You can include a webhook with your requests.

```go
// Webhook without an ID
webhook, err := fullcontact.NewWebhook("http://your-web-hook.url/")

// Webhook with an ID
webhook, err := fullcontact.NewWebhook("http://your-web-hook.url/", "ID-123")

```

#### Person Response

```json
{
  "status": {"type":"number"},
  "requestId": {"type":"string"},
  "likelihood": {"type":"number"},
  "contactInfo": {
    "familyName": {"type":"string"},
    "givenName": {"type":"string"},
    "fullName": {"type":"string"},
    "middleNames":
    [
      {"type":"string"}
    ],
    "websites":
    [
      {
        "url": {"type":"string"}
      }
    ],
    "chats":
    [
      {
        "handle": {"type":"string"},
        "client": {"type":"string"}
      }
    ]
  },
  "demographics": {
    "locationGeneral": {"type":"string"},
    "locationDeduced": {
      "normalizedLocation": {"type":"string"},
      "deducedLocation" : {"type":"string"},
      "city" : {
        "deduced" : {"type":"boolean"},
        "name" : {"type":"string"}
      },
      "state" : {
        "deduced" : {"type":"boolean"},
        "name" : {"type":"string"},
        "code" : {"type":"string"}
      },
      "country" : {
        "deduced" : {"type":"boolean"},
        "name" : {"type":"string"},
        "code" : {"type":"string"}
      },
      "continent" : {
        "deduced" : {"type":"boolean"},
        "name" : {"type":"string"}
      },
      "county" : {
        "deduced" : {"type":"boolean"},
        "name" : {"type":"string"},
        "code" : {"type":"string"}
      },
      "likelihood" : {"type":"number"}
    },
    "age": {"type":"string"},
    "gender": {"type":"string"},
    "ageRange": {"type":"string"}
  },
  "photos":
  [
    {
      "typeId": {"type":"string"},
      "typeName": {"type":"string"},
      "url": {"type":"string"},
      "isPrimary": {"type":"boolean"}
    }
  ],
  "socialProfiles":
  [
    {
      "typeId": {"type":"string"},
      "typeName": {"type":"string"},
      "id": {"type":"string"},
      "username": {"type":"string"},
      "url": {"type":"string"},
      "bio": {"type":"string"},
      "rss": {"type":"string"},
      "following": {"type":"number"},
      "followers": {"type":"number"}
    }
  ],
  "digitalFootprint": {
    "topics":
    [
      {
        "value": {"type":"string"},
        "provider": {"type":"string"}
      }
    ],
    "scores":
    [
      {
        "provider": {"type":"string"},
        "type": {"type":"string"},
        "value": {"type":"number"}
      }
    ]
  },
  "organizations":
  [
    {
      "title": {"type":"string"},
      "name": {"type":"string"},
      "startDate": {"type":"string"},   // formatted as "YYYY-MM"
      "endDate":  {"type":"string"},    // formatted as "YYYY-MM"
      "isPrimary": {"type":"boolean"}
      "current": {"type":"boolean"}
    }
  ]
}
```
