package constants

import "reflect"

type VerificationType struct {
	Email        string
	Phone        string
	Age          string
	Address      string
	CreditCard   string
	ProfilePhoto string
	Username     string
}

var VerificationTypes = VerificationType{
	Email:        "MAIL",
	Phone:        "PHNE",
	Age:          "UAGE",
	Address:      "ADDR",
	CreditCard:   "CRCD",
	ProfilePhoto: "PROF",
	Username:     "UNAM",
}

func GetAllVerificationTypes() []string {
	v := reflect.ValueOf(VerificationTypes)
	var types []string

	for i := 0; i < v.NumField(); i++ {
		types = append(types, v.Field(i).String())
	}

	return types
}
