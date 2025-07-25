package main

import (
	"base/db"
	"base/helpers"
	"base/loaders"
	"fmt"
	"log"
	"net/http"
	"oauth/handlers"
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
	//|| Preload Translations
	//||------------------------------------------------------------------------------------------------||
	loaders.LoadTranslations()
	//||------------------------------------------------------------------------------------------------||
	//|| Setup Router
	//||------------------------------------------------------------------------------------------------||
	router := mux.NewRouter()
	//||------------------------------------------------------------------------------------------------||
	//|| Static Assets
	//||------------------------------------------------------------------------------------------------||
	fs := http.FileServer(http.Dir("public"))
	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))
	//||------------------------------------------------------------------------------------------------||
	//|| Middleware
	//||------------------------------------------------------------------------------------------------||
	router.Use(LoggerMiddleware)
	router.HandleFunc("/cli/translate", helpers.TranslateI18nHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Global Routes
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/v1/authorize", handlers.ServeOAuthHandler).Methods("GET")
	router.HandleFunc("/v1/deny", handlers.DenyOAuthHandler).Methods("GET")
	router.HandleFunc("/v1/approve", handlers.ApproveOAuthHandler).Methods("GET")
	router.HandleFunc("/v1/private", handlers.ServePrivateKeyForm).Methods("GET")
	router.HandleFunc("/v1/private/submit", handlers.SubmitPrivateKeyHandler).Methods("GET")
	//||------------------------------------------------------------------------------------------------||
	//|| Simple Up Check
	//||------------------------------------------------------------------------------------------------||
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintf(
			w,
			`{"status":"ok","time":"%s"}`,
			time.Now().Format(time.RFC3339),
		)
	}).Methods("GET")
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
	fmt.Println("OAUTH server running on :" + os.Getenv("PORT_HTTP_OAUTH"))
	log.Fatal(
		http.ListenAndServe(
			":"+os.Getenv("PORT_HTTP_OAUTH"),
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
