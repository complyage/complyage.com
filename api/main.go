package main

import (
	"api/handlers"
	"fmt"

	"github.com/complyage/base/adapters"
	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/encrypted"
	"github.com/complyage/base/ips"
	"github.com/complyage/base/verify"
	"github.com/complyage/base/zones"

	"github.com/joho/godotenv"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth"
	"github.com/ralphferrara/aria/auth/setup"
	"github.com/ralphferrara/aria/http"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Main
//||------------------------------------------------------------------------------------------------||

func main() {
	//||------------------------------------------------------------------------------------------------||
	//|| Load Env
	//||------------------------------------------------------------------------------------------------||
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("No .env file found, continuing...")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Starting switch-over to aria
	//||------------------------------------------------------------------------------------------------||
	app.Init("../config.json")
	app.ListenForShutdown()
	//||------------------------------------------------------------------------------------------------||
	//|| Initialize Auth
	//||------------------------------------------------------------------------------------------------||
	setup.Setup.Functions.OnAccountComplete = abstract.OnAccountComplete
	setup.Setup.Functions.OnAuthCheck = abstract.AuthMe
	setup.Setup.Domain = ".complyage.com"
	auth.Init(app.SQLDB["main"].DB, app.Config.Auth, ".complyage.com")
	//||------------------------------------------------------------------------------------------------||
	//|| Cors
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("Environment:", app.Config.App.Env)
	log.PrettyPrint(app.Config.App)
	if app.Config.App.Env == "development" {
		app.HTTP["api"].SetDefaults(http.Defaults{
			Cors: http.CORS([]string{
				"http://localhost:5173",
				"http://localhost:5174",
				"http://127.0.0.1:5173",
				"http://127.0.0.1:5174",
				"http://localhost:8080",
				"http://localhost:8081",
				"*", // allow wildcard only in dev for flexibility
			}),
			Middleware: http.Logger(),
		})
	} else {
		app.HTTP["api"].SetDefaults(http.Defaults{
			Cors: http.CORS([]string{
				"https://complyage.com",
				"https://www.complyage.com",
				"https://api.complyage.com",
				"https://gate.complyage.com",
				"https://oauth.complyage.com",
			}),
			Middleware: http.Logger(),
		})
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Start HTTP
	//||------------------------------------------------------------------------------------------------||
	app.HTTP["api"].Start()
	//||------------------------------------------------------------------------------------------------||
	//|| Register Handlers
	//||------------------------------------------------------------------------------------------------||
	handlers.InitRoutes()
	//||------------------------------------------------------------------------------------------------||
	//|| Initialize Stripe
	//||------------------------------------------------------------------------------------------------||
	adapters.InitStripe()
	//||------------------------------------------------------------------------------------------------||
	//|| Load IP Ranges
	//||------------------------------------------------------------------------------------------------||
	if app.Config.App.Env != "development" {
		ips.LoadIPRanges()
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Load Zones
	//||------------------------------------------------------------------------------------------------||
	zones.LoadZones()
	//||------------------------------------------------------------------------------------------------||
	//|| Verify Init
	//||------------------------------------------------------------------------------------------------||
	verify.Init(app.Storages["verifications"], app.SQLDB["main"].DB)
	encrypted.Init(app.Storages["encrypted"])
	//||------------------------------------------------------------------------------------------------||
	//|| Keep Alive
	//||------------------------------------------------------------------------------------------------||
	select {}
}
