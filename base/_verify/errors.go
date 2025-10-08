package verify

import (
	"github.com/ralphferrara/aria/app"
)

func initErrors() {
	//||------------------------------------------------------------------------------------------------||
	//|| Verify Steps
	//||------------------------------------------------------------------------------------------------||
	app.Err("Verify").Add("VERIFY_PANIC", "An unknown error has occured", true)
	app.Err("Verify").Add("NO_UUID", "No UUID provided", true)
	app.Err("Verify").Add("RECORD_NOT_FOUND", "Verification Record not found", true)
	app.Err("Verify").Add("RECORD_CORRUPT", "Verification Record was not able to be parsed", true)
	app.Err("Verify").Add("SAVE_FAIL", "Failed to save verification", true)
	app.Err("Verify").Add("CREATE_FAILED", "Failed to create verification record", true)
}
