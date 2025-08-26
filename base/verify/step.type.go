package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| StepType (iota-based enum with code/description)
//||------------------------------------------------------------------------------------------------||

type StepType int

const (
	StepInitial StepType = iota
	StepStatusChange
	StepAgentLevel1
	StepModerate
	StepSentEmail
	StepSentSMS
	StepPayProcess
	StepCodeEntry
	StepFaceMatch
	StepFaceLive
	StepMediaUpload
	StepOCR
	StepDOB
	StepAgeVerified
	StepComplete
)

//||------------------------------------------------------------------------------------------------||
//|| StepType String, Code, Description
//||------------------------------------------------------------------------------------------------||

func (t StepType) Code() string {
	switch t {
	case StepInitial:
		return "INITIAL"
	case StepStatusChange:
		return "STATUS_CHANGE"
	case StepAgentLevel1:
		return "STEP_AGENT_L1"
	case StepModerate:
		return "MODERATE"
	case StepSentEmail:
		return "SENT_EMAIL"
	case StepSentSMS:
		return "SENT_SMS"
	case StepPayProcess:
		return "PAY_PROCESS"
	case StepCodeEntry:
		return "CODE_ENTRY"
	case StepFaceMatch:
		return "FACE_MATCH"
	case StepFaceLive:
		return "FACE_LIVE"
	case StepMediaUpload:
		return "MEDIA_UPLOAD"
	case StepOCR:
		return "OCR"
	case StepDOB:
		return "DOB"
	case StepAgeVerified:
		return "AGE_VERIFIED"
	case StepComplete:
		return "COMPLETE"
	default:
		return "UNKNOWN"
	}
}

func (t StepType) Description(more string) string {
	switch t {
	case StepInitial:
		return "Initial verification step"
	case StepStatusChange:
		return "Status changed to " + more
	case StepAgentLevel1:
		return "Agent Level 1 processing"
	case StepModerate:
		return "Moderation step: " + more
	case StepSentEmail:
		return "Email sent to user"
	case StepSentSMS:
		return "SMS sent to user"
	case StepPayProcess:
		return "Payment processed"
	case StepCodeEntry:
		return "Code entry verification"
	case StepFaceMatch:
		return "Face matching verification"
	case StepFaceLive:
		return "Live face verification"
	case StepMediaUpload:
		return "Media uploaded"
	case StepOCR:
		return "Optical Character Recognition performed"
	case StepDOB:
		return "DOB verified (" + more + ")"
	case StepAgeVerified:
		return "Age verified (" + more + ")"
	case StepComplete:
		return "Verification process completed"
	default:
		return "Unknown step"
	}
}

func (t StepType) String() string {
	return t.Code()
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (t StepType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.Code())
}

func (t *StepType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "INITIAL":
		*t = StepInitial
	case "STATUS_CHANGE":
		*t = StepStatusChange
	case "STEP_AGENT_L1":
		*t = StepAgentLevel1
	case "MODERATE":
		*t = StepModerate
	case "SENT_EMAIL":
		*t = StepSentEmail
	case "SENT_SMS":
		*t = StepSentSMS
	case "PAY_PROCESS":
		*t = StepPayProcess
	case "CODE_ENTRY":
		*t = StepCodeEntry
	case "FACE_MATCH":
		*t = StepFaceMatch
	case "FACE_LIVE":
		*t = StepFaceLive
	case "MEDIA_UPLOAD":
		*t = StepMediaUpload
	case "OCR":
		*t = StepOCR
	case "DOB":
		*t = StepDOB
	case "AGE_VERIFIED":
		*t = StepAgeVerified
	case "COMPLETE":
		*t = StepComplete
	default:
		return fmt.Errorf("invalid StepType: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Dot notation namespace for StepType
//||------------------------------------------------------------------------------------------------||

type nsStepType struct {
	Initial         StepType
	StatusChange    StepType
	StepAgentLevel1 StepType
	Moderate        StepType
	SentEmail       StepType
	SentSMS         StepType
	PayProcess      StepType
	CodeEntry       StepType
	FaceMatch       StepType
	FaceLive        StepType
	MediaUpload     StepType
	OCR             StepType
	DOB             StepType
	AgeVerified     StepType
	Complete        StepType
}

var STEPTYPES = nsStepType{
	Initial:         StepInitial,
	StatusChange:    StepStatusChange,
	StepAgentLevel1: StepAgentLevel1,
	Moderate:        StepModerate,
	SentEmail:       StepSentEmail,
	SentSMS:         StepSentSMS,
	PayProcess:      StepPayProcess,
	CodeEntry:       StepCodeEntry,
	FaceMatch:       StepFaceMatch,
	FaceLive:        StepFaceLive,
	MediaUpload:     StepMediaUpload,
	OCR:             StepOCR,
	DOB:             StepDOB,
	AgeVerified:     StepAgeVerified,
	Complete:        StepComplete,
}
