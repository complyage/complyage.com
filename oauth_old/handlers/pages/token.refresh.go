package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"

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

	val, err := app.CacheRedis["oauth"].Get("refresh:" + refreshToken)
	if err != nil {
		responses.ErrorJSON(w, "Invalid or expired refresh token", http.StatusUnauthorized)
		return
	}

	var session OAuthSession
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

	access := OAuthAccess{
		AccountID:     session.AccountID,
		APIKey:        session.ClientID,
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

	err = app.CacheRedis["oauth"].Set("access:"+newAccessKey, accessData, 60*time.Minute)
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
	app.CacheRedis["oauth"].Del("refresh:" + refreshToken)
	err = app.CacheRedis["oauth"].Set("refresh:"+newRefreshKey, refreshData, 30*24*time.Hour)
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
