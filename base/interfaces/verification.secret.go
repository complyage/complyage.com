package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Meta Data
//||------------------------------------------------------------------------------------------------||

type VerificationSecret struct {
	Attempts   int    `json:"attempts,omitempty"`
	Expiration string `json:"expiration,omitempty"`
	CheckCode  string `json:"checkCode,omitempty"`
}
