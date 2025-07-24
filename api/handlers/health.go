package handlers

import (
	"fmt"
	"net/http"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Health Handler
//||------------------------------------------------------------------------------------------------||

func HealthHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(
		w,
		`{"status":"ok","time":"%s"}`,
		time.Now().Format(time.RFC3339),
	)
}
