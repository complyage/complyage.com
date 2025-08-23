package main

import (
	"api/handlers"
	"base/adapters"
	"base/db"
	"base/helpers"
	"base/loaders"
	"base/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/ralphferrara/aria/storage"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Use DB vs In-Memory
//||------------------------------------------------------------------------------------------------||

var UseInMemory bool

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
	//|| Should we use in-memory storage or DB?
	//||------------------------------------------------------------------------------------------------||
	env := os.Getenv("ENV_MODE")
	fmt.Println("ENV_MODE = " + env)
	if env == "production" {
		fmt.Println("Running in production mode, using In memory storage")
		UseInMemory = true
	} else {
		fmt.Println("Running in development mode, using DB")
		UseInMemory = false
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Starting switch-over to aria
	//||------------------------------------------------------------------------------------------------||
	if err := app.Init("config.json"); err != nil {
		app.Log.Error("main", "Startup failed: %v", err)
		os.Exit(1)
	}
	app.Log.Info("main", "API started")
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to Storage
	//||------------------------------------------------------------------------------------------------||
	var Store storage.Storage
	Store.Init()
	//||------------------------------------------------------------------------------------------------||
	//|| Open DB Connection
	//||------------------------------------------------------------------------------------------------||
	db.ConnectMySQL()
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to Redis
	//||------------------------------------------------------------------------------------------------||
	db.ConnectRedis() // Redis
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to RabbitMQ
	//||------------------------------------------------------------------------------------------------||
	db.ConnectMQ() // RabbitMQ
	//||------------------------------------------------------------------------------------------------||
	//|| GORM
	//||------------------------------------------------------------------------------------------------||
	db.DB.AutoMigrate(&models.Account{})
	//||------------------------------------------------------------------------------------------------||
	//|| Initialize Stripe
	//||------------------------------------------------------------------------------------------------||
	adapters.InitStripe()
	//||------------------------------------------------------------------------------------------------||
	//|| Open DB Connection
	//||------------------------------------------------------------------------------------------------||
	if UseInMemory {
		fmt.Println("Loading IP ranges...Please be patient, takes minutes")
		if err := loaders.LoadIPRanges(); err != nil {
			panic(fmt.Sprintf("Failed to load IP ranges: %v", err))
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Load Zones
	//||------------------------------------------------------------------------------------------------||
	loaders.LoadZones()
	//||------------------------------------------------------------------------------------------------||
	//|| Load Sites
	//||------------------------------------------------------------------------------------------------||
	loaders.StartSiteLoader()
	//||------------------------------------------------------------------------------------------------||
	//|| Setup Router
	//||------------------------------------------------------------------------------------------------||
	router := mux.NewRouter()
	//||------------------------------------------------------------------------------------------------||
	//|| Middleware
	//||------------------------------------------------------------------------------------------------||
	router.Use(LoggerMiddleware)
	//||------------------------------------------------------------------------------------------------||
	//|| Get News/Zones Public
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/news", handlers.NewsHandler).Methods("GET")
	router.HandleFunc("/v1/api/zones", handlers.ZoneHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Member Sites
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/sites/zones", handlers.SitesZoneHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/vtypes", handlers.VerificationTypesListHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/list", handlers.SitesListHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/load", handlers.SitesLoadHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/upload", handlers.UploadHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/create", handlers.SitesNewHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/copy", handlers.SitesCopyHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/update", handlers.SitesUpdateHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/delete", handlers.SitesDeleteHandler).Methods("DELETE")
	//||------------------------------------------------------------------------------------------------||
	//|| Publc Routes
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/auth/signup", handlers.SignupHandler).Methods("POST")
	router.HandleFunc("/auth/forgot", handlers.ForgotPasswordHandler).Methods("POST")
	router.HandleFunc("/auth/twofactor", handlers.TwoFactorHandler).Methods("POST")
	router.HandleFunc("/auth/me", handlers.AuthMeHandler).Methods("GET")
	router.HandleFunc("/auth/complete", handlers.CompleteHandler).Methods("POST", "GET")
	router.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", handlers.LogoutHandler).Methods("GET")
	router.HandleFunc("/auth/generate", handlers.GenerateKeyPairHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| User Routes
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/user/dashboard", handlers.UserDashboard).Methods("GET")
	router.HandleFunc("/user/verifications", handlers.UserVerifications).Methods("GET")
	router.HandleFunc("/user/reset", handlers.ResetPasswordHandler).Methods("POST")
	router.HandleFunc("/user/quit", handlers.QuitHandler).Methods("POST")
	router.HandleFunc("/user/delete-account", handlers.DeleteAccountHandler).Methods("POST")
	router.HandleFunc("/user/shared", handlers.UserSharedHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/init", handlers.VerificationInit).Methods("GET")
	router.HandleFunc("/v1/api/verify/load", handlers.VerificationCodeLoad).Methods("GET")
	router.HandleFunc("/v1/api/verify/check", handlers.VerificationCheckCodeHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/list", handlers.GetVerificationsList).Methods("GET")
	router.HandleFunc("/v1/api/verify/qr/generate", handlers.QRCodeGenerate).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - ID
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/id/init", handlers.IDVerifyInitHandler).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/status", handlers.VerifyIDStatusHandler).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/media/upload", handlers.VerifyIDMediaUpload).Methods("POST")
	router.HandleFunc("/v1/api/verify/id/media/fetch", handlers.VerifyIDMediaFetch).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/success", handlers.VerifyIDSuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Card
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/phone", handlers.PhoneVerifyInitHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Card
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/card", handlers.CCVerifyInitHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/card/check", handlers.CCVerifyCheckHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/card/success", handlers.CCVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Address
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/address", handlers.AddressVerifyInitHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/address/success", handlers.AddressVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Tools
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/currency", handlers.ConvertUSDHandler).Methods("GET")
	router.HandleFunc("/v1/api/geo/address", handlers.GoogleAddressVerifyHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Status Update - ID Verification
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/internal/verify/id/progress", handlers.VerifyIDProgressHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Simple Up Check
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/agent/fatal", handlers.AgentFatalHandler).Methods("POST")
	router.HandleFunc("/health", handlers.HealthHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Cors Middleware - Need to update to handle CORS properly
	//||------------------------------------------------------------------------------------------------||
	allowedOrigins := []string{"*"} // or list your domains
	//||------------------------------------------------------------------------------------------------||
	//|| Logger Middleware
	//||------------------------------------------------------------------------------------------------||
	router.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s => 404\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
		http.Error(w, "404 page not found", http.StatusNotFound)
	})
	//||------------------------------------------------------------------------------------------------||
	//|| We are up and running
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("API server running on :" + os.Getenv("PORT_HTTP_API"))
	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("PORT_HTTP_API"),
			helpers.CORSMiddleware(allowedOrigins, router),
		),
	)

}

//||------------------------------------------------------------------------------------------------||
//|| Logger Middleware
//||------------------------------------------------------------------------------------------------||

func LoggerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Printf("[%s] %s %s\n", time.Now().Format(time.RFC3339), r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
