package constants

import "reflect"

type VerificationStatus struct {
	Pending             string
	PendingVerification string
	Verified            string
	Rejected            string
	Escalated           string
	Expired             string
	Cancelled           string
	Missing             string
}

var VerificationStatuses = VerificationStatus{
	Pending:             "PEND",
	PendingVerification: "PEVF",
	Verified:            "VERF",
	Rejected:            "RJCT",
	Escalated:           "ESCL",
	Expired:             "EXPD",
	Cancelled:           "CNCL",
	Missing:             "MISS",
}

func GetAllVerificationStatuses() []string {
	v := reflect.ValueOf(VerificationStatuses)
	var statuses []string

	for i := 0; i < v.NumField(); i++ {
		statuses = append(statuses, v.Field(i).String())
	}

	return statuses
}
