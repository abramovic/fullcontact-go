package fullcontact

import (
	"fmt"
)

var (
	libraryName  = "FullContact-Go"
	errUnknown   = fmt.Errorf("%s: Error unknown", libraryName)
	errLibrary   = fmt.Errorf("%s: Library-specific error", libraryName)
	ErrStatus400 = fmt.Errorf("%s: Your request was malformed", libraryName)
	ErrStatus403 = fmt.Errorf("%s: Your API key is invalid, missing, or has exceeded its quota. **Paid plans will not receive a 403 response when they exceed their alotted matches. They will only receive a 403 for exceeding rate limit quotas", libraryName)
	ErrStatus404 = fmt.Errorf("%s: This person was searched in the past 24 hours and nothing was found", libraryName)
	ErrStatus405 = fmt.Errorf("%s: You have queried the API with an unsupported HTTP method. Retry your query with either GET or POST", libraryName)
	ErrStatus422 = fmt.Errorf("%s: Invalid or missing API query parameter", libraryName)
	ErrStatus500 = fmt.Errorf("%s: There was an unexpected error on our server. If you see this please contact support@fullcontact.com", libraryName)
)
