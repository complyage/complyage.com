package templates

import "github.com/ralphferrara/aria/base/template"

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| Register Templates
	//||------------------------------------------------------------------------------------------------||
	template.Register("oauth", "./assets/html/oauth.authorize.html")
	template.Register("error", "./assets/html/error.html")
	template.Register("sub.permission", "./assets/html/sub.permission.html")
}
