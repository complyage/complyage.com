package identity

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type IdentityRecord struct {
	Display      string `json:"display"`
	Verification string `json:"verification"`
}
