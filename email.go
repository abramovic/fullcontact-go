package fullcontact

// EmailResponse is the response from FullContact's Email API
/*
{
     "status": {"type":"number"},  // common status codes can be found below
     "usernameSubAddressing": {"type":"string"},  // true, false, or unknown
     "disposableEmailDomain": {"type":"string"},  // true, false, or retry
     "message": {"type":"string"}
}
*/
type EmailResponse struct {
	Status           int64  `json:"status"`
	SubAddressing    string `json:"usernameSubAddressing"` // true, false, unknown or retry
	DisposableDomain string `json:"disposableEmailDomain"` // true, false, or retry
	Message          string `json:"message"`
}
