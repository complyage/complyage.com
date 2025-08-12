package agent_helpers

import "agent/agent_interfaces"

// type APIAgentRespondProfile struct {
// 	Identity   string                 `json:"identity,omitempty"` // Unique identifier for the agent
// 	Params     map[string]interface{} `json:"params,omitempty"`
// 	Front      AgentMedia             `json:"front,omitempty"`
// 	Back       AgentMedia             `json:"back,omitempty"`
// 	Profile    AgentMedia             `json:"profile,omitempty"`
// 	FinalLevel int                    `json:"level,omitempty"`      // 0 = Base AI, 1 = Vision AI, 3 = Human
// 	MaxLevel   int                    `json:"max_level,omitempty"`  // 0 = Base AI, 1 = Vision AI, 3 = Human
// 	Status     string                 `json:"status,omitempty"`     // Status of the request (e.g., "pending", "completed", "failed")
// 	Message    string                 `json:"message,omitempty"`    // Message related to the request
// 	ErrorCode  string                 `json:"error_code,omitempty"` // Error code if any error occurred
// }

func CallbackProfile(status string, errorCode string, response agent_interfaces.AgentResponse) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	callbackProfile := agent_interfaces.APIAgentRequestProfile{
		Identity:  response.Request.Identity,
		Params:    response.Request.Params,
		Front:     response.Request.Front,
		Back:      response.Request.Back,
		Profile:   response.Request.Profile,
		Level:     response.Request.Level,
		MaxLevel:  response.Request.MaxLevel,
		Status:    status,
		Message:   response.Message,
		ErrorCode: errorCode,
	}

	if agent_helpers.Elapsed > 0 {
		callbackData["elapsed"] = agent_helpers.Elapsed
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Send Callback
	//||------------------------------------------------------------------------------------------------||

	return db.SendCallback(session.CallBack, callbackData)
}
