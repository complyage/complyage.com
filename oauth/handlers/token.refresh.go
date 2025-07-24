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
	"time"

	"github.com/google/uuid"
)

//||------------------------------------------------------------------------------------------------||
//|| Refresh Token Handler
//||------------------------------------------------------------------------------------------------||

func RefreshTokenHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Form
	//||------------------------------------------------------------------------------------------------||

	err := r.ParseForm()
	if err != nil {
		responses.ErrorJSON(w, "Invalid form submission", http.StatusBadRequest)
		return
	}

	refreshToken := r.FormValue("refresh_token")
	if refreshToken == "" {
		responses.ErrorJSON(w, "Missing refresh_token", http.StatusBadRequest)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Lookup Refresh Token in Redis
	//||------------------------------------------------------------------------------------------------||

	ctx := context.Background()
	val, err := db.Redis.Get(ctx, "refresh:"+refreshToken).Result()
	if err != nil {
		responses.ErrorJSON(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	var session interfaces.OAuthSession
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		responses.ErrorJSON(w, "Failed to parse refresh token session", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate New Access Token
	//||------------------------------------------------------------------------------------------------||

	newAccessKey := uuid.NewString()
	session.AccessKey = newAccessKey
	session.Expires = time.Now().Unix() + 3600
	session.Created = time.Now().Unix()

	access := interfaces.OAuthAccess{
		AccountID:     session.AccountID,
		APIKey:        session.APIKey,
		AccessKey:     session.AccessKey,
		State:         session.State,
		Scope:         session.Scope,
		Expires:       session.Expires,
		Created:       session.Created,
		Status:        "APPR",
		Verifications: session.Verifications,
	}

	accessData, err := json.Marshal(access)
	if err != nil {
		responses.ErrorJSON(w, "Failed to encode access data", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save New Access Key to Redis
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(ctx, "access:"+newAccessKey, accessData, 60*time.Minute).Err()
	if err != nil {
		responses.ErrorJSON(w, "Failed to store access token", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Optionally Rotate Refresh Token
	//||------------------------------------------------------------------------------------------------||

	newRefreshKey := uuid.NewString()
	refreshData, _ := json.Marshal(session)

	// Delete old and write new
	db.Redis.Del(ctx, "refresh:"+refreshToken)
	err = db.Redis.Set(ctx, "refresh:"+newRefreshKey, refreshData, 30*24*time.Hour).Err()
	if err != nil {
		responses.ErrorJSON(w, "Failed to rotate refresh token", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Respond With New Tokens
	//||------------------------------------------------------------------------------------------------||

	responses.SuccessJSON(w, map[string]interface{}{
		"access_token":  newAccessKey,
		"token_type":    "bearer",
		"expires_in":    3600,
		"refresh_token": newRefreshKey,
	})
}
