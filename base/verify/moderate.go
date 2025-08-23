package verify

//||------------------------------------------------------------------------------------------------||
//|| Approval Struct
//||------------------------------------------------------------------------------------------------||

type Moderate struct {
	Type      ModerateType   `json:"type"`
	Status    ModerateStatus `json:"status"`
	Moderator int64          `json:"moderator,omitempty"`
	Timestamp string         `json:"timestamp,omitempty"`
	Details   string         `json:"details,omitempty"`
}
