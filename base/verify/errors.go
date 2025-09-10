package verify

import (
	"github.com/ralphferrara/aria/app"
)

func initErrors() {
	//||------------------------------------------------------------------------------------------------||
	//|| Verify Steps
	//||------------------------------------------------------------------------------------------------||
	app.Err("verify").Add("VERIFY_PANIC", "An unknown error has occured", true)
	app.Err("verify").Add("NO_UUID", "No UUID provided", true)
	app.Err("verify").Add("RECORD_NOT_FOUND", "Verification Record not found", true)
	app.Err("verify").Add("RECORD_CORRUPT", "Verification Record was not able to be parsed", true)
	app.Err("verify").Add("SAVE_FAIL", "Failed to save verification", true)
}
