package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Deny OAuth Handler
//||------------------------------------------------------------------------------------------------||

func DenyOAuthHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Get
	//||------------------------------------------------------------------------------------------------||

	oauth := r.URL.Query().Get("oauth")
	if oauth == "" {
		responses.ErrorHTML(w, "Missing OAuth session ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch session from Redis
	//||------------------------------------------------------------------------------------------------||

	val, err := app.CacheRedis["oauth"].Get("oauth:" + oauth)
	if err != nil {
		responses.ErrorHTML(w, "Invalid or expired OAuth session")
		return
	}

	var session OAuthSession
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
	app.CacheRedis["oauth"].Set("oauth:"+oauth, data, ttl)

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
	if session.State != "" {
		q.Set("state", session.State)
	}
	if len(session.Scope) > 0 {
		q.Set("scope", strings.Join(session.Scope, ","))
	}
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
