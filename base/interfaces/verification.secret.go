package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Meta Data
//||------------------------------------------------------------------------------------------------||

type VerificationSecret struct {
	Attempts   int    `json:"attempts"`
	Expiration string `json:"expiration"`
	CheckCode  string `json:"checkCode"`
}
