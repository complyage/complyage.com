package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Meta Data
//||------------------------------------------------------------------------------------------------||

type VerificationMeta struct {
	Approval VerificationMetaApproval `json:"approval"`
	Step     int                      `json:"step,omitempty"`
	Steps    []VerificationMetaStep   `json:"steps,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Approval Data
//||------------------------------------------------------------------------------------------------||

type VerificationMetaApproval struct {
	ApprovedBy     string `json:"approvedBy"`
	ApprovedAt     string `json:"approvedAt"`
	ApprovedMethod string `json:"approvedMethod"`
}

//||------------------------------------------------------------------------------------------------||
//|| Approval Data
//||------------------------------------------------------------------------------------------------||

type VerificationMetaStep struct {
	StepName      string `json:"stepName"`
	StepStatus    string `json:"stepStatus"`            // e.g., "pending", "completed", "failed"
	StepDetails   string `json:"stepDetails,omitempty"` // Additional information about the step
	StepTimestamp string `json:"stepTimestamp"`         // Timestamp of when the step was processed
}
