package handlers

import (
	"oauth/handlers/pages"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

//||------------------------------------------------------------------------------------------------||
//|| Auth
//||------------------------------------------------------------------------------------------------||

func InitRoutes() {
	//||------------------------------------------------------------------------------------------------||
	//|| Global Routes
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["oauth"].Router.HandleFunc("/v1/authorize", pages.ServeOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/deny", pages.DenyOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/approve", pages.ApproveOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/private", pages.ServePrivateKeyForm).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/return", pages.OAuthReturnHandler).Methods("GET")
}
