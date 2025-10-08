package verify

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Approval Struct
//||------------------------------------------------------------------------------------------------||

type Moderate struct {
	Status    string    `json:"status"`
	Moderator string    `json:"moderator,omitempty"`
	Timestamp time.Time `json:"timestamp,omitempty"`
	Details   string    `json:"details,omitempty"`
}
