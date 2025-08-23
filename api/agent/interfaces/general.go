package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Media
//||------------------------------------------------------------------------------------------------||

type AgentMedia struct {
	Blob   string `json:"blob,omitempty"`
	Base64 string `json:"base64,omitempty"`
	Mime   string `json:"type"`
}
