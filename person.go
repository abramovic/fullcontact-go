package fullcontact

import (
	"encoding/json"
	"fmt"
)

// PersonResponse is the response from FullContact's Person API
/*
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
*/
type PersonResponse struct {
	Status           int64                  `json:"status"`
	RequestID        string                 `json:"requestId"`
	Likelidhood      float64                `json:"likelihood"`
	Photos           []PersonPhoto          `json:"photos"`
	ContactInfo      PersonContactInfo      `json:"confactInfo"`
	Organizations    []PersonOrganization   `json:"organizations"`
	Demographics     PersonDemographics     `json:"demographics"`
	SocialProfiles   []SocialProfile        `json:"socialProfiles"`
	DigitalFootprint PersonDigitalFootprint `json:"digitalFootprint"`
}

// PersonPhoto is a sub-model of PersonResponse
type PersonPhoto struct {
	Type      string `json:"type"`
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

// PersonContactInfo is a sub-model of PersonResponse
type PersonContactInfo struct {
	Websites   []Website `json:"websites"`
	FamilyName string    `json:"familyName"`
	FullName   string    `json:"fullName"`
	GivenName  string    `json:"givenName"`
}

// Website is a sub-model of PersonContactInfo
type Website struct {
	URL string `json:"url"`
}

// PersonOrganization is a sub-model of PersonResponse
type PersonOrganization struct {
	Name      string `json:"name"`
	IsPrimary bool   `json:"isPrimary"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Title     string `json:"title"`
	IsCurrent bool   `json:"current"`
}

// PersonDemographics is a sub-model of PersonResponse
type PersonDemographics struct {
	LocationDeduced LocationDeduced `json:"locationDeduced"`
	LocationGeneral string          `json:"locationGeneral"`
	Gender          string          `json:"gender"`
	Age             string          `json:"age"`
	AgeRange        string          `json:"ageRange"`
}

// LocationDeduced is a sub-model of PersonDemographics
type LocationDeduced struct {
	NormalizedLocation string       `json:"normalizedLocation"`
	DeducedLocation    string       `json:"deducedLocation"`
	Likelihood         float64      `json:"likelihood"`
	City               LocationPart `json:"city"`
	County             LocationPart `json:"county"`
	State              LocationPart `json:"state"`
	Country            LocationPart `json:"country"`
	Continent          LocationPart `json:"continent"`
}

// LocationPart is a sub-model of LocationDeduced
type LocationPart struct {
	IsDeduced bool   `json:"deduced"`
	Name      string `json:"name"`
	Code      string `json:"code"`
}

// SocialProfile is a sub-model of PersonResponse
type SocialProfile struct {
	ID string `json:"id"`
	SocialProfileShared
}

// SocialProfileinterface is used because FullContact uses both ints and strings
type SocialProfileinterface struct {
	ID interface{} `json:"id"`
	SocialProfileShared
}

// SocialProfileShared are the attributes that have shown to be consistent with SocialProfile
type SocialProfileShared struct {
	Type      string `json:"type"`
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Followers int64  `json:"followers"`
	Following int64  `json:"following"`
	RSS       string `json:"rss"`
}

// UnmarshalJSON needs to be refactored. It does what it can to ensure that the ID is always a string
func (s *SocialProfile) UnmarshalJSON(data []byte) error {
	var aux SocialProfileinterface
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// TODO: we can clean this up
	s.Type = aux.Type
	s.TypeID = aux.TypeID
	s.TypeName = aux.TypeName
	s.URL = aux.URL
	s.Username = aux.Username
	s.Bio = aux.Bio
	s.Followers = aux.Followers
	s.Following = aux.Following
	s.RSS = aux.RSS

	// FullContact is not consistent between "type" and  "typeId"
	if s.Type == "" {
		s.Type = s.TypeID
	}
	if s.TypeID == "" {
		s.TypeID = s.Type
	}

	// FullContact is not consistent with giving us the same ID type
	if value, ok := aux.ID.(string); ok {
		s.ID = value
		return nil
	}
	if value, ok := aux.ID.(uint); ok {
		s.ID = fmt.Sprintf("%d", value)
		return nil
	}
	if value, ok := aux.ID.(int); ok {
		s.ID = fmt.Sprintf("%d", value)
		return nil
	}
	if value, ok := aux.ID.(int32); ok {
		s.ID = fmt.Sprintf("%d", value)
		return nil
	}
	if value, ok := aux.ID.(int64); ok {
		s.ID = fmt.Sprintf("%d", value)
		return nil
	}
	if value, ok := aux.ID.(float32); ok {
		s.ID = fmt.Sprintf("%d", int64(value))
		return nil
	}
	if value, ok := aux.ID.(float64); ok {
		s.ID = fmt.Sprintf("%d", int64(value))
		return nil
	}
	return fmt.Errorf("Could not convert primary ID %v", aux.ID)
}

// PersonDigitalFootprint is a sub-model of PersonResponse
type PersonDigitalFootprint struct {
	Scores []DigitalFootPrintScore `json:"scores"`
	Topics []DigitalFootPrintTopic `json:"topics"`
}

// DigitalFootPrintScore is a sub-model of PersonDigitalFootprint
type DigitalFootPrintScore struct {
	Provider string      `json:"provider"`
	Type     string      `json:"general"`
	Value    interface{} `json:"value"`
}

// DigitalFootPrintTopic is a sub-model of PersonDigitalFootprint
type DigitalFootPrintTopic struct {
	Provider string `json:"provider"`
	Value    string `json:"value"`
}
