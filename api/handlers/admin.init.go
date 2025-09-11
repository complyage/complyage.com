package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/handlers/admin"
	"net/http"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Init Admin Routes
//||------------------------------------------------------------------------------------------------||

func InitAdminRoutes() {
	//||------------------------------------------------------------------------------------------------||
	//|| Admin
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Router.HandleFunc("/admin/widget/verifications", AdminOnly(admin.VerificationsWidgetHandler, 100)).Methods("GET")
}

//||------------------------------------------------------------------------------------------------||
//|| AdminOnly Middleware Wrapper
//||------------------------------------------------------------------------------------------------||

func AdminOnly(next http.HandlerFunc, minLevel int) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, account, _, err := actions.LoadSessionAccount(r)
		if err != nil || account.ID == 0 {
			responses.Error(w, http.StatusUnauthorized, "ERR_LOGIN")
			return
		}
		if account.Level < minLevel {
			responses.Error(w, http.StatusUnauthorized, "UNAUTHORIZED")
			return
		}
		next(w, r)
	}
}
