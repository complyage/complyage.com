package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/sites"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/auth/types"
	"github.com/ralphferrara/aria/base/template"
	"github.com/ralphferrara/aria/responses"

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

	site, err := sites.FetchSiteByPublic(clientID)
	if err != nil {
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

	var session types.SessionRecord
	if cookie, err := r.Cookie("session"); err == nil {
		session, err = actions.FetchSession(cookie.Value)
		if err != nil {
			session = types.SessionRecord{ID: 0, Username: "Anonymous", Status: "RMVD"}
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

	oauthSession := OAuthSession{
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
	_ = app.CacheRedis["oauth"].Set("oauth:"+referenceKey, sessionData, 60*time.Minute)

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
