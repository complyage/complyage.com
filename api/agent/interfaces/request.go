package interfaces

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type AgentRequest struct {
	Identifier string       `json:"identifier"`
	Process    string       `json:"process,omitempty"`
	Media      []AgentMedia `json:"media,omitempty"`
	CallBack   string       `json:"callback,omitempty"`
	Timestamp  string       `json:"timestamp,omitempty"`
}
