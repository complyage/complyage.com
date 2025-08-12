package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type APIAgentRequestProfile struct {
	AgentKey string                 `json:"agent_key,omitempty"`
	Params   map[string]interface{} `json:"params,omitempty"`
	Front    AgentMedia             `json:"front,omitempty"`
	Back     AgentMedia             `json:"back,omitempty"`
	Profile  AgentMedia             `json:"profile,omitempty"`
	Level    int                    `json:"level,omitempty"`     // 0 = Base AI, 1 = Vision AI, 3 = Human
	MaxLevel int                    `json:"max_level,omitempty"` // 0 = Base AI, 1 = Vision AI, 3 = Human
	CallBack string                 `json:"callback,omitempty"`
}
