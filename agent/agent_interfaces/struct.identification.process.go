package agent_interfaces

import (
	"base/interfaces"
)

type IdentificationProcess struct {
	Step           int                               `json:"step"`
	RawText        string                            `json:"raw_text,omitempty"`
	FaceMatch      bool                              `json:"face_match,omitempty"`
	IDVerified     bool                              `json:"address_read,omitempty"`
	Identification interfaces.VerifiedIdentification `json:"identification,omitempty"`
	Error          string                            `json:"error,omitempty"`
}
