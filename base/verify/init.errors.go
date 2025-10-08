package verify

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Errors
//||------------------------------------------------------------------------------------------------||

func init() {
	app.Err("Verify").Add("DATABASE_UPDATE_FAILED", "Could not update Verification database record", false)
	app.Err("Verify").Add("DATABASE_INSERT_FAILED", "Could not insert Verification database record", false)
	app.Err("Verify").Add("NIL_VERIFICATION", "Verification is not initialized", false)
	app.Err("Verify").Add("MISSING_IDENTIFIER", "Identifier missing", false)
	app.Err("Verify").Add("VERIFY_LOAD_ACCOUNT_MISMATCH", "Verification Mismatch", false)
	app.Err("Verify").Add("VERIFY_SAVE_MARSHAL", "Could not convert verification for saving", false)
	app.Err("Verify").Add("VERIFY_SAVE_PUT", "Could not save Verification to Storage", false)
	app.Err("Verify").Add("VERIFY_CREATE_FAILED", "Could not create a new verification failed", false)
	app.Err("Verify").Add("VERIFY_LOAD_UUID", "Invalid Verification UUID", false)
	app.Err("Verify").Add("VERIFY_LOAD_FAILED", "Failed to fetch Verification", false)
	app.Err("Verify").Add("VERIFY_LOAD_UNMARSHAL", "Failed to open Verification", false)
	app.Err("Verify").Add("TWOFACTOR_EXPIRED", "Verification Code Expired", false)
	app.Err("Verify").Add("TWOFACTOR_TOOMANY", "Too many attempts", false)
	app.Err("Verify").Add("TWOFACTOR_NOTSET", "Verification Code Not Set", false)
	app.Err("Verify").Add("TWOFACTOR_INVALID", "Verification Code Incorrect", false)
	app.Err("Verify").Add("UNKNOWN_VERIFICATION_TYPE", "Unknown Verification Type", false)
}
