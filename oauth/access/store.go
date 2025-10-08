package access

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/random"
)

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func LoadAccess(token string) (OAuthAccess, error) {
	fmt.Println("Loading Access for Token:", token)
	val, err := app.CacheRedis["oauth"].Get("access:" + token)
	if err != nil {
		fmt.Println(err)
		return OAuthAccess{}, err
	}
	var oa OAuthAccess
	err = json.Unmarshal([]byte(val), &oa)
	if err != nil {
		return OAuthAccess{}, err
	}
	return oa, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Remove Access
//||------------------------------------------------------------------------------------------------||

func (oa OAuthAccess) RemoveAccess() error {
	fmt.Println("Removing Access for Token:", oa.Token)
	err := app.CacheRedis["oauth"].Del("access:" + oa.Token)
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
	err := app.CacheRedis["oauth"].Set("access:"+oa.Token, oa.String(), 30*time.Minute)
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
