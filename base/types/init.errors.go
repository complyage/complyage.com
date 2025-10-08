package types

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Types :: Errors
//||------------------------------------------------------------------------------------------------||

func init() {
	app.Err("Types").Add("INVALID_PHONE", "Invalid Phone Number", false)
}
