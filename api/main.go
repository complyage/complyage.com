package main

import (
	"api/handlers/auth"
	"api/handlers/public"
	"api/handlers/user"
	"api/handlers/utils"
	"api/handlers/v1/agent"
	"api/handlers/v1/sites"
	"api/handlers/v1/verifier"
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
	//|| Auth
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/auth/signup", auth.SignupHandler).Methods("POST")
	router.HandleFunc("/auth/forgot", auth.ForgotPasswordHandler).Methods("POST")
	router.HandleFunc("/auth/twofactor", auth.TwoFactorHandler).Methods("POST")
	router.HandleFunc("/auth/me", auth.AuthMeHandler).Methods("GET")
	router.HandleFunc("/auth/complete", auth.CompleteHandler).Methods("POST", "GET")
	router.HandleFunc("/auth/login", auth.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", auth.LogoutHandler).Methods("GET")
	router.HandleFunc("/auth/generate", auth.GenerateKeyPairHandler).Methods("GET")
	router.HandleFunc("/auth/reset", auth.ResetPasswordHandler).Methods("POST")
	router.HandleFunc("/auth/quit", auth.QuitHandler).Methods("POST")
	router.HandleFunc("/auth/delete-account", auth.DeleteAccountHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| User Routes
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/user/dashboard", user.UserDashboard).Methods("GET")
	router.HandleFunc("/user/verifications", user.UserVerifications).Methods("GET")
	router.HandleFunc("/user/shared", user.UserSharedHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Get News/Zones Public
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/news", public.NewsHandler).Methods("GET")
	router.HandleFunc("/v1/api/zones", public.ZoneHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Member Sites
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/sites/zones", sites.SitesZoneHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/vtypes", sites.VerificationTypesListHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/list", sites.SitesListHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/load", sites.SitesLoadHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/upload", sites.UploadHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/create", sites.SitesNewHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/copy", sites.SitesCopyHandler).Methods("GET")
	router.HandleFunc("/v1/api/sites/update", sites.SitesUpdateHandler).Methods("POST")
	router.HandleFunc("/v1/api/sites/delete", sites.SitesDeleteHandler).Methods("DELETE")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/init", verifier.VerificationInit).Methods("GET")
	router.HandleFunc("/v1/api/verify/list", verifier.GetVerificationsList).Methods("GET")
	router.HandleFunc("/v1/api/verify/load", verifier.VerificationCodeLoad).Methods("GET")
	router.HandleFunc("/v1/api/verify/check", verifier.VerificationCodeCheck).Methods("POST")
	router.HandleFunc("/v1/api/verify/qr/generate", verifier.QRCodeGenerate).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - ID
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/id/init", verifier.IDVerifyInitHandler).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/status", verifier.VerifyIDStatusHandler).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/media/upload", verifier.VerifyIDMediaUpload).Methods("POST")
	router.HandleFunc("/v1/api/verify/id/media/fetch", verifier.VerifyIDMediaFetch).Methods("GET")
	router.HandleFunc("/v1/api/verify/id/success", verifier.VerifyIDSuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Card
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/phone", verifier.PhoneVerifyInitHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Card
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/card", verifier.CCVerifyInitHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/card/check", verifier.CCVerifyCheckHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/card/success", verifier.CCVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify - Address
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/verify/address", verifier.AddressVerifyInitHandler).Methods("POST")
	router.HandleFunc("/v1/api/verify/address/success", verifier.AddressVerifySuccessHandler).Methods("POST")
	//||------------------------------------------------------------------------------------------------||
	//|| Tools
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/currency", utils.ConvertUSDHandler).Methods("GET")
	router.HandleFunc("/health", utils.HealthHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Agent
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/agent/fatal", agent.AgentFatalHandler).Methods("POST")
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
