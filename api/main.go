package main

import (
	"api/handlers"
	"base/db"
	"base/helpers"
	"base/loaders"
	"base/models"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
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
	err := godotenv.Load("../base/.env")
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
	//|| Open DB Connection
	//||------------------------------------------------------------------------------------------------||
	db.ConnectMySQL()
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to Redis
	//||------------------------------------------------------------------------------------------------||
	db.ConnectRedis() // Redis
	//||------------------------------------------------------------------------------------------------||
	//|| GORM
	//||------------------------------------------------------------------------------------------------||
	db.DB.AutoMigrate(&models.Account{})
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
	//|| Global Routes
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/api/client", handlers.CheckClientEnforcement).Methods("GET")
	router.HandleFunc("/v1/api/oauth", handlers.OAuthResponseHandler).Methods("GET")
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
	router.HandleFunc("/auth/twofactor", handlers.TwoFactorHandler).Methods("POST")
	router.HandleFunc("/auth/me", handlers.AuthMeHandler).Methods("GET")
	router.HandleFunc("/auth/complete", handlers.CompleteHandler).Methods("POST", "GET")
	router.HandleFunc("/auth/login", handlers.LoginHandler).Methods("POST")
	router.HandleFunc("/auth/logout", handlers.LogoutHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Simple Up Check
	//||------------------------------------------------------------------------------------------------||
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
