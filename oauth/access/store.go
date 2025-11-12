package access

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func SessionKey(token string) string {
	return "access:" + token
}

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func RetrieveAuthorizationToken(authorizationCode, clientId, clientSecret string) (string, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Split the Authorization Code into Token and Code
	//||------------------------------------------------------------------------------------------------||
	parts := strings.Split(authorizationCode, ",")
	if len(parts) != 2 {
		return "", app.Err("OAuth").Error("INVALID_AUTHORIZATION_CODE")
	}
	token := parts[0]
	authCode := parts[1]
	//||------------------------------------------------------------------------------------------------||
	//|| Retrieve from Cache
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("Loading Access for Token:", token)
	val, err := app.CacheRedis["oauth"].Get(SessionKey(token))
	if err != nil {
		return "", app.Err("OAuth").Error("MISSING_ACCESS")
	}
	var oa OAuthAccess
	err = json.Unmarshal([]byte(val), &oa)
	if err != nil {
		return "", app.Err("OAuth").Error("CORRUPT_ACCESS")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check the Access Code
	//||------------------------------------------------------------------------------------------------||
	if oa.AuthorizeToken != authCode {
		return "", app.Err("OAuth").Error("INVALID_AUTHORIZATION_CODE")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if Client ID matches
	//||------------------------------------------------------------------------------------------------||
	if oa.Enforcement.Site.ClientId != clientId {
		return "", app.Err("OAuth").Error("INVALID_CLIENT_ID")
	}
	if oa.Enforcement.Site.Private != clientSecret {
		return "", app.Err("OAuth").Error("INVALID_CLIENT_SECRET")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if Expired
	//||------------------------------------------------------------------------------------------------||
	if time.Now().After(oa.ExpiresAt) {
		return "", app.Err("OAuth").Error("AUTHORIZATION_CODE_EXPIRED")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if Approved
	//||------------------------------------------------------------------------------------------------||
	if !oa.Approved {
		return "", app.Err("OAuth").Error("AUTHORIZATION_NOT_APPROVED")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if Used Already
	//||------------------------------------------------------------------------------------------------||
	if oa.AuthorizeKeyUsed {
		return "", app.Err("OAuth").Error("AUTHORIZATION_CODE_ALREADY_USED")
	}
	oa.AuthorizeKeyUsed = true
	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||
	oa.Store()
	//||------------------------------------------------------------------------------------------------||
	//|| Check if
	//||------------------------------------------------------------------------------------------------||
	return oa.AuthorizeToken, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func LoadAccessCode(token string) (OAuthAccess, error) {
	fmt.Println("Loading Access for Token:", token)
	val, err := app.CacheRedis["oauth"].Get(SessionKey(token))
	if err != nil {
		return OAuthAccess{}, app.Err("OAuth").Error("MISSING_ACCESS")
	}
	var oa OAuthAccess
	err = json.Unmarshal([]byte(val), &oa)
	if err != nil {
		return OAuthAccess{}, app.Err("OAuth").Error("CORRUPT_ACCESS")
	}
	return oa, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Remove Access
//||------------------------------------------------------------------------------------------------||

func (oa OAuthAccess) RemoveAccess() error {
	fmt.Println("Removing Access for Token:", oa.Token)
	err := app.CacheRedis["oauth"].Del(SessionKey(oa.Token))
	if err != nil {
		return err
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Store
//||------------------------------------------------------------------------------------------------||

func (oa OAuthAccess) Store() error {
	fmt.Println("Storing Access for Token:", oa.Token)
	err := app.CacheRedis["oauth"].Set(SessionKey(oa.Token), oa.String(), 30*time.Minute)
	if err != nil {
		return err
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (oa OAuthAccess) String() string {
	jsonStr, err := json.Marshal(oa)
	if err != nil {
		fail, fErr := json.Marshal(OAuthAccess{
			Token:        "BADTOKEN",
			Approved:     false,
			AuthorizeKey: random.RandomString(48),
			DenyKey:      random.RandomString(48),
		})
		if fErr != nil {
			return "{}"
		}
		return string(fail)
	}
	return string(jsonStr)
}
