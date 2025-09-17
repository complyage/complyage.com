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
	template.Register("oauth", "templates/html/oauth.html")
	template.Register("private", "templates/html/private.html")
	template.Register("private_locked", "templates/html/private.locked.html")
	template.Register("private_unlocked", "templates/html/private.unlocked.html")
	template.Register("private_key", "templates/html/private.key.html")
	template.Register("private_bip39", "templates/html/private.bip39.html")
	template.Register("permission", "templates/html/permission.html")
}
