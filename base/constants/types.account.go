package constants

type AccountType struct {
	User   string
	Vendor string
}

var AccountTypes = AccountType{
	User:   "USER",
	Vendor: "VNDR",
}
