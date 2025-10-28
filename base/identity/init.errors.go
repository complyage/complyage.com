package identity

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Identity :: Errors
//||------------------------------------------------------------------------------------------------||

func init() {
	app.Err("Identity").Add("IDENTITY_LOAD_FAILED", "Failed to Load Identity Data", false)
	app.Err("Identity").Add("IDENTITY_UPDATE_FAILED", "Failed to Update Identity Data", false)
}
