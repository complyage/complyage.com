package identity

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type IdentityUsernme struct {
	IDSite       int64  `json:"idSite"`
	Username     string `json:"username"`
	Verification string `json:"verification"`
}
