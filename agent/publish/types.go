package publish

//||------------------------------------------------------------------------------------------------||
//|| Processes
//||------------------------------------------------------------------------------------------------||

type ProcessTypes string

const (
	ProcessVerifyID  ProcessTypes = "verify_id"
	ProcessFacialAge ProcessTypes = "facial_age"
)

//||------------------------------------------------------------------------------------------------||
//|| AgentVerification
//||------------------------------------------------------------------------------------------------||

type AgentVerification struct {
	Level      int          `json:"level"`
	Process    ProcessTypes `json:"process"`
	Identifier string       `json:"identifier"`
	Timestamp  string       `json:"timestamp,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type AgentRequest struct {
	Level     int          `json:"level"`
	Model     string       `json:"identifier"`
	Prompt    string       `json:"process,omitempty"`
	Media     []AgentMedia `json:"media,omitempty"`
	CallBack  string       `json:"callback,omitempty"`
	Timestamp string       `json:"timestamp,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Media
//||------------------------------------------------------------------------------------------------||

type AgentMedia struct {
	Base64 string `json:"base64,omitempty"`
	Mime   string `json:"type"`
}
