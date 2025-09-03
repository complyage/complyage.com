package verify

//||------------------------------------------------------------------------------------------------||
//|| Approval Struct
//||------------------------------------------------------------------------------------------------||

type Moderate struct {
	Status    string `json:"status"`
	Moderator string `json:"moderator,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Details   string `json:"details,omitempty"`
}
