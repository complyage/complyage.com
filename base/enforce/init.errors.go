package enforce

import "github.com/ralphferrara/aria/app"

func init() {
	app.Err("Enforce").Add("MISSING_CLIENT_ID", "The API key is required", false)
	app.Err("Enforce").Add("INVALID_SITE_KEY", "The provided site key is invalid", false)
	app.Err("Enforce").Add("INVALID_SITE_STATUS", "The site provided is not Active", false)
	app.Err("Enforce").Add("INVALID_DOMAIN", "This domain is not authorized for this api key", false)
}
