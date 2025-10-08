package main

import "github.com/ralphferrara/aria/app"

//||------------------------------------------------------------------------------------------------||
//|| Errors :: Gate
//||------------------------------------------------------------------------------------------------||

func init() {
	app.Err("Gate").Add("API_KEY_MISSING", "Site Key not provided", false)
	app.Err("Gate").Add("API_KEY_INVALID", "API Key invalid", false)
	app.Err("Gate").Add("RECORD_MISSING", "Keeper Record Missing", false)
	app.Err("Gate").Add("RECORD_MISMATTCH", "Keeper Record Mismatch", false)

}
