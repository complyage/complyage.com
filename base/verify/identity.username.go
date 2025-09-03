package verify

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type IdentityUsername struct {
	IDSite       int64  `json:"idSite"`
	Username     string `json:"username"`
	Verification string `json:"verification"`
}
