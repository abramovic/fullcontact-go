package fullcontact

// CompanyResponse is the response from FullContact's Company API
/*
{
  "status" : "type":"number",
  "requestId" : "type":"string",
  "category" : [ {
     "name" : "type":"string",
     "code" : "type":"string" // Currently "ADULT", "EMAIL_PROVIDER", "EDUCATION", "SMS", "OTHER"
   } ],
  "logo" : "type":"string",
  "website" : "type":"string",
  "languageLocale" : "type":"string",
  "organization" : {
    "name" : "type":"string",
    "approxEmployees" : "type":"number",
    "founded" : "type":"string",   // ISO 8601 Date
    "overview" : "type":"string",
    "contactInfo" : {
      "emailAddresses" : [ {
        "value" : "type":"string",
        "label" : "type":"string"   // Current labels include "other", "support", "sales", "publicity", "jobs", or "general"
      } ],
      "phoneNumbers" : [ {
        "number" : "type":"string",
        "label" : "type":"string"
      } ],
      "addresses" : [ {
        "addressLine1" : "type":"string",
        "addressLine2" : "type":"string",
        "locality" : "type":"string",
        "region" : {
          "name" : "type":"string",
          "code" : "type":"string"
        },
        "country" : {
          "name" : "type":"string",
          "code" : "type":"string"   // ISO 3166-1 Alpha 2
        },
        "postalCode" : "type":"string",
        "label" : "type":"string"
      } ]
    },
    "keyPeople": [ {
      "name" : "type":"string",
      "title" : "type":"string",
      "link" : "type":"string"   // A link to Person API for this individual. You will need to add your API Key.
    } ],
    "links" : [ {
      "url" : "type":"string",
      "label" : "type":"string"
    }],
    "images" : [ {
      "url" : "type":"string",
      "label" : "type":"string"
    }],
    "keywords" : [ "type":"string" ]
  },
  "socialProfiles" : [ {
    "bio" : "type":"string",
    "followers" : "type":"number",
    "following" : "type":"number",
    "typeId" : "type":"string",
    "typeName" : "type":"string",
    "url" : "type":"string",
    "username" : "type":"string",
    "id" : "type":"string"
  }],
  "traffic" : {
    "topCountryRanking" : [ {
      "rank" : "type":"number",
      "locale" : "type":"string"   // Up to three results of type ISO 3166-1 Alpha 2
    } ],
    "ranking" : [ {
      "rank" : "type":"number",
      "locale" : "type":"string"   // Up to two results, of type "global" and/or "us"
    } ]
  }
}
*/
type CompanyResponse struct {
	Status         int64           `json:"status"`
	RequestID      string          `json:"requestId"`
	Category       Category        `json:"category"`
	Logo           string          `json:"logo"`
	Website        string          `json:"website"`
	Locale         string          `json:"languageLocale"`
	Organization   Organization    `json:"organization"`
	SocialProfiles []SocialProfile `json:"socialProfiles"`
	Traffic        Traffic         `json:"traffic"`
}

// Category is a sub-model of CompanyResponse
type Category struct {
	Name string `json:"name"`
	Code string `json:"code"` // Currently "ADULT", "EMAIL_PROVIDER", "EDUCATION", "SMS", "OTHER"
}

// Organization is a sub-model of CompanyResponse
type Organization struct {
	Name        string                  `json:"name"`
	Employees   int64                   `json:"approxEmployees"`
	Founded     string                  `json:"founded"` // ISO 8601 Date
	Overview    string                  `json:"overview"`
	ContactInfo OrganizationContactInfo `json:"contactInfo"`
	KeyPeople   []KeyPeople             `json:"keyPeople"`
	Links       []Link                  `json:"links"`
	Images      []Link                  `json:"images"`
	Keywords    []string                `json:"keywords"`
}

// OrganizationContactInfo is a sub-model of Organization
type OrganizationContactInfo struct {
	Emails       []ContactInfoEmail   `json:"emailAddresses"`
	PhoneNumbers []ContactInfoEmail   `json:"phoneNumbers"`
	Addresses    []ContactInfoAddress `json:"addresses"`
}

// ContactInfoEmail is a sub-model of OrganizationContactInfo
type ContactInfoEmail struct {
	Value string `json:"value"`
	Label string `json:"label"` // Current labels include "other", "support", "sales", "publicity", "jobs", or "general"
}

// ContactInfoPhone is a sub-model of OrganizationContactInfo
type ContactInfoPhone struct {
	Number string `json:"number"`
	Label  string `json:"label"`
}

// ContactInfoAddress is a sub-model of OrganizationContactInfo
type ContactInfoAddress struct {
	Label        string             `json:"label"`
	AddressLine1 string             `json:"addressLine1"`
	AddressLine2 string             `json:"addressLine2"`
	Locality     string             `json:"locality"`
	PostalCode   string             `json:"postalCode"`
	Region       ContactInfoRegion  `json:"region"`
	Country      ContactInfoCountry `json:"country"`
}

// ContactInfoRegion is a sub-model of ContactInfoAddress
type ContactInfoRegion struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// ContactInfoCountry is a sub-model of ContactInfoAddress
type ContactInfoCountry struct {
	Name string `json:"name"`
	Code string `json:"code"`
}

// KeyPeople is a sub-model of Organization
type KeyPeople struct {
	Name  string `json:"name"`
	Title string `json:"title"`
	Link  string `json:"link"`
}

// Link is a sub-model of Organization
type Link struct {
	Label string `json:"label"`
	URL   string `json:"url"`
}

// Traffic is a sub-model of CompanyResponse
type Traffic struct {
	TopCountries []LocaleRank `json:"topCountryRanking"`
	Ranking      []LocaleRank `json:"ranking"`
}

// LocaleRank is a sub-model of Traffic
type LocaleRank struct {
	Rank   int64  `json:"rank"`
	Locale string `json:"locale"`
}
