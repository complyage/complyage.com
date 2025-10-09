package templates

import "github.com/ralphferrara/aria/base/template"

//||------------------------------------------------------------------------------------------------||
//|| Init
//||------------------------------------------------------------------------------------------------||

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| Main
	//||------------------------------------------------------------------------------------------------||
	template.Register("oauth", "./assets/html/oauth.html")
	template.Register("error", "./assets/html/error.html")
	template.Register("verified", "./assets/html/verified.html")
	//||------------------------------------------------------------------------------------------------||
	//|| Main
	//||------------------------------------------------------------------------------------------------||
	template.Register("sub.age", "./assets/html/sub.age.html")
	//||------------------------------------------------------------------------------------------------||
	//|| Permissions
	//||------------------------------------------------------------------------------------------------||
	template.Register("sub.permission", "./assets/html/sub.permission.html")
	template.Register("sub.permission.age", "./assets/html/sub.permission.age.html")
	//||------------------------------------------------------------------------------------------------||
	//|| Badge/Button
	//||------------------------------------------------------------------------------------------------||
	template.Register("verify.badge", "./assets/html/verify.badge.html")
	template.Register("verify.button", "./assets/html/verify.button.html")
	template.Register("verify.login", "./assets/html/verify.login.html")
	template.Register("verify.use", "./assets/html/verify.use.html")
	//||------------------------------------------------------------------------------------------------||
	//|| Login Status
	//||------------------------------------------------------------------------------------------------||
	template.Register("status.loggedin", "./assets/html/status.loggedin.html")
	template.Register("status.loggedout", "./assets/html/status.loggedout.html")
	//||------------------------------------------------------------------------------------------------||
	//|| Footer
	//||------------------------------------------------------------------------------------------------||
	template.Register("footer.notrequired", "./assets/html/footer.notrequired.html")
	template.Register("footer.loggedout", "./assets/html/footer.loggedout.html")
	template.Register("footer.requires", "./assets/html/footer.requires.html")
	template.Register("footer.private", "./assets/html/footer.private.html")
	template.Register("footer.verified", "./assets/html/footer.verified.html")
	template.Register("footer.verified.age", "./assets/html/footer.verified.age.html")
}
