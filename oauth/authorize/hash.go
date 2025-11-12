package authorize

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Generate the Hash
//||------------------------------------------------------------------------------------------------||

func GenerateHash(action string, session AuthorizeSession) string {
	mac := hmac.New(sha256.New, []byte(session.Salt))
	mac.Write([]byte(action + session.SessionId + session.ClientId + session.RedirectURI + session.Scope + session.State + session.ExternalUserId))
	sum := mac.Sum(nil)
	return base64.RawURLEncoding.EncodeToString(sum)
}

//||------------------------------------------------------------------------------------------------||
//|| Validate the Action Hash
//||------------------------------------------------------------------------------------------------||

func ActionCheck(action string, session AuthorizeSession, providedHash string) error {
	expectedHash := GenerateHash(action, session)
	if !hmac.Equal([]byte(expectedHash), []byte(providedHash)) {
		return app.Err("Authorize").Error("INVALID_ACTION_HASH")
	}
	return nil
}
