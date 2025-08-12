package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| AgentMessage
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Agent Response
//||------------------------------------------------------------------------------------------------||

type AgentResponse struct {
	Success bool                   `json:"success"`
	Message string                 `json:"message"` // ✅ now properly matches agent output
	Details map[string]interface{} `json:"details"`
	Request AgentRequest           `json:"request"`
	Data    map[string]interface{} `json:"data"`
	Elapsed int                    `json:"elapsed,omitempty"`
}
