package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| OAuthRedirectHandler – Handles /redirect after login
//||------------------------------------------------------------------------------------------------||

func OAuthReturnHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the oauth query parameter
	//||------------------------------------------------------------------------------------------------||

	oauthKey := r.URL.Query().Get("oauth")
	if oauthKey == "" {
		responses.Error(w, http.StatusBadRequest, "Missing oauth parameter")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load OAuth Session from Redis
	//||------------------------------------------------------------------------------------------------||

	sessionData, err := db.Redis.Get(r.Context(), "oauth:"+oauthKey).Result()
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "OAuth session not found or expired")
		return
	}

	var oauthSession interfaces.OAuthSession
	if err := json.Unmarshal([]byte(sessionData), &oauthSession); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to decode OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check expiration
	//||------------------------------------------------------------------------------------------------||

	if time.Now().Unix() > oauthSession.Expires {
		responses.Error(w, http.StatusUnauthorized, "OAuth session expired")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Redirect URL with required query params
	//||------------------------------------------------------------------------------------------------||

	redirectURL, err := url.Parse("/v1/authorize")
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Parsing the URL")
		return
	}
	q := redirectURL.Query()
	q.Set("client_id", oauthSession.ClientID)
	q.Set("state", oauthSession.State)
	q.Set("scope", strings.Join(oauthSession.Scope, ","))
	redirectURL.RawQuery = q.Encode()

	//||------------------------------------------------------------------------------------------------||
	//|| Redirect to the client site
	//||------------------------------------------------------------------------------------------------||

	http.Redirect(w, r, redirectURL.String(), http.StatusSeeOther)
}
