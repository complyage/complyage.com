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
	"base/template"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/google/uuid"
)

//||------------------------------------------------------------------------------------------------||
//|| Serve
//||------------------------------------------------------------------------------------------------||

func ServePrivateKeyForm(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Params
	//||------------------------------------------------------------------------------------------------||

	clientID := r.URL.Query().Get("client_id")
	state := r.URL.Query().Get("state")

	if clientID == "" {
		responses.ErrorHTML(w, "client_id is required")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Site
	//||------------------------------------------------------------------------------------------------||

	site := loaders.GetSiteByPublic(clientID)
	if site == nil {
		responses.ErrorHTML(w, "Invalid apiKey")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Scopes
	//||------------------------------------------------------------------------------------------------||

	var requestedScopes []string
	for _, p := range strings.Split(site.SitePermissions, ",") {
		if s := strings.TrimSpace(p); s != "" {
			requestedScopes = append(requestedScopes, s)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session
	//||------------------------------------------------------------------------------------------------||

	var session interfaces.SessionRecord
	if cookie, err := r.Cookie("session"); err == nil {
		session, err = helpers.FetchSession(cookie.Value)
		if err != nil {
			session = interfaces.SessionRecord{ID: 0, Username: "Anonymous", Status: "RMVD"}
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| OAuth Session
	//||------------------------------------------------------------------------------------------------||

	referenceKey := uuid.NewString()
	siteRedirect := site.SiteRedirect
	if siteRedirect == "" {
		siteRedirect = site.SiteURL + "/oauth/complete"
	}

	oauthSession := interfaces.OAuthSession{
		AccountID:   session.ID,
		Private:     session.Private,
		PrivateHash: session.PrivateHash,
		ClientID:    clientID,
		AccessKey:   uuid.NewString(),
		State:       state,
		Redirect:    siteRedirect,
		Scope:       requestedScopes,
		Expires:     time.Now().Unix() + 3600,
		Created:     time.Now().Unix(),
		Status:      "PEND",
	}

	sessionData, _ := json.Marshal(oauthSession)
	_ = db.Redis.Set(r.Context(), "oauth:"+referenceKey, sessionData, 60*time.Minute).Err()

	//||------------------------------------------------------------------------------------------------||
	//|| Template
	//||------------------------------------------------------------------------------------------------||

	tpl := template.Create("private").
		Add("SITE_URL", site.SiteURL).
		Add("SITE_NAME", site.SiteName).
		Add("APIKEY", clientID).
		Add("OAUTHAPPR", os.Getenv("VITE_COMPLYAGE_OAUTH_URL")+"/v1/approve?oauth="+referenceKey).
		Add("LOGINSTATUS", func() string {
			if session.Level > 0 && session.Status == "ACTV" {
				return "loggedin"
			}
			return "loggedout"
		}()).
		Add("USERNAME", session.Username).
		Add("USERLEVEL", fmt.Sprintf("%d", session.Level)).
		Add("COMPLYAGE_UI_URL", os.Getenv("SITE_URL")).
		Add("COMPLYAGE_CLIENT_URL", os.Getenv("LOCAL_URL"))

	//||------------------------------------------------------------------------------------------------||
	//|| Modal
	//||------------------------------------------------------------------------------------------------||

	sub := template.Create("private_locked")
	subHTML, err := sub.Compile()
	if err != nil {
		responses.ErrorHTML(w, "Template rendering failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Submodal
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("SUBMODAL", subHTML)

	//||------------------------------------------------------------------------------------------------||
	//|| Translate
	//||------------------------------------------------------------------------------------------------||

	tpl = tpl.Translate(r)

	//||------------------------------------------------------------------------------------------------||
	//|| Compile and Output
	//||------------------------------------------------------------------------------------------------||

	html, err := tpl.Compile()
	if err != nil {
		responses.ErrorHTML(w, "Template rendering failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| HTML Output
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
