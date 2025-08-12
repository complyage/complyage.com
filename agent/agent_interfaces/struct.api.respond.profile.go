package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| APIAgentRespondProfile
//||------------------------------------------------------------------------------------------------||

type APIAgentRespondProfile struct {
	Identity   string                 `json:"identity,omitempty"` // Unique identifier for the agent
	Params     map[string]interface{} `json:"params,omitempty"`
	Front      AgentMedia             `json:"front,omitempty"`
	Back       AgentMedia             `json:"back,omitempty"`
	Profile    AgentMedia             `json:"profile,omitempty"`
	FinalLevel int                    `json:"level,omitempty"`      // 0 = Base AI, 1 = Vision AI, 3 = Human
	MaxLevel   int                    `json:"max_level,omitempty"`  // 0 = Base AI, 1 = Vision AI, 3 = Human
	Status     string                 `json:"status,omitempty"`     // Status of the request (e.g., "pending", "completed", "failed")
	Message    string                 `json:"message,omitempty"`    // Message related to the request
	ErrorCode  string                 `json:"error_code,omitempty"` // Error code if any error occurred
}
