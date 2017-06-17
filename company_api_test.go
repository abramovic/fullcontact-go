package fullcontact

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func TestCompanyDomain(t *testing.T) {
	client, _ := NewClient("")

	_, err := client.Company.Domain("localhost", nil)
	if err == nil {
		t.Errorf("Company API: %s", "the domain 'localhost' should be invalid as it's missing a .")
		return
	}

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}

	domain, err := client.Company.Domain("github.com", nil)
	if err != nil {
		t.Errorf("Company API: %s", err.Error())
		return
	}

	if domain.Status != 200 {
		t.Errorf("Expected Domain Status 200. Got %v", domain.Status)
		return
	}

	if domain.Category[0].Code != "OTHER" {
		t.Errorf("Expected Domain Category Code OTHER. Got %v", domain.Category[0].Code)
		return
	}

}

func TestCompanyDgetDomain(t *testing.T) {
	client, _ := NewClient("abcdef")

	r, err := client.Company.getDomain("github.com", nil)
	if err != nil {
		t.Errorf("Company API: %s", err.Error())
		return
	}
	if r.URL.String() != "https://api.fullcontact.com/v2/company/lookup.json?domain=github.com" {
		t.Errorf("Company API: Got invalid URL: %s", r.URL.String())
		return
	}
}

func TestCompanydomain(t *testing.T) {
	client, _ := NewClient("")

	server := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, domainJSON)
		}),
	)
	defer server.Close()

	r, _ := http.NewRequest("GET", server.URL, nil)

	domain, err := client.Company.domain(r)
	if err != nil {
		t.Errorf("Company API: %s", err.Error())
		return
	}

	if domain.Status != 200 {
		t.Errorf("Expected Domain Status 200. Got %v", domain.Status)
		return
	}

	if domain.Category[0].Code != "OTHER" {
		t.Errorf("Expected Domain Category Code OTHER. Got %v", domain.Category[0].Code)
		return
	}
}

var domainJSON = `{
  "status" : 200,
  "requestId" : "7ce74ced-562b-40f7-bd5e-958bf04dcb29",
  "category" : [ {
    "code" : "OTHER",
    "name" : "Other"
  } ],
  "logo" : "https://d2ojpxxtu63wzl.cloudfront.net/static/2f29a8b707850060264ce907fe3a1e78_7b420235df579597868d6af48d04eb522bd77c14d10362ac6cc2ae666a52c298",
  "website" : "http://github.com",
  "organization" : {
    "name" : "GitHub, Inc.",
    "approxEmployees" : 750,
    "founded" : "2008",
    "overview" : "GitHub is the best way to build software together. Over ten million people use GitHub to share code and build amazing things with friends, co-workers, classmates, and complete strangers. Whether it's your company's application or a powerful open source project, GitHub helps everyone work together better by providing tools for easier collaboration and code sharing. With the collaborative features of GitHub.com, GitHub desktop and mobile applications, and GitHub Enterprise, it has never been easier for individuals and teams to write better code, faster.",
    "contactInfo" : {
      "emailAddresses" : [ {
        "value" : "support@github.com",
        "label" : "support"
      } ],
      "phoneNumbers" : [ {
        "number" : "+1 (415) 448-6673",
        "label" : "other"
      }, {
        "number" : "+1 (415) 371-0060",
        "label" : "other"
      } ],
      "addresses" : [ {
        "addressLine1" : "88 Colin P Kelly Jr. Street",
        "locality" : "San Francisco",
        "region" : {
          "name" : "California",
          "code" : "CA"
        },
        "country" : {
          "name" : "United States",
          "code" : "US"
        },
        "postalCode" : "94107",
        "label" : "other"
      } ]
    },
    "links" : [ {
      "url" : "https://en-gb.facebook.com/officialctf/about",
      "label" : "facebook"
    }, {
      "url" : "http://github.com/blog",
      "label" : "github"
    } ],
    "images" : [ {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/642370569987d8dee7ff0af832c9e869_5b02ab84772b8aa6c8b343363b0ffaa6bfea2944e7b9cbc2e3b1978ccbd06d4c",
      "label" : "facebook"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/061656d3578811b382af8bdd69a16fee_62291574321b8f336199001d8e10e54ee4458a86e538894f202b6576fd718def",
      "label" : "logo"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/106819aa808d5384f8d5e0bc98b1fee9_b9545927960d95ec29947d828b03d537bb504f941a50b6750bb8cd43f1ac41e1",
      "label" : "logo"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/02648ce6a508d5a8c0f8de9b2ece33fa_9ae25ee4cde2d03e84ab5945d23daa499905338803fbd2248b61732c971965cf",
      "label" : "other"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/586fe1db62023365c05311ba8fa7b22c_b34dbf93f5ac8047198ca438eec750a5db9c1e78283a2e10932c767e312c19d2",
      "label" : "other"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/20dde6be4ba0efcef7eaec9bb66c1156_562b6238f632a9ac28ebacb5ab3346ea5a9b5067db01647765cec5b9bb5fdc8e",
      "label" : "gravatar"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/336243f16faabc7ffe7251da1dd4babb_70eca3858a3f3b30db57e4968ed339f413038fbd3d16107691b1dbc97c629b1c",
      "label" : "foursquare"
    }, {
      "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/8d11d79d3121a7351d0d6e36fe7cb5c0_846c08e187bcf6da90a2a6e12c3958fcc132df72cd531ad0db5306ed2c6703c0",
      "label" : "logo"
    } ],
    "keywords" : [ "Charities", "Developer Tools", "GitHub", "JavaScript", "Linux", "Mobile Applications", "Open Source", "Portland", "Seoul", "Servers", "Services", "Software", "Software Consulting", "Technology Consulting" ]
  },
  "socialProfiles" : [ {
    "typeId" : "facebook",
    "typeName" : "Facebook",
    "url" : "https://www.facebook.com/GitHub"
  }, {
    "bio" : "How people build software",
    "followers" : 649272,
    "following" : 174,
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "https://twitter.com/github",
    "username" : "github",
    "id" : "13334762"
  }, {
    "bio" : "GitHub is the best place to share code with friends, co-workers, classmates, and complete strangers. Over two million people use GitHub to build amazing things together.\n\nWith the collaborative features of GitHub.com, our desktop and mobile apps, and GitHub Enterprise, it has never been easier for individuals and teams to write better code, faster.",
    "followers" : 380,
    "typeId" : "angellist",
    "typeName" : "AngelList",
    "url" : "https://angel.co/github",
    "username" : "github",
    "id" : "60436"
  }, {
    "bio" : "GitHub is a web-based Git repository hosting service offering distributed revision control and source code management functionality of Git.",
    "typeId" : "crunchbasecompany",
    "typeName" : "CrunchBase",
    "url" : "http://www.crunchbase.com/organization/github",
    "username" : "github"
  }, {
    "typeId" : "google",
    "typeName" : "GooglePlus",
    "url" : "https://plus.google.com/+GitHub",
    "username" : "GitHub"
  }, {
    "typeId" : "klout",
    "typeName" : "Klout",
    "url" : "http://klout.com/github",
    "username" : "github",
    "id" : "32369627086608380"
  }, {
    "typeId" : "github",
    "typeName" : "Github",
    "url" : "https://github.com/github",
    "username" : "github"
  }, {
    "bio" : "GitHub is how people build software. With a community of more than 14 million people, developers can discover, use and contribute to over 25 million projects using a powerful, collaborative workflow. Whether using GitHub.com or your own instance of GitHub Enterprise, you can integrate GitHub with third party tools, from project management to continuous deployment, to build software in the way that works best for you.",
    "followers" : 53170,
    "typeId" : "linkedincompany",
    "typeName" : "LinkedIn",
    "url" : "https://www.linkedin.com/company/github",
    "username" : "github",
    "id" : "1418841"
  }, {
    "typeId" : "gravatar",
    "typeName" : "Gravatar",
    "url" : "https://gravatar.com/github",
    "username" : "github",
    "id" : "4024325"
  }, {
    "typeId" : "foursquare",
    "typeName" : "Foursquare",
    "url" : "https://foursquare.com/user/61997196",
    "id" : "61997196"
  } ],
  "traffic" : {
    "topCountryRanking" : [ {
      "rank" : 40,
      "locale" : "us"
    }, {
      "rank" : 52,
      "locale" : "cn"
    }, {
      "rank" : 21,
      "locale" : "jp"
    } ],
    "ranking" : [ {
      "rank" : 64,
      "locale" : "global"
    }, {
      "rank" : 40,
      "locale" : "us"
    } ]
  }
}`
