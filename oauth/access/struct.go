package access

import "github.com/complyage/base/enforce"

//||------------------------------------------------------------------------------------------------||
//|| Access Data stored in Redis
//||------------------------------------------------------------------------------------------------||

type OAuthAccess struct {
	Token        string              `json:"token"`
	Enforcement  enforce.Enforcement `json:"enforcement"`
	Approved     bool                `json:"approved"`
	PreviewKey   string              `json:"previewKey"`
	AuthorizeKey string              `json:"authorizeKey"`
	DenyKey      string              `json:"denyKey"`
	BypassKey    string              `json:"bypassKey"`
}
