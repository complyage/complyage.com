package identity

import "agent/verify"

//||------------------------------------------------------------------------------------------------||
//|| Identity
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	//||------------------------------------------------------------------------------------------------||
	//|| Meta
	//||------------------------------------------------------------------------------------------------||
	ID           int64           `json:"id,omitempty"`
	VerifiedDOB  verify.DOB      `json:"verifiedDOB,omitempty"`
	VerifiedType verify.DataType `json:"verifiedMethod,omitempty"`
	Approved     []string        `json:"approved,omitempty"`
	//||------------------------------------------------------------------------------------------------||
	//|| Areas
	//||------------------------------------------------------------------------------------------------||
	Address    IdentityRecord             `json:"address,omitempty"`
	CreditCard IdentityRecord             `json:"creditCard,omitempty"`
	Email      IdentityRecord             `json:"email,omitempty"`
	Face       IdentityRecord             `json:"face,omitempty"`
	IDCard     IdentityRecord             `json:"idCard,omitempty"`
	Phone      IdentityRecord             `json:"phone,omitempty"`
	Usernames  map[int64]IdentityUsername `json:"usernames,omitempty"`
}
