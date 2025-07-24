package interfaces

//||------------------------------------------------------------------------------------------------||
//|| OAuth Session
//||------------------------------------------------------------------------------------------------||

type OAuthSession struct {
	APIKey    string   `json:"apiKey"`
	AccessKey string   `json:"accessKey"`
	State     string   `json:"state"`
	Redirect  string   `json:"redirect"`
	Scope     []string `json:"scope"`
	Expires   int64    `json:"expires"`
	Created   int64    `json:"created"`
	Status    string   `json:"status"`
}

//||------------------------------------------------------------------------------------------------||
//|| OAuth Response Types
//||------------------------------------------------------------------------------------------------||

type OAuthVerification struct {
	Type   string      `json:"type"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type OAuthRequirements struct {
	Type     string `json:"type"`
	Optional bool   `json:"optional"`
}

type OAuthUser struct {
	ID            string              `json:"id"`
	Status        string              `json:"status"`
	Username      string              `json:"username"`
	Verifications []OAuthVerification `json:"verifications"`
}

type OAuthSite struct {
	Name        string `json:"name"`
	URL         string `json:"url"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
}

type OAuthZone struct {
	State         string              `json:"state"`
	Country       string              `json:"country"`
	IP            string              `json:"ip"`
	Requirements  []OAuthRequirements `json:"requirements"`
	Description   string              `json:"description"`
	Law           string              `json:"law"`
	EffectiveDate string              `json:"effectiveDate"`
}

type OAuthResponse struct {
	Site         OAuthSite           `json:"site"`
	User         OAuthUser           `json:"user"`
	Zone         OAuthZone           `json:"zone"`
	Status       string              `json:"status"`
	Requirements []OAuthRequirements `json:"requirements"`
}
