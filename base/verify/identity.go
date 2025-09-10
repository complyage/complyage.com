package verify

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
	Address    IdentityRecord             `json:"ADDR,omitempty"`
	CreditCard IdentityRecord             `json:"CRCD,omitempty"`
	Email      IdentityRecord             `json:"MAIL,omitempty"`
	Face       IdentityRecord             `json:"FACE,omitempty"`
	IDCard     IdentityRecord             `json:"IDEN,omitempty"`
	Phone      IdentityRecord             `json:"PHNE,omitempty"`
	Usernames  map[int64]IdentityUsername `json:"usernames,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity Username
//||------------------------------------------------------------------------------------------------||

type IdentityRecord struct {
	Verified     bool   `json:"verified"`
	Age          int    `json:"age,omitempty"`
	DOB          DOB    `json:"dob,omitempty"`
	Display      string `json:"display,omitempty"`
	Verification string `json:"verification,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity Username
//||------------------------------------------------------------------------------------------------||

type IdentityUsername struct {
	IDSite       int64  `json:"idSite"`
	Username     string `json:"username"`
	Verification string `json:"verification"`
}
