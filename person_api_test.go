package fullcontact

import (
	"encoding/json"
	"os"
	"testing"
)

func TestPersonEmail(t *testing.T) {
	client, _ := NewClient("")

	_, missingEmail := client.Person.Email("", nil)
	if missingEmail == nil {
		t.Errorf("Person API: Email - Error %v", "Email is not valid. We should have gotten an error")
		return
	}

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}

	email, err := client.Person.Email("bart@fullcontact.com", nil)
	if err != nil {
		t.Errorf("Person API: Email - Error %v", err.Error())
		return
	}

	if email.Status != 200 {
		t.Errorf("Person API: Email - Got non-200 Status %v", email.Status)
		return
	}
}

func TestPersonTwitter(t *testing.T) {
	client, _ := NewClient("")

	_, missingErr := client.Person.Twitter("", nil) //
	if missingErr == nil {
		t.Errorf("Person API: Twitter - %s", "Missing Twitter ID")
		return
	}

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}

	twitter, err := client.Person.Twitter("jack", nil) //
	if err != nil {
		t.Errorf("Person API: Twitter - %s", err.Error())
		return
	}

	if twitter.Status != 200 {
		t.Errorf("Person API: Twitter - Got non-200 Status %v", twitter.Status)
		return
	}

	empty := SocialProfile{}
	if twitter.social("fullcontact") != empty {
		t.Errorf("Person API: Twitter - Should get empty struct")
		return
	}

	socialTwitter := twitter.social("twitter")

	if socialTwitter.ID != "12" {
		t.Errorf("Person API: Twitter - Invalid Twitter ID %s", socialTwitter.ID)
		return
	}

	if socialTwitter.URL != "https://twitter.com/jack" {
		t.Errorf("Person API: Twitter - Invalid Twitter URL %s", socialTwitter.URL)
		return
	}

	if twitter.social("facebook").URL != "https://www.facebook.com/jackdorsey" {
		t.Errorf("Person API: Twitter - Invalid Facebook URL %s", twitter.social("facebook").URL)
		return
	}

}

func TestSocialID(t *testing.T) {

	jsonString := `"abc"`
	var strID SocialID
	if err := json.Unmarshal([]byte(jsonString), &strID); err != nil {
		t.Errorf("SocialID: String: - JSON Error %s", err.Error())
	}
	if strID != "abc" {
		t.Errorf("SocialID: String: Could Not Parse - %s", strID)
		return
	}

	jsonNumeric := `"123"`

	var numID SocialID
	if err := json.Unmarshal([]byte(jsonNumeric), &numID); err != nil {
		t.Errorf("SocialID: Numeric: - JSON Error %s", err.Error())
		return
	}
	if numID != "123" {
		t.Errorf("SocialID: Numeric: Could Not Parse - %s", numID)
		return
	}

	jsonNumericObj := `{"id": 123}`

	var user struct {
		ID SocialID `json:"id"`
	}

	if err := json.Unmarshal([]byte(jsonNumericObj), &user); err != nil {
		t.Errorf("SocialID: Numeric: - JSON Error %s", err.Error())
		return
	}
	if user.ID != "123" {
		t.Errorf("SocialID: Numeric: Could Not Parse - %s", user.ID)
		return
	}

	var person struct {
		ID SocialID `json:"id"`
	}
	person.ID = "abc-dec"

	personJSON, err := json.Marshal(person)
	if err != nil {
		t.Errorf("SocialID: Object: - JSON Error %s", err.Error())
		return
	}
	if string(personJSON) != `{"id":"abc-dec"}` {
		t.Errorf("SocialID: Object: Could Not Parse - %s", string(personJSON))
		return
	}
}

func TestPersonPhone(t *testing.T) {
	client, _ := NewClient("")

	if os.Getenv("FULLCONTACT_API") == "" {
		t.Skip("FULLCONTACT_API environment variable not set.")
	}

	phone, err := client.Person.Phone("888-330-6943", nil) // FullContact Phone
	if err != nil {
		t.Errorf("Person API: Phone - %s", err.Error())
		return
	}

	if phone.Status != 200 {
		t.Errorf("Person API: Phone - Got non-200 Status %v", phone.Status)
		return
	}

}

func TestPersonPhoneInvalid(t *testing.T) {
	client, _ := NewClient("")

	phone, err := client.Person.Phone("8883306943", nil)
	if err == nil || phone != nil {
		t.Errorf("Person API: Phone Should be Invalid Format")
	}

}

var jsonPersonEmail = `{
  "status" : 200,
  "requestId" : "73df8546-f876-4d09-b4b1-adce2c84ccfb",
  "likelihood" : 0.99,
  "photos" : [ {
    "type" : "foursquare",
    "typeId" : "foursquare",
    "typeName" : "Foursquare",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/a7e6a5aba590d4933e35eaadabd97fd2_44e00e968ac57725a15b32f9ca714827aff8e4818d290cb0c611f2e2585253b3",
    "isPrimary" : true
  } ],
  "contactInfo" : {
    "chats" : [ {
      "client" : "gtalk",
      "handle" : "lorangb@gmail.com"
    }, {
      "client" : "skype",
      "handle" : "bart.lorang"
    } ],
    "websites" : [ {
      "url" : "https://fullcontact.com/bart"
    }, {
      "url" : "http://bartlorang.com"
    } ],
    "familyName" : "Lorang",
    "fullName" : "Bart Lorang",
    "givenName" : "Bart"
  },
  "organizations" : [ {
    "name" : "V1.vc",
    "startDate" : "2015-07",
    "title" : "Co-Founder & Managing Director",
    "current" : true
  } ],
  "demographics" : {
    "locationDeduced" : {
      "normalizedLocation" : "Boulder, Colorado, United States",
      "deducedLocation" : "Boulder, Colorado, United States",
      "city" : {
        "name" : "Boulder"
      },
      "state" : {
        "name" : "Colorado",
        "code" : "CO"
      },
      "country" : {
        "name" : "United States",
        "code" : "US"
      },
      "continent" : {
        "deduced" : true,
        "name" : "North America"
      },
      "county" : {
        "deduced" : true,
        "name" : "Boulder"
      },
      "likelihood" : 1.0
    },
    "age" : "37",
    "ageRange" : "32-42",
    "gender" : "Male",
    "locationGeneral" : "Boulder, Colorado, United States"
  },
  "socialProfiles" : [ {
    "bio" : "Co-Founder and CEO of FullContact",
    "type" : "aboutme",
    "typeId" : "aboutme",
    "typeName" : "About.me",
    "url" : "https://about.me/lorangb",
    "username" : "lorangb"
  }, {
    "bio" : "Entrepreneur, Tech Nerd; CEO and Co-Founder of @fullcontact. Passionate about solving the world's contact information problem and helping entrepreneurs.",
    "followers" : 2919,
    "type" : "angellist",
    "typeId" : "angellist",
    "typeName" : "AngelList",
    "url" : "https://angel.co/bartlorang",
    "username" : "bartlorang",
    "id" : "182"
  }, {
    "type" : "facebook",
    "typeId" : "facebook",
    "typeName" : "Facebook",
    "url" : "https://www.facebook.com/bart.lorang"
  }, {
    "type" : "flickr",
    "typeId" : "flickr",
    "typeName" : "Flickr",
    "url" : "https://www.flickr.com/people/39267654@N00",
    "username" : "39267654@n00",
    "id" : "39267654@N00"
  }, {
    "type" : "github",
    "typeId" : "github",
    "typeName" : "Github",
    "url" : "https://github.com/lorangb",
    "username" : "lorangb"
  }, {
    "followers" : 1,
    "type" : "google",
    "typeId" : "google",
    "typeName" : "GooglePlus",
    "url" : "https://plus.google.com/111748526539078793602",
    "id" : "111748526539078793602"
  }, {
    "bio" : "http://about.me/lorangb",
    "type" : "gravatar",
    "typeId" : "gravatar",
    "typeName" : "Gravatar",
    "url" : "https://gravatar.com/blorang",
    "username" : "blorang",
    "id" : "18197740"
  }, {
    "type" : "hackernews",
    "typeId" : "hackernews",
    "typeName" : "HackerNews",
    "url" : "http://news.ycombinator.com/user?id=lorangb",
    "username" : "lorangb"
  }, {
    "type" : "instagram",
    "typeId" : "instagram",
    "typeName" : "Instagram",
    "url" : "https://instagram.com/bartlorang"
  }, {
    "bio" : "CEO & Founder of FullContactManaging Director of v1.vcTech Entrepreneur, Investor",
    "type" : "keybase",
    "typeId" : "keybase",
    "typeName" : "Keybase",
    "url" : "https://keybase.io/bartlorang",
    "username" : "bartlorang",
    "id" : "b4efc8f483638567f42dca0561caa319"
  }, {
    "bio" : "Mr. Lorang is a proven entrepreneur, executive and manager in the global technology industry. Mr. Lorang is active in the startup technology community as an angel investor, strategic advisor and speaker at industry events. Bart serves as Co-Founder & CEO of FullContact. Mr. Lorang is responsible for communicating FullContact's vision and strategy. Mr. Lorang is a visionary technologist with extensive experience conceiving, designing, building, marketing and selling enterprise software solutions on a global scale. Bart is also Co-Founder and Managing Director of v1.vc, a $5M seed stage fund dedicated to helping crazy entrepreneurs change the world. Bart serves on the Board of the Colorado Technology Association, Rapt Media and is on the Advisory Board of Education Funding Partners. Bart is a regular guest on FOX Business channel and has been featured by ABC, CNN, FOX News, MSNBC, Forbes, FastCompany, Yahoo, Inc Magazine and TechCrunch. Prior to founding FullContact, Mr. Lorang was an owner in Dimension Technology Solutions where he served as President and oversaw all day to day operations, customer engagements, partner relations, product development, sales and marketing functions. Mr. Lorang is recognized for providing solutions that are simple and work reliably. He strongly believes in using technology to solve problems as opposed to using problems to demonstrate technology. Mr. Lorang holds a Bachelor of Science degree in Computer Science from the University of Colorado and an MBA from the Daniels College of Business at University of Denver. Specialties: Investments, Startups, Financial Analysis, Sales, Technical Sales, Implementations, System Integration, Project Management, Leadership",
    "followers" : 500,
    "following" : 500,
    "type" : "linkedin",
    "typeId" : "linkedin",
    "typeName" : "LinkedIn",
    "url" : "https://www.linkedin.com/in/bartlorang",
    "username" : "bartlorang",
    "id" : "8995706"
  }, {
    "followers" : 89,
    "following" : 26,
    "type" : "pinterest",
    "typeId" : "pinterest",
    "typeName" : "Pinterest",
    "url" : "http://www.pinterest.com/lorangb/",
    "username" : "lorangb"
  }, {
    "type" : "plancast",
    "typeId" : "plancast",
    "typeName" : "Plancast",
    "url" : "http://www.plancast.com/lorangb",
    "username" : "lorangb"
  }, {
    "type" : "quora",
    "typeId" : "quora",
    "typeName" : "Quora",
    "url" : "http://www.quora.com/bart-lorang",
    "username" : "bart-lorang"
  }, {
    "bio" : "CEO & Co-Founder of @FullContact, Managing Director @v1vc_. Tech Entrepreneur, Investor. Husband to @parkerbenson and Father to Greyson Lorang",
    "followers" : 5454,
    "following" : 741,
    "type" : "twitter",
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "https://twitter.com/bartlorang",
    "username" : "bartlorang",
    "id" : "5998422"
  }, {
    "type" : "xing",
    "typeId" : "xing",
    "typeName" : "Xing",
    "url" : "https://www.xing.com/profile/bart_lorang2",
    "username" : "bart_lorang2"
  }, {
    "type" : "youtube",
    "typeId" : "youtube",
    "typeName" : "YouTube",
    "url" : "https://youtube.com/user/lorangb",
    "username" : "lorangb"
  } ],
  "digitalFootprint" : {
    "topics" : [ {
      "provider" : "aboutme",
      "value" : "Angel Investor"
    }, {
      "provider" : "aboutme",
      "value" : "Entrepreneur"
    }, {
      "provider" : "aboutme",
      "value" : "Husband"
    }, {
      "provider" : "aboutme",
      "value" : "Tech Nerd"
    }, {
      "provider" : "aboutme",
      "value" : "Technology"
    } ]
  }
}`

var jsonPersonTwitter = `{
  "status" : 200,
  "requestId" : "8ee78d86-dac9-41a3-a7e3-88430f905522",
  "likelihood" : 0.95,
  "photos" : [ {
    "typeId" : "userclaim",
    "typeName" : "userClaim",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/edaa53d9a080aea37ddfb85d775620a9_98a2d7beef6a5b4a53f43da4dd1a90bda21dc18f755394fdbf9b6cf3283853a0",
    "isPrimary" : true
  }, {
    "typeId" : "userclaim",
    "typeName" : "userClaim",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/922c5b3acda4a1b5ea3d44943f8c033a_45183f024950993c27e5b6c4abc8a4a4d62df77197cfd8a7194d03b6b35c2915"
  } ],
  "contactInfo" : {
    "websites" : [ {
      "url" : "https://www.fullcontact.com"
    } ]
  },
  "organizations" : [ {
    "isPrimary" : true,
    "name" : "FullContact Inc.",
    "current" : true
  } ],
  "socialProfiles" : [ {
    "typeId" : "googleplus",
    "typeName" : "Google Plus",
    "url" : "https://plus.google.com/u/0/107620035082673219790",
    "id" : "u/0/107620035082673219790"
  }, {
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "http://twitter.com/fullcontact",
    "username" : "fullcontact"
  }, {
    "typeId" : "gravatar",
    "typeName" : "Gravatar",
    "url" : "http://gravatar.com/fullcontactinfo",
    "username" : "fullcontactinfo"
  } ]
}`

var jsonPersonPhone = `{
  "status" : 200,
  "requestId" : "80a6fd54-2cae-4a3b-ac0e-5663b212d106",
  "likelihood" : 0.95,
  "photos" : [ {
    "type" : "twitter",
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/2a4f9c5c54331c4e1148a7e799279a1a_2245a77eba2970b8f965f0f6d9aa3adc36617847114d4fd80c46138e53c975d5",
    "isPrimary" : true
  }, {
    "type" : "other",
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/267aa47cf473e6123c4268ac1ec98cd7_03bf0b911b1ca9ea84662eca95a9cf113e3786002f576e128387ef237869a3af"
  }, {
    "type" : "angellist",
    "typeId" : "angellist",
    "typeName" : "AngelList",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/868904add6ce778444589f010693949e_5e3c361b611b7d54cbacf923e35ae63d4b552087ae56699a0a673fb176cc05d5"
  }, {
    "type" : "google",
    "typeId" : "google",
    "typeName" : "GooglePlus",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/930121c9557ceefb051403d801882b7b_9dec2397963b6fe85b6e46b4c8e1f8b73bfe4de22d057a2f95c5ab194cce9550"
  }, {
    "type" : "instagram",
    "typeId" : "instagram",
    "typeName" : "Instagram",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/768c512b3ed4fe2e5e97aa3532f5b318_680b880e8682ddf1a57d7946b99053bca65b12f8cccec34f8c8eec1cd525006a"
  }, {
    "type" : "pinterest",
    "typeId" : "pinterest",
    "typeName" : "Pinterest",
    "url" : "https://d2ojpxxtu63wzl.cloudfront.net/static/fe03b65c8164e0b2e5185f44e684b858_72891030873bc95d85fa08b3d3ffd5adef2fd0bb49f7ebb8149a0227441d0179"
  } ],
  "contactInfo" : {
    "websites" : [ {
      "url" : "http://www.fullcontact.com"
    }, {
      "url" : "http://www.kutenda.com"
    }, {
      "url" : "http://www.kutendaformsps.com"
    } ],
    "familyName" : "Vojslavek",
    "fullName" : "Andrew Vojslavek",
    "givenName" : "Andrew"
  },
  "organizations" : [ {
    "name" : "FullContact Inc.",
    "startDate" : "2013-06",
    "title" : "Director Channels and Partnerships",
    "current" : true
  }, {
    "name" : "FullContact",
    "title" : "Account Executive"
  } ],
  "demographics" : {
    "locationDeduced" : {
      "normalizedLocation" : "Denver, Colorado, United States",
      "deducedLocation" : "Denver, Colorado, United States",
      "city" : {
        "name" : "Denver"
      },
      "state" : {
        "name" : "Colorado",
        "code" : "CO"
      },
      "country" : {
        "name" : "United States",
        "code" : "US"
      },
      "continent" : {
        "deduced" : true,
        "name" : "North America"
      },
      "county" : {
        "deduced" : true,
        "name" : "Denver"
      },
      "likelihood" : 1.0
    },
    "ageRange" : "25-35",
    "gender" : "Male",
    "locationGeneral" : "Denver, Colorado, United States"
  },
  "socialProfiles" : [ {
    "bio" : "<p>Married to the best person ever. Keeping myself busy by selling APIs, bouldering, running, and drinking amazing beer.</p>",
    "type" : "aboutme",
    "typeId" : "aboutme",
    "typeName" : "About.me",
    "url" : "https://about.me/avojslavek",
    "username" : "avojslavek"
  }, {
    "followers" : 7,
    "type" : "angellist",
    "typeId" : "angellist",
    "typeName" : "AngelList",
    "url" : "https://angel.co/andrew-vojslavek",
    "username" : "andrew-vojslavek",
    "id" : "795207"
  }, {
    "type" : "facebook",
    "typeId" : "facebook",
    "typeName" : "Facebook",
    "url" : "https://www.facebook.com/andrew.vojslavek"
  }, {
    "type" : "foursquare",
    "typeId" : "foursquare",
    "typeName" : "Foursquare",
    "url" : "https://foursquare.com/user/2584215",
    "id" : "2584215"
  }, {
    "followers" : 20,
    "type" : "google",
    "typeId" : "google",
    "typeName" : "GooglePlus",
    "url" : "https://plus.google.com/100215368740929270230",
    "id" : "100215368740929270230"
  }, {
    "type" : "instagram",
    "typeId" : "instagram",
    "typeName" : "Instagram",
    "url" : "https://instagram.com/coarv86"
  }, {
    "type" : "klout",
    "typeId" : "klout",
    "typeName" : "Klout",
    "url" : "http://klout.com/avojslavek",
    "username" : "avojslavek",
    "id" : "79375950882299374"
  }, {
    "bio" : "Specialties: cold calling, consultative sales approach, question based selling, relationship building, closing new business, account management.",
    "followers" : 500,
    "following" : 500,
    "type" : "linkedin",
    "typeId" : "linkedin",
    "typeName" : "LinkedIn",
    "url" : "https://www.linkedin.com/in/andrewvojslavek",
    "username" : "andrewvojslavek",
    "id" : "70769536"
  }, {
    "followers" : 43,
    "following" : 23,
    "type" : "pinterest",
    "typeId" : "pinterest",
    "typeName" : "Pinterest",
    "url" : "http://www.pinterest.com/avoj/",
    "username" : "avoj"
  }, {
    "bio" : "Life is about... My wife, growing @FullContactInc, Rock Climbing, and lots of reading!",
    "followers" : 178,
    "following" : 466,
    "type" : "twitter",
    "typeId" : "twitter",
    "typeName" : "Twitter",
    "url" : "https://twitter.com/avojslavek",
    "username" : "avojslavek",
    "id" : "473146231"
  }, {
    "type" : "vimeo",
    "typeId" : "vimeo",
    "typeName" : "Vimeo",
    "url" : "http://vimeo.com/user1059183",
    "username" : "user1059183",
    "id" : "1059183"
  } ],
  "digitalFootprint" : {
    "scores" : [ {
      "provider" : "klout",
      "type" : "general",
      "value" : 35
    } ],
    "topics" : [ {
      "provider" : "aboutme",
      "value" : "Apis"
    }, {
      "provider" : "aboutme",
      "value" : "Beer"
    }, {
      "provider" : "aboutme",
      "value" : "Bouldering"
    }, {
      "provider" : "aboutme",
      "value" : "Running"
    }, {
      "provider" : "aboutme",
      "value" : "Start Ups"
    }, {
      "provider" : "klout",
      "value" : "Gmail"
    }, {
      "provider" : "klout",
      "value" : "Sales"
    } ]
  }
}`
