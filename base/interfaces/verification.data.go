package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Verified Data
//||------------------------------------------------------------------------------------------------||

type VerificationData struct {
	UAGE DOB            `json:"UAGE,omitempty"`
	MAIL EmailAddress   `json:"MAIL,omitempty"`
	PHNE PhoneNumber    `json:"PHNE,omitempty"`
	ADDR Address        `json:"ADDR,omitempty"`
	CRCD CreditCard     `json:"CRCD,omitempty"`
	IDEN Identification `json:"IDEN,omitempty"`
}
