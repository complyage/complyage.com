package agent_interfaces

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Allowed Models
//||------------------------------------------------------------------------------------------------||

const (
	ModelFace   = "face"
	ModelNSFW   = "nsfw"
	ModelOCR    = "ocr"
	ModelVision = "vision"
)

//||------------------------------------------------------------------------------------------------||
//|| AllowAgentModels
//||------------------------------------------------------------------------------------------------||

var AllowAgentModels = []string{
	ModelFace,
	ModelNSFW,
	ModelOCR,
	ModelVision,
}

//||------------------------------------------------------------------------------------------------||
//|| Agent Actions
//||------------------------------------------------------------------------------------------------||

const (
	AgentActionFaceCompare   = "face.compare"
	AgentActionFaceDetect    = "face.detect"
	AgentActionNSFWDetect    = "nsfw.detect"
	AgentActionOCRExtract    = "ocr.extract"
	AgentActionVisionAnalyze = "vision.analyze"
	AgentActionVisionContent = "vision.content"
)

//||------------------------------------------------------------------------------------------------||
//|| AllowedAgentActions
//||------------------------------------------------------------------------------------------------||

var AllowedAgentActions = []string{
	AgentActionFaceCompare,
	AgentActionFaceDetect,
	AgentActionNSFWDetect,
	AgentActionOCRExtract,
	AgentActionVisionAnalyze,
	AgentActionVisionContent,
}

//||------------------------------------------------------------------------------------------------||
//|| Media
//||------------------------------------------------------------------------------------------------||

type AgentMedia struct {
	Blob   string `json:"blob,omitempty"`
	Base64 string `json:"base64,omitempty"`
	Mime   string `json:"type"`
}

//||------------------------------------------------------------------------------------------------||
//|| AgentRequest
//||------------------------------------------------------------------------------------------------||

type AgentRequest struct {
	Identity  string                 `json:"identity,omitempty"`
	Model     string                 `json:"model,omitempty"`
	Action    string                 `json:"action"`
	Params    map[string]interface{} `json:"params,omitempty"`
	Media     []AgentMedia           `json:"media,omitempty"`
	Prompt    string                 `json:"prompt,omitempty"`
	CallBack  string                 `json:"callback,omitempty"`
	Timestamp string                 `json:"timestamp,omitempty"`
}
