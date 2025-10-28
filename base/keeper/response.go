package keeper

import (
	"encoding/json"
	"net/http"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Set ComplyAge Cookie
//||------------------------------------------------------------------------------------------------||

func setComplyAgeCookie(w http.ResponseWriter, keeperId string) {
	http.SetCookie(w, &http.Cookie{
		Name:  "complyage_session",
		Value: keeperId,
		Path:  "/",
		// Domain:   GATE_DOMAIN, // uncomment if you want cross-subdomain cookies
		Expires:  time.Now().Add(24 * time.Hour),
		Secure:   false, // true if HTTPS
		HttpOnly: false, // must be false so JS can read it
		SameSite: http.SameSiteLaxMode,
	})
}

//||------------------------------------------------------------------------------------------------||
//|| Respond with JSON
//||------------------------------------------------------------------------------------------------||

func (record *KeeperRecord) Respond(w http.ResponseWriter) {
	setComplyAgeCookie(w, record.KeeperId)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(record)
}

//||------------------------------------------------------------------------------------------------||
//|| Respond with Redirect
//||------------------------------------------------------------------------------------------------||

func (record *KeeperRecord) Redirect(w http.ResponseWriter, r *http.Request) {
	setComplyAgeCookie(w, record.KeeperId)
	returnURL := record.ReturnURL
	if returnURL == "" {
		returnURL = "/" // fallback if nothing provided
	}
	http.Redirect(w, r, returnURL, http.StatusFound) // 302 redirect
}
