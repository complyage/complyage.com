package interfaces

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type AgentRequestVerifyID struct {
	UUID   string     `json:"uuid"`
	Front  AgentMedia `json:"front"`
	Back   AgentMedia `json:"back"`
	Selfie AgentMedia `json:"selfie"`
}
