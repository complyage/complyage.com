package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"sync"
	"time"

	"github.com/joho/godotenv"
)

//||---------------------------------------------------------------------------------------------||
//|| Rate Limiter
//||---------------------------------------------------------------------------------------------||

type rateLimiter struct {
	mu     sync.Mutex
	tokens map[string]time.Time
}

func newRateLimiter() *rateLimiter {
	return &rateLimiter{tokens: make(map[string]time.Time)}
}

func (r *rateLimiter) allow(ip string, interval time.Duration) bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	last, ok := r.tokens[ip]
	if !ok || time.Since(last) > interval {
		r.tokens[ip] = time.Now()
		return true
	}
	return false
}

//||---------------------------------------------------------------------------------------------||
//|| Main
//||---------------------------------------------------------------------------------------------||

func main() {

	//||---------------------------------------------------------------------------------------------||
	//|| Load .env
	//||---------------------------------------------------------------------------------------------||

	_ = godotenv.Load("../.env")

	//||---------------------------------------------------------------------------------------------||
	//|| Get Deploy Key
	//||---------------------------------------------------------------------------------------------||

	deployKey := os.Getenv("DEPLOY_KEY")
	if deployKey == "" {
		log.Fatal("DEPLOY_KEY not set in environment")
	}

	limiter := newRateLimiter()

	//||---------------------------------------------------------------------------------------------||
	//|| Handler
	//||---------------------------------------------------------------------------------------------||

	http.HandleFunc("/deploy", func(w http.ResponseWriter, r *http.Request) {

		//||---------------------------------------------------------------------------------------------||
		//|| Get IP
		//||---------------------------------------------------------------------------------------------||

		ip, _, _ := net.SplitHostPort(r.RemoteAddr)

		//||---------------------------------------------------------------------------------------------||
		//|| 1 Request Rate Limiting
		//||---------------------------------------------------------------------------------------------||

		if !limiter.allow(ip, 5*time.Second) {
			http.Error(w, "too many requests", http.StatusTooManyRequests)
			log.Printf("Rate limited IP: %s\n", ip)
			return
		}

		//||---------------------------------------------------------------------------------------------||
		//|| Token Check
		//||---------------------------------------------------------------------------------------------||

		key := r.URL.Query().Get("check")
		if key != deployKey {
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			log.Printf("Unauthorized deploy attempt from %s\n", r.RemoteAddr)
			return
		}

		HandleDeployment(w, r)
	})

	//||---------------------------------------------------------------------------------------------||
	//|| Listen
	//||---------------------------------------------------------------------------------------------||

	log.Println("Listening on :9090...")
	log.Fatal(http.ListenAndServe(":9090", nil))

}
