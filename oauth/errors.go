package main

import "github.com/ralphferrara/aria/app"

func init() {
	app.Err("OAuth").Add("MISSING_CLIENT_ID", "Client ID is required", false)
	app.Err("OAuth").Add("INVALID_API_KEY", "API Key is invalud", false)
	app.Err("OAuth").Add("INVALID_SITE_STATUS", "Invalid Site Status", false)
	app.Err("OAuth").Add("INVALID_SITE_KEY", "Site key was not found", false)
	app.Err("OAuth").Add("INVALID_SITE", "Site is not active or pending approval", false)
	app.Err("OAuth").Add("INVALID_SCOPE", "Invalid Scope Requested", false)
	app.Err("OAuth").Add("DISABLED_SCOPE", "Disabled Scope Requested", false)
	app.Err("OAuth").Add("AUTO_NO_IP", "IP was not registered", false)
	app.Err("OAuth").Add("INVALID_LOCATION", "No Location found", false)
	app.Err("OAuth").Add("INVALID_DOMAIN", "Domain is not allowed", false)
	app.Err("OAuth").Add("MISSING_PRIVATE_COOKIE", "Missing private cookie", false)
}
