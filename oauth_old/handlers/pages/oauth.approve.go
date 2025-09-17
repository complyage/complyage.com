package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Approve OAuth Handler
//||------------------------------------------------------------------------------------------------||

func ApproveOAuthHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Context. OAuth
	//||------------------------------------------------------------------------------------------------||

	oauth := r.URL.Query().Get("oauth")
	if oauth == "" {
		responses.ErrorHTML(w, "Missing OAuth session ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Redis Session
	//||------------------------------------------------------------------------------------------------||

	val, err := app.CacheRedis["oauth"].Get("oauth:" + oauth)
	if err != nil {
		responses.ErrorHTML(w, "Invalid or expired OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Session
	//||------------------------------------------------------------------------------------------------||

	var session OAuthSession
	if err := json.Unmarshal([]byte(val), &session); err != nil {
		responses.ErrorHTML(w, "Failed to parse OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| If we don't have the key or cannot verify it, redirect to /private
	//||------------------------------------------------------------------------------------------------||

	if session.Private == "" || session.PrivateHash == "" || encrypt.CheckPrivateKey(session.Private, session.PrivateHash) != nil {
		http.Redirect(w, r, "/v1/private?oauth="+oauth, http.StatusSeeOther)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load the Verifications
	//||------------------------------------------------------------------------------------------------||

	var verifications []OAuthVerification
	err = app.SQLDB["main"].DB.
		Where("fid_account = ?", session.AccountID).
		Where("verification_type IN ?", session.Scope).
		Where("verification_status = ?", "APPR").
		Find(&verifications).Error

	if err != nil {
		responses.ErrorHTML(w, "Failed to load verifications")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check if all required verifications are present
	//||------------------------------------------------------------------------------------------------||

	if (len(session.Scope) > 0 && len(verifications) < len(session.Scope)) ||
		(len(session.Scope) == 0 && len(verifications) == 0) {
		responses.ErrorHTML(w, "Missing required verifications")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decrypt each verification
	//||------------------------------------------------------------------------------------------------||

	for i := range verifications {
		raw, ok := verifications[i].Data.(string)
		if !ok {
			responses.ErrorHTML(w, "Invalid verification data format")
			return
		}
		decrypted, pErr := encrypt.DecryptWithPrivateKey([]byte(raw), session.Private)
		if pErr != nil {
			responses.ErrorHTML(w, "Failed to decrypt verification data")
			return
		}
		verifications[i].Data = string(decrypted)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create OAuthAccess Payload
	//||------------------------------------------------------------------------------------------------||

	access := OAuthAccess{
		AccountID:     session.AccountID,
		APIKey:        session.ClientID,
		AccessKey:     session.AccessKey,
		State:         session.State,
		Scope:         session.Scope,
		Expires:       session.Expires,
		Created:       session.Created,
		Status:        "APPR",
		Verifications: verifications,
	}

	accessData, err := json.Marshal(access)
	if err != nil {
		responses.ErrorHTML(w, "Failed to encode access payload")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Delete old oauth key and save new access key
	//||------------------------------------------------------------------------------------------------||

	_ = app.CacheRedis["oauth"].Del("oauth:" + oauth)
	err = app.CacheRedis["oauth"].Set("access:"+session.AccessKey, accessData, 60*time.Minute)
	if err != nil {
		responses.ErrorHTML(w, "Failed to store access session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Redirect
	//||------------------------------------------------------------------------------------------------||

	redirectURL := session.Redirect
	if redirectURL == "" {
		redirectURL = "/"
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		responses.ErrorHTML(w, "Invalid redirect URI")
		return
	}

	query := parsedURL.Query()
	query.Set("code", session.AccessKey)
	if session.State != "" {
		query.Set("state", session.State)
	}
	parsedURL.RawQuery = query.Encode()

	//||------------------------------------------------------------------------------------------------||
	//|| Redirect to original client with query params
	//||------------------------------------------------------------------------------------------------||

	http.Redirect(w, r, parsedURL.String(), http.StatusSeeOther)

}
