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
	"net/url"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Deny OAuth Handler
//||------------------------------------------------------------------------------------------------||

func DenyOAuthHandler(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	oauth := r.URL.Query().Get("oauth")

	if oauth == "" {
		responses.ErrorHTML(w, "Missing OAuth session ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch session from Redis
	//||------------------------------------------------------------------------------------------------||

	val, err := db.Redis.Get(ctx, "oauth:"+oauth).Result()
	if err != nil {
		responses.ErrorHTML(w, "Invalid or expired OAuth session")
		return
	}

	var session interfaces.OAuthSession
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		responses.ErrorHTML(w, "Failed to parse OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update status and store again
	//||------------------------------------------------------------------------------------------------||

	session.Status = "denied"
	data, _ := json.Marshal(session)
	ttl := time.Until(time.Unix(session.Expires, 0))
	db.Redis.Set(ctx, "oauth:"+oauth, data, ttl)

	//||------------------------------------------------------------------------------------------------||
	//|| Build redirect with query params
	//||------------------------------------------------------------------------------------------------||

	redirectURL := session.Redirect
	if redirectURL == "" {
		redirectURL = "/" // fallback
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build redirect with query params
	//||------------------------------------------------------------------------------------------------||

	q := url.Values{}
	q.Set("state", session.State)
	q.Set("scope", strings.Join(session.Scope, ","))
	q.Set("accessKey", session.AccessKey)
	q.Set("status", session.Status)

	finalURL := redirectURL
	if strings.Contains(redirectURL, "?") {
		finalURL += "&" + q.Encode()
	} else {
		finalURL += "?" + q.Encode()
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Send it
	//||------------------------------------------------------------------------------------------------||

	http.Redirect(w, r, finalURL, http.StatusFound)
}
