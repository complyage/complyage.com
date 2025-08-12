package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type AgentRequestProfile struct {
	SiteId    uint                   `json:"agent_key,omitempty"`
	Identity  string                 `json:"identity,omitempty"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Front     AgentMedia             `json:"front,omitempty"`
	Back      AgentMedia             `json:"back,omitempty"`
	Profile   AgentMedia             `json:"profile,omitempty"`
	Level     int                    `json:"level,omitempty"`     // 0 = Base AI, 1 = Vision AI, 3 = Human
	MaxLevel  int                    `json:"max_level,omitempty"` // 0 = Base AI, 1 = Vision AI, 3 = Human
	CallBack  string                 `json:"callback,omitempty"`
	Timestamp time.Time              `json:"timestamp,omitempty"`
	Elapsed   int64                  `json:"elapsed,omitempty"`
	Process   IdentificationProcess  `json:"process,omitempty"` // Process for the identification
}
