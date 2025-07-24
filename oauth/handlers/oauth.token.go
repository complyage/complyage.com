package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/interfaces"
	"base/responses"
	"context"
	"encoding/json"
	"net/http"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Token Exchange Handler
//|| OAuth2-compliant: Accepts code, returns access token JSON
//||------------------------------------------------------------------------------------------------||

func TokenExchangeHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Form and Required Fields
	//||------------------------------------------------------------------------------------------------||

	r.ParseForm()
	code := r.FormValue("code")
	clientID := r.FormValue("client_id")

	if code == "" || clientID == "" {
		responses.ErrorJSON(w, "Missing required fields: code and client_id", 400)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load OAuth Access from Redis
	//||------------------------------------------------------------------------------------------------||

	val, err := db.Redis.Get(ctx, "access:"+code).Result()
	if err != nil {
		responses.ErrorJSON(w, "Invalid or expired code", 400)
		return
	}

	var access interfaces.OAuthAccess
	if err := json.Unmarshal([]byte(val), &access); err != nil {
		responses.ErrorJSON(w, "Malformed access session", 500)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate client ID matches
	//||------------------------------------------------------------------------------------------------||

	if access.APIKey != clientID {
		responses.ErrorJSON(w, "Client ID mismatch", 403)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Optionally delete code to enforce one-time use
	//||------------------------------------------------------------------------------------------------||

	_ = db.Redis.Del(ctx, "access:"+code)

	//||------------------------------------------------------------------------------------------------||
	//|| Build and Return Access Token Response
	//||------------------------------------------------------------------------------------------------||

	resp := map[string]interface{}{
		"access_token": access.AccessKey,
		"token_type":   "bearer",
		"expires_in":   int(access.Expires - time.Now().Unix()),
		"scope":        strings.Join(access.Scope, " "),
	}

	responses.SuccessJSON(w, resp)
}
