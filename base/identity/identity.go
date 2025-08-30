package identity

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	Email       IdentityRecord            `json:"email,omitempty"`
	Age         IdentityRecord            `json:"age,omitempty"`
	Phone       IdentityRecord            `json:"phone,omitempty"`
	Address     IdentityRecord            `json:"address,omitempty"`
	CreditCard  IdentityRecord            `json:"creditCard,omitempty"`
	IDCard      IdentityRecord            `json:"idCard,omitempty"`
	Usernames   map[int64]IdentityUsernme `json:"usernames,omitempty"`
	Approved    []string                  `json:"approved,omitempty"`
	Verified    bool                      `json:"verified"`
	VerifiedAge int                       `json:"verifiedAge,omitempty"`
}
