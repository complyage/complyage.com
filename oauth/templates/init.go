package templates

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

func InitTemplates() {
	//||------------------------------------------------------------------------------------------------||
	//|| Register Templates
	//||------------------------------------------------------------------------------------------------||
	template.Register("oauth", "html/oauth.html")
	template.Register("private", "html/private.html")
	template.Register("private_locked", "html/private.locked.html")
	template.Register("private_unlocked", "html/private.unlocked.html")
	template.Register("private_key", "html/private.key.html")
	template.Register("private_bip39", "html/private.bip39.html")
	template.Register("permission", "html/permission.html")
}
