package constants

type AccountStatusType struct {
	Created  string
	Pending  string
	Verified string
	Active   string
	Banned   string
	Closed   string
}

var AccountStatus = AccountStatusType{
	Created:  "PNEW",
	Pending:  "PEND",
	Verified: "VERF",
	Active:   "ACTV",
	Banned:   "BNND",
	Closed:   "RMVD",
}
