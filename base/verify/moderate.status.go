package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Moderate (iota-based)
//||------------------------------------------------------------------------------------------------||

type ModerateStatus int

const (
	ModeratePending ModerateStatus = iota
	ModerateProcessing
	ModerateApproved
	ModerateDenied
)

//||------------------------------------------------------------------------------------------------||
//|| String Functions
//||------------------------------------------------------------------------------------------------||

func (s ModerateStatus) String() string {
	switch s {
	case ModeratePending:
		return "PENDING"
	case ModerateProcessing:
		return "PROCESSING"
	case ModerateApproved:
		return "APPROVED"
	case ModerateDenied:
		return "DENIED"
	default:
		return "UNKNOWN"
	}
}

//||------------------------------------------------------------------------------------------------||
//|| String Functions
//||------------------------------------------------------------------------------------------------||

func (s ModerateStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

//||------------------------------------------------------------------------------------------------||
//|| String Functions
//||------------------------------------------------------------------------------------------------||

func (s *ModerateStatus) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "PENDING":
		*s = ModeratePending
	case "PROCESSING":
		*s = ModerateProcessing
	case "APPROVED":
		*s = ModerateApproved
	case "DENIED":
		*s = ModerateDenied
	default:
		return fmt.Errorf("invalid ApprovalStatus: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Make a namespace
//||------------------------------------------------------------------------------------------------||

type nsModerateStatus struct {
	Pending    ModerateStatus
	Processing ModerateStatus
	Approved   ModerateStatus
	Denied     ModerateStatus
}

//||------------------------------------------------------------------------------------------------||
//|| Make a namespace
//||------------------------------------------------------------------------------------------------||

var MODERATE_STATUS = nsModerateStatus{
	Pending:    ModeratePending,
	Processing: ModerateProcessing,
	Approved:   ModerateApproved,
	Denied:     ModerateDenied,
}
