package main

import "github.com/ralphferrara/aria/app"

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| API Base
	//||------------------------------------------------------------------------------------------------||
	app.Err("API").Add("ROUTE_NOT_FOUND", "The requested route was not found", false)
	app.Err("API").Add("INVALID_JSON", "The provided JSON was not valid", false)
	app.Err("API").Add("MISSING_FIELDS", "One or more required fields were missing", false)
	app.Err("API").Add("INVALID_FIELDS", "One or more fields were invalid", false)
	app.Err("API").Add("UNAUTHORIZED", "You are not authorized to access this resource", false)
	app.Err("API").Add("FORBIDDEN", "You do not have permission to access this resource", false)
	app.Err("API").Add("NOT_FOUND", "The requested resource was not found", false)
	app.Err("API").Add("METHOD_NOT_ALLOWED", "The requested method is not allowed for this resource", false)
	app.Err("API").Add("INTERNAL_SERVER_ERROR", "An internal server error occurred", true)
	app.Err("API").Add("BAD_GATEWAY", "A bad gateway error occurred", true)
	app.Err("API").Add("MISSING_COOKIE", "No Private Cookie foud", true)
	app.Err("API").Add("MISSING_PRIVATE", "Private Session", true)
	app.Err("API").Add("BAD_CACHE", "Could not retrieve from cache", true)
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypted
	//||------------------------------------------------------------------------------------------------||
	app.Err("API").Add("ENCRYPTED_FAILED", "Could not create encrypted data", true)
	app.Err("API").Add("ENCRYPTED_UPLOAD_FAILED", "Could not stored encrypted data", true)
}
