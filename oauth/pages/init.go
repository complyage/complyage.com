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
	//|| Age Gate Handlers
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["oauth"].Router.HandleFunc("/v1/age", ServeAgeGateHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/verified", VerifiedHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/bypass", ApproveOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/preview", PreviewShareHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| OAuth Handlers
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["oauth"].Router.HandleFunc("/v1/authorize", ServeOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/deny", DenyOAuthHandler).Methods("GET")
	app.HTTP["oauth"].Router.HandleFunc("/v1/approve", ApproveOAuthHandler).Methods("GET")
}
