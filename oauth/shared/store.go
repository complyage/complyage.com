package shared

import (
	"encoding/json"
	"time"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Create The Shared Wrapper
//||------------------------------------------------------------------------------------------------||

func (oas OAuthSharedAccess) Store() error {
	jsonData, err := json.Marshal(oas)
	if err != nil {
		return err
	}
	err = app.CacheRedis["oauth"].Set("shared:"+oas.Token, string(jsonData), time.Duration(15)*time.Minute)
	if err != nil {
		return err
	}
	return nil
}
