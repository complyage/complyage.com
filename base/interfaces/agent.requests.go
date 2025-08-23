package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Fatal Response
//||------------------------------------------------------------------------------------------------||

type AgentFatalRequest struct {
	Identifier string `json:"identifier"`
	ErrorCode  string `json:"errorCode"`
	Message    string `json:"message"`
	Process    string `json:"process,omitempty"`
}
