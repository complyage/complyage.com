package identity

import "github.com/complyage/base/types"

//||------------------------------------------------------------------------------------------------||
//|| Identity
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	//||------------------------------------------------------------------------------------------------||
	//|| Meta
	//||------------------------------------------------------------------------------------------------||
	ID       int64    `json:"id,omitempty"`
	Approved []string `json:"approved,omitempty"`
	//||------------------------------------------------------------------------------------------------||
	//|| Areas
	//||------------------------------------------------------------------------------------------------||
	Address    IdentityRecord              `json:"ADDR"`
	CreditCard IdentityRecord              `json:"CRCD"`
	Email      IdentityRecord              `json:"MAIL"`
	Face       IdentityRecord              `json:"FACE"`
	IDCard     IdentityRecord              `json:"IDEN"`
	Phone      IdentityRecord              `json:"PHNE"`
	Usernames  map[string]IdentityUsername `json:"usernames"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity Username
//||------------------------------------------------------------------------------------------------||

type IdentityRecord struct {
	Verified     bool      `json:"verified"`
	Age          int       `json:"age"`
	DOB          types.DOB `json:"dob"`
	Display      string    `json:"display"`
	Verification string    `json:"verification"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity Username
//||------------------------------------------------------------------------------------------------||

type IdentityUsername struct {
	IDSite       int64  `json:"idSite"`
	Username     string `json:"username"`
	Verification string `json:"verification"`
}
