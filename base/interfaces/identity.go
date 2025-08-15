package interfaces

import (
	"base/constants"
)

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	Email      string                       `json:"email,omitempty"`
	Age        string                       `json:"age,omitempty"`
	Phone      string                       `json:"phone,omitempty"`
	Address    string                       `json:"address,omitempty"`
	CreditCard string                       `json:"creditCard,omitempty"`
	Usernames  map[int64]string             `json:"usernames,omitempty"`
	Approved   []constants.VerificationType `json:"approved,omitempty"`
}
