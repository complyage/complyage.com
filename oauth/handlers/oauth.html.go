package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/loaders"
	"base/responses"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

// ||------------------------------------------------------------------------------------------------||
// || Serves the HTML file with dynamic replacements
// ||------------------------------------------------------------------------------------------------||

func ServeOAuthHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Serves the HTML file with dynamic replacements
	//||------------------------------------------------------------------------------------------------||

	apiKey := r.URL.Query().Get("client_id")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")

	//||------------------------------------------------------------------------------------------------||
	//|| APIKey
	//||------------------------------------------------------------------------------------------------||

	if apiKey == "" {
		responses.ErrorHTML(w, "client_id is required")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Site By API Key
	//||------------------------------------------------------------------------------------------------||

	site := loaders.GetSiteByPublic(apiKey)
	if site == nil {
		responses.ErrorHTML(w, "Invalid apiKey")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Scope (default to site's allowed if missing)
	//||------------------------------------------------------------------------------------------------||

	var requestedScopes []string
	var siteScopes []string

	//||------------------------------------------------------------------------------------------------||
	//|| Clean Scope Requests
	//||------------------------------------------------------------------------------------------------||

	for _, p := range strings.Split(site.SitePermissions, ",") {
		p = strings.TrimSpace(p)
		if p != "" {
			siteScopes = append(siteScopes, p)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Handle Scope
	//||------------------------------------------------------------------------------------------------||

	if scope == "" {
		requestedScopes = siteScopes
	} else {
		scopeParts := strings.Split(scope, ",")
		for _, s := range scopeParts {
			s = strings.TrimSpace(strings.ToUpper(s))
			if s == "" {
				continue
			}

			//||------------------------------------------------------------------------------------------------||
			//|| Check if this scope is in allowed site permissions
			//||------------------------------------------------------------------------------------------------||

			found := false
			for _, allowed := range siteScopes {
				if strings.EqualFold(s, allowed) {
					found = true
					break
				}
			}

			//||------------------------------------------------------------------------------------------------||
			//|| An unapproved scope was requested
			//||------------------------------------------------------------------------------------------------||

			if !found {
				responses.ErrorHTML(w, fmt.Sprintf("Invalid scope requested: %s", s))
				return
			}

			requestedScopes = append(requestedScopes, s)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	var session interfaces.SessionRecord

	cookie, err := r.Cookie("session")
	if err == nil {
		session, err = helpers.FetchSession(cookie.Value)
		if err != nil {
			session = interfaces.SessionRecord{
				ID:            "0",
				Username:      "Anonymous",
				Status:        "RMVD",
				Type:          "USER",
				Private:       "",
				PrivateCheck:  "",
				Level:         0,
				Verifications: []interfaces.UserVerification{},
			}
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serves the HTML file with dynamic replacements
	//||------------------------------------------------------------------------------------------------||

	lang := "en"
	if al := r.Header.Get("Accept-Language"); al != "" {
		primary := strings.SplitN(al, ",", 2)[0]
		if code := strings.SplitN(primary, "-", 2)[0]; code != "" {
			lang = strings.ToLower(code)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Translations
	//||------------------------------------------------------------------------------------------------||

	translations, _ := loaders.GetTranslations(lang)

	//||------------------------------------------------------------------------------------------------||
	//|| Get The Template
	//||------------------------------------------------------------------------------------------------||

	tpl := ""
	b, err := ioutil.ReadFile("assets/oauth.html")
	if err != nil {
		http.Error(w, "template not found", http.StatusInternalServerError)
		fmt.Printf("read template error: %v\n", err)
		return
	}
	tpl = string(b)

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Template  Markers
	//||------------------------------------------------------------------------------------------------||

	vars := make(map[string]string, len(translations)+5)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Permission HTML Template
	//||------------------------------------------------------------------------------------------------||

	permissionHTMLBytes, err := os.ReadFile("assets/permission.html")
	if err != nil {
		responses.ErrorHTML(w, "Failed to load permission.html: "+err.Error())
		return
	}
	permissionHTML := string(permissionHTMLBytes)

	//||------------------------------------------------------------------------------------------------||
	//|| Permission Counts
	//||------------------------------------------------------------------------------------------------||

	totalPermissions := 0
	matchPermissions := 0

	//||------------------------------------------------------------------------------------------------||
	//|| Loop through the Requested Permissions
	//||------------------------------------------------------------------------------------------------||

	var htmlPermissions strings.Builder
	for _, perm := range requestedScopes {

		//||------------------------------------------------------------------------------------------------||
		//|| Figure out Permissions and Add to Total Count
		//||------------------------------------------------------------------------------------------------||

		perm = strings.TrimSpace(perm)
		if perm == "" {
			continue
		}
		totalPermissions++

		//||------------------------------------------------------------------------------------------------||
		//|| Generate Permission Based on Session & Verification Status
		//||------------------------------------------------------------------------------------------------||

		statusClass := "loggedout"
		if session.Level > 0 && session.Status == "ACTV" {

			//||------------------------------------------------------------------------------------------------||
			//|| Unverified until Verified
			//||------------------------------------------------------------------------------------------------||

			statusClass = "unverified"

			//||------------------------------------------------------------------------------------------------||
			//|| Loop Through Users Verifications
			//||------------------------------------------------------------------------------------------------||

			for _, v := range session.Verifications {
				if strings.EqualFold(v.Type, perm) && strings.EqualFold(v.Status, "VERF") {
					statusClass = "verified"
					matchPermissions++
					break
				}
			}
		}

		//||------------------------------------------------------------------------------------------------||
		//|| Permissions Template Markers
		//||------------------------------------------------------------------------------------------------||

		html := strings.ReplaceAll(permissionHTML, "[%%PERMCODE%%]", perm)
		html = strings.ReplaceAll(html, "[%%PERM_STATUS_"+perm+"%%]", statusClass)

		//||------------------------------------------------------------------------------------------------||
		//|| Add to Main HTML
		//||------------------------------------------------------------------------------------------------||

		htmlPermissions.WriteString(html)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate and Store OAuth Session in Redis
	//||------------------------------------------------------------------------------------------------||

	referenceKey := uuid.NewString()

	//||------------------------------------------------------------------------------------------------||
	//|| Redirect URL
	//||------------------------------------------------------------------------------------------------||

	siteRedirect := site.SiteRedirect
	if siteRedirect == "" {
		siteRedirect = site.SiteURL + "/oauth/complete"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Write OAuth Session to Redis
	//||------------------------------------------------------------------------------------------------||

	oauthSession := interfaces.OAuthSession{
		AccountID:    session.ID,
		Private:      session.Private,
		PrivateCheck: session.PrivateCheck,
		APIKey:       apiKey,
		AccessKey:    uuid.NewString(),
		State:        state,
		Redirect:     siteRedirect,
		Scope:        requestedScopes,
		Expires:      time.Now().Unix() + 3600,
		Created:      time.Now().Unix(),
		Status:       "PEND",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Convert to String
	//||------------------------------------------------------------------------------------------------||

	sessionData, err := json.Marshal(oauthSession)
	if err != nil {
		responses.ErrorHTML(w, "Failed to encode OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Write OAuth Session to Redis
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(r.Context(), "oauth:"+referenceKey, sessionData, 60*time.Minute).Err()
	if err != nil {
		responses.ErrorHTML(w, "Failed to store OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add the Permission to the Template
	//||------------------------------------------------------------------------------------------------||

	tpl = strings.ReplaceAll(tpl, "[%%PERMISSIONS%%]", htmlPermissions.String())

	//||------------------------------------------------------------------------------------------------||
	//|| Add Translations
	//||------------------------------------------------------------------------------------------------||

	for key, val := range translations {
		vars[key] = val
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Match Status on Permissions
	//||------------------------------------------------------------------------------------------------||

	vars["VERIFICATIONSTATUS"] = ""
	if session.Level > 0 && session.Status == "ACTV" {
		vars["VERIFICATIONSTATUS"] = "matched"
		if (matchPermissions == totalPermissions) || (totalPermissions == 0) {
			vars["VERIFICATIONSTATUS"] = "matched"
		}
	}

	vars["REMAININGCOUNT"] = fmt.Sprintf("%d", totalPermissions-matchPermissions)

	//||------------------------------------------------------------------------------------------------||
	//|| Do we need to request the private key?
	//||------------------------------------------------------------------------------------------------||

	pErr := helpers.CheckPrivateKey(session.Private, session.PrivateCheck)
	if pErr != nil || session.Private == "" {
		vars["OAUTHAPPR"] = os.Getenv("VITE_COMPLYAGE_OAUTH_URL") + "/v1/private?oauth=" + referenceKey
	} else {
		vars["OAUTHAPPR"] = os.Getenv("VITE_COMPLYAGE_OAUTH_URL") + "/v1/approve?oauth=" + referenceKey
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Translations
	//||------------------------------------------------------------------------------------------------||

	vars["LOGINURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/login?oauth=" + referenceKey
	vars["SIGNUPURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/signup?oauth=" + referenceKey
	vars["OAUTHDENY"] = os.Getenv("VITE_COMPLYAGE_OAUTH_URL") + "/v1/deny?oauth=" + referenceKey
	vars["EXITURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/exit"
	vars["ACCOUNTURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/members/"
	vars["VERIFYURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/members/verify?oauth=" + referenceKey
	vars["VERIFYALLURL"] = os.Getenv("VITE_COMPLYAGE_UI_URL") + "/members/verifyAll?oauth=" + referenceKey

	//||------------------------------------------------------------------------------------------------||
	//|| Logo URL
	//||------------------------------------------------------------------------------------------------||

	hashSecret := os.Getenv("MINIO_HASH")
	host := os.Getenv("VITE_COMPLYAGE_MINIO_URL")
	bucket := os.Getenv("VITE_MINIO_BUCKET")
	hashInput := fmt.Sprintf("%s%d%s", hashSecret, site.IDSite, hashSecret)
	hashBytes := sha256.Sum256([]byte(hashInput))
	hashHex := fmt.Sprintf("%d_%s", site.IDSite, hex.EncodeToString(hashBytes[:]))
	logoURL := fmt.Sprintf("%s/%s/sites/logos/%s.webp", host, bucket, hashHex)

	//||----------------------------------------------------------------------------------------------||
	//|| User / Session
	//||----------------------------------------------------------------------------------------------||

	if session.Level > 0 && session.Status == "ACTV" {
		vars["LOGINSTATUS"] = "loggedin"
		vars["USERNAME"] = session.Username
		vars["USERLEVEL"] = fmt.Sprintf("%d", session.Level)
	} else {
		vars["LOGINSTATUS"] = "loggedout"
		vars["USERNAME"] = ""
		vars["USERLEVEL"] = "0"
	}

	//||----------------------------------------------------------------------------------------------||
	//|| Site URL/Zone
	//||----------------------------------------------------------------------------------------------||

	vars["SITE_URL"] = site.SiteURL
	vars["SITE_NAME"] = site.SiteName
	vars["SITE_LOGO"] = logoURL
	vars["APIKEY"] = apiKey

	//||----------------------------------------------------------------------------------------------||
	//|| Local
	//||----------------------------------------------------------------------------------------------||

	vars["COMPLYAGE_UI_URL"] = os.Getenv("SITE_URL")
	vars["COMPLYAGE_CLIENT_URL"] = os.Getenv("LOCAL_URL")

	//||----------------------------------------------------------------------------------------------||
	//|| Single pass replacement of [%%KEY%%] → value
	//||----------------------------------------------------------------------------------------------||

	for key, val := range vars {
		marker := fmt.Sprintf("[%%%%%s%%%%]", key)
		tpl = strings.ReplaceAll(tpl, marker, val)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl))

}
