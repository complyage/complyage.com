package authorize

import (
	"github.com/ralphferrara/aria/app"
)

func init() {
	app.Err("Authorize").Add("ROUTE_NOT_FOUND", "The requested route was not found", false)
}
