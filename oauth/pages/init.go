package pages

import (
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
	app.HTTP["oauth"].Router.HandleFunc("/v1/age", ServeAgeGateHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/login", ServeLogin).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/authorize", ServeOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/deny", DenyOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/approve", ApproveOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/bypass", BypassOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/preview", PreviewShareHandler).Methods("GET")
}
