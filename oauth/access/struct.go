package access

import (
	"go/types"
	"time"

	"github.com/complyage/base/enforce"
)

//||------------------------------------------------------------------------------------------------||
//|| Access Data stored in Redis
//||------------------------------------------------------------------------------------------------||

type OAuthAccess struct {
	Token              string              `json:"token"`
	ExpiresAt          time.Time           `json:"expiresAt"`
	Enforcement        enforce.Enforcement `json:"enforcement"`
	Approved           bool                `json:"approved"`
	PreviewKey         string              `json:"previewKey"`
	DenyKey            string              `json:"denyKey"`
	BypassKey          string              `json:"bypassKey"`
	AuthorizeKey       string              `json:"authorizeKey"`
	AuthorizeToken     string              `json:"authorizeToken"`
	AuthorizeKeyUsed   bool                `json:"authorizeKeyUsed"`
	AuthorizeTokenUsed bool                `json:"authorizeTokenUsed"`
	UserInfo           OAuthUserInfo       `json:"userInfo"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserInfo
//||------------------------------------------------------------------------------------------------||

type OAuthUserInfo struct {
	UserId  int64         `json:"userId"`
	Address types.Address `json:"ADDR"`
}
