package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| AgentMedia
//||------------------------------------------------------------------------------------------------||

type AgentProcessMedia struct {
	Filename string `json:"filename,omitempty"`  // Name of the media file
	MimeType string `json:"mime_type,omitempty"` // MIME type of the media file
	FileSize int    `json:"file_size,omitempty"` // Size of the media file in bytes
	Blob     string `json:"blob,omitempty"`      // Base64 encoded media content
}

//||------------------------------------------------------------------------------------------------||
//|| AgentProcessGlobal
//||------------------------------------------------------------------------------------------------||

type AgentProcessGlobalStep struct {
	StepNumber      int                    `json:"step_number,omitempty"`      // Step number in the process
	StepDescription string                 `json:"step_description,omitempty"` // Description of the step
	ProcessCode     string                 `json:"process_code,omitempty"`     // Unique code for the process
	Success         bool                   `json:"success,omitempty"`          // Indicates if the step was successful
	ErrorCode       string                 `json:"error_code,omitempty"`       // Error code if any error occurred
	Elapsed         int                    `json:"elapsed,omitempty"`          // Time taken for the step in milliseconds
	Debug           map[string]interface{} `json:"debug,omitempty"`            // Debug information for the step
	Data            map[string]interface{} `json:"data,omitempty"`             // Data
}

//||------------------------------------------------------------------------------------------------||
//|| AgentProcessGlobalMedia
//||------------------------------------------------------------------------------------------------||

type AgentProcessGlobalMediaOverview struct {
	Front   AgentMedia `json:"front,omitempty"`   // Overview of the front media
	Back    AgentMedia `json:"back,omitempty"`    // Overview of the back media
	Profile AgentMedia `json:"profile,omitempty"` // Overview of the profile media
	Other   AgentMedia `json:"other,omitempty"`   // Overview of any other media
}

//||------------------------------------------------------------------------------------------------||
//|| AgentProcessGlobal
//||------------------------------------------------------------------------------------------------||

type AgentProcessGlobal struct {
	AgentKey    string                          `json:"agent_key,omitempty"`
	Identity    string                          `json:"identity,omitempty"` // Unique identifier for the agent
	Process     string                          `json:"process,omitempty"`  // Name of the process being executed
	Params      map[string]interface{}          `json:"params,omitempty"`
	Media       AgentProcessGlobalMediaOverview `json:"media,omitempty"`     // Overview of the media involved in the process
	Level       int                             `json:"level,omitempty"`     // 0 = Base AI, 1 = Vision AI, 3 = Human
	MaxLevel    int                             `json:"max_level,omitempty"` // 0 = Base AI, 1 = Vision AI, 3 = Human
	CallBack    string                          `json:"callback,omitempty"`
	Status      string                          `json:"status,omitempty"`       // Status of the request (e.g., "pending", "completed", "failed")
	Message     string                          `json:"message,omitempty"`      // Message related to the request
	ErrorCode   string                          `json:"error_code,omitempty"`   // Error code if any error occurred
	CurrentStep int                             `json:"current_step,omitempty"` // Current step in the process
	Steps       []AgentProcessGlobalStep        `json:"steps,omitempty"`        // Steps in the process
}
