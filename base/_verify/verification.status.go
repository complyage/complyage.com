package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| StatusType
//||------------------------------------------------------------------------------------------------||

type StatusType int

const (
	StatusPending StatusType = iota
	StatusPendingVerification
	StatusInProgress
	StatusEscalated
	StatusVerified
	StatusRejected
	StatusExpired
)

//||------------------------------------------------------------------------------------------------||
//|| StatusType (helpers)
//||------------------------------------------------------------------------------------------------||

func (s StatusType) Code() string {
	switch s {
	case StatusPending:
		return "PEND"
	case StatusPendingVerification:
		return "PEVF"
	case StatusInProgress:
		return "INPR"
	case StatusEscalated:
		return "ESCL"
	case StatusVerified:
		return "VERF"
	case StatusRejected:
		return "RJCT"
	case StatusExpired:
		return "EXPD"
	default:
		return "UNKN"
	}
}

func (s StatusType) Description() string {
	switch s {
	case StatusPending:
		return "Queued and awaiting processing"
	case StatusPendingVerification:
		return "Pending Verification"
	case StatusInProgress:
		return "Currently being processed"
	case StatusEscalated:
		return "Operation was escalated for review"
	case StatusVerified:
		return "Successfully verified"
	case StatusRejected:
		return "Verification was rejected"
	case StatusExpired:
		return "Verification has expired"
	default:
		return "Unknown status"
	}
}

func (s StatusType) String() string {
	return s.Code()
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (s StatusType) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Code())
}

func (s *StatusType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "PEND":
		*s = StatusPending
	case "PEVF":
		*s = StatusPendingVerification
	case "INPR":
		*s = StatusInProgress
	case "ESCL":
		*s = StatusEscalated
	case "VERF":
		*s = StatusVerified
	case "RJCT":
		*s = StatusRejected
	case "EXPD":
		*s = StatusExpired
	default:
		return fmt.Errorf("invalid StatusType: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Dot notation namespace for StatusType
//||------------------------------------------------------------------------------------------------||

type nsStatusType struct {
	Pending             StatusType
	PendingVerification StatusType
	InProgress          StatusType
	Escalated           StatusType
	Verified            StatusType
	Rejected            StatusType
	Expired             StatusType
}

var STATUSES = nsStatusType{
	Pending:             StatusPending,
	PendingVerification: StatusPendingVerification,
	InProgress:          StatusInProgress,
	Escalated:           StatusEscalated,
	Verified:            StatusVerified,
	Rejected:            StatusRejected,
	Expired:             StatusExpired,
}
