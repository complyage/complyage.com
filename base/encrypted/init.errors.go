package encrypted

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Encrypted :: Errors
//||------------------------------------------------------------------------------------------------||

func init() {
	app.Err("Encrypted").Add("ENCRYPT_CREATE_FAILED", "Failed to Encrypt Data", false)
	app.Err("Encrypted").Add("ENCRYPT_STORE_FAILED", "Failed to Store Encrypt Data", false)
	app.Err("Encrypted").Add("ENCRYPT_LOAD_FAILED", "Failed to Load Encrypt Data", false)
	app.Err("Encrypted").Add("ENCRYPT_LOAD_FAILED", "Failed to Load Encrypt Data", false)
	app.Err("Encrypted").Add("LOAD_PREVIEW_INVALID_INPUT", "Failed to Load Encrypted Data for preview", false)
}
