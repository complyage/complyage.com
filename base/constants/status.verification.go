package constants

import "reflect"

type VerificationStatus struct {
	Pending   string
	Verified  string
	Rejected  string
	Escalated string
	Expired   string
	Cancelled string
}

var VerificationStatuses = VerificationStatus{
	Pending:   "PEND",
	Verified:  "VERF",
	Rejected:  "RJCT",
	Escalated: "ESCL",
	Expired:   "EXPD",
	Cancelled: "CNCL",
}

func GetAllVerificationStatuses() []string {
	v := reflect.ValueOf(VerificationStatuses)
	var statuses []string

	for i := 0; i < v.NumField(); i++ {
		statuses = append(statuses, v.Field(i).String())
	}

	return statuses
}
