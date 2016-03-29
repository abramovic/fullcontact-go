package fullcontact

import (
	"encoding/json"
	"fmt"
)

type PersonResponse struct {
	Status           int                      `json:"status"`
	RequestID        string                   `json:"requestId"`
	Likelidhood      float64                  `json:"likelihood"`
	Photos           []PhotoResponse          `json:"photos"`
	ContactInfo      ContactInfoResponse      `json:"confactInfo"`
	Organizations    []OrganizationResponse   `json:"organizations"`
	Demographics     DemographicsResponse     `json:"demographics"`
	SocialProfiles   []SocialProfileResponse  `json:"socialProfiles"`
	DigitalFootprint DigitalFootprintResponse `json:"digitalFootprint"`
}

type PhotoResponse struct {
	Type      string `json:"type"`
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	IsPrimary bool   `json:"isPrimary"`
}

type ContactInfoResponse struct {
	Websites   []WebsiteResponse `json:"websites"`
	FamilyName string            `json:"familyName"`
	FullName   string            `json:"fullName"`
	GivenName  string            `json:"givenName"`
}

type WebsiteResponse struct {
	URL string `json:"url"`
}

type OrganizationResponse struct {
	Name      string `json:"name"`
	IsPrimary bool   `json:"isPrimary"`
	StartDate string `json:"startDate"`
	EndDate   string `json:"endDate"`
	Title     string `json:"title"`
	IsCurrent bool   `json:"current"`
}

type DemographicsResponse struct {
	LocationDeduced LocationDeducedResponse `json:"locationDeduced"`
	LocationGeneral string                  `json:"locationGeneral"`
	Gender          string                  `json:"gender"`
	Age             string                  `json:"age"`
	AgeRange        string                  `json:"ageRange"`
}

type LocationDeducedResponse struct {
	NormalizedLocation string               `json:"normalizedLocation"`
	DeducedLocation    string               `json:"deducedLocation"`
	Likelihood         float64              `json:"likelihood"`
	City               LocationPartResponse `json:"city"`
	County             LocationPartResponse `json:"county"`
	State              LocationPartResponse `json:"state"`
	Country            LocationPartResponse `json:"country"`
	Continent          LocationPartResponse `json:"continent"`
}

type LocationPartResponse struct {
	IsDeduced bool   `json:"deduced"`
	Name      string `json:"name"`
	Code      string `json:"code"`
}

type SocialProfileResponse struct {
	ID string `json:"id"`
	SocialProfileShared
}

type SocialProfileResponseInterface struct {
	ID interface{} `json:"id"`
	SocialProfileShared
}

type SocialProfileShared struct {
	Type      string `json:"type"`
	TypeID    string `json:"typeId"`
	TypeName  string `json:"typeName"`
	URL       string `json:"url"`
	Username  string `json:"username"`
	Bio       string `json:"bio"`
	Followers int    `json:"followers"`
	Following int    `json:"following"`
	RSS       string `json:"rss"`
}

func (s *SocialProfileResponse) UnmarshalJSON(data []byte) error {
	var aux SocialProfileResponseInterface
	err := json.Unmarshal(data, &aux)
	if err != nil {
		return err
	}

	// TODO: we can clean this up
	s.Type = s.Type
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
		s.ID = fmt.Sprintf("%d", int(value))
		return nil
	}
	if value, ok := aux.ID.(float64); ok {
		s.ID = fmt.Sprintf("%d", int(value))
		return nil
	}
	return fmt.Errorf("Could not convert primary ID %v", aux.ID)
}

type DigitalFootprintResponse struct {
	Scores []DigitalFootPrintScore `json:"scores"`
	Topics []DigitalFootPrintTopic `json:"topics"`
}

type DigitalFootPrintScore struct {
	Provider string      `json:"provider"`
	Type     string      `json:"general"`
	Value    interface{} `json:"value"`
}

type DigitalFootPrintTopic struct {
	Provider string `json:"provider"`
	Value    string `json:"value"`
}
