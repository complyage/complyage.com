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
)

//||------------------------------------------------------------------------------------------------||
//|| Access Token Data Handler
//||------------------------------------------------------------------------------------------------||

func AccessTokenDataHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query or Authorization Header
	//||------------------------------------------------------------------------------------------------||

	accessKey := r.URL.Query().Get("access_token")
	if accessKey == "" {
		bearer := r.Header.Get("Authorization")
		if len(bearer) > 7 && bearer[:7] == "Bearer " {
			accessKey = bearer[7:]
		}
	}

	if accessKey == "" {
		responses.ErrorJSON(w, "Missing access token", http.StatusBadRequest)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Lookup Redis
	//||------------------------------------------------------------------------------------------------||

	ctx := context.Background()
	val, err := db.Redis.Get(ctx, "access:"+accessKey).Result()
	if err != nil {
		responses.ErrorJSON(w, "Invalid or expired access token", http.StatusUnauthorized)
		return
	}

	var access interfaces.OAuthAccess
	if err := json.Unmarshal([]byte(val), &access); err != nil {
		responses.ErrorJSON(w, "Failed to parse access token", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return User Info
	//||------------------------------------------------------------------------------------------------||

	responses.SuccessJSON(w, map[string]interface{}{
		"userId":        access.AccountID,
		"verifications": access.Verifications,
	})
}
