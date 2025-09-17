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
	"github.com/ralphferrara/aria/locale"
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
	for _, p := range strings.Split(site.Permissions, ",") {
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
	siteRedirect := site.Redirect
	if siteRedirect == "" {
		siteRedirect = site.URL + "/oauth/complete"
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

	tpl := template.Create("private", locale.Request(r))
	tpl.Add("SITE_URL", site.URL)
	tpl.Add("SITE_NAME", site.Name)
	tpl.Add("APIKEY", clientID)
	tpl.Add("OAUTHAPPR", os.Getenv("VITE_COMPLYAGE_OAUTH_URL")+"/v1/approve?oauth="+referenceKey)
	tpl.Add("LOGINSTATUS", func() string {
		if session.Level > 0 && session.Status == "ACTV" {
			return "loggedin"
		}
		return "loggedout"
	}())
	tpl.Add("USERNAME", session.Username)
	tpl.Add("USERLEVEL", fmt.Sprintf("%d", session.Level))
	tpl.Add("COMPLYAGE_UI_URL", os.Getenv("SITE_URL"))
	tpl.Add("COMPLYAGE_CLIENT_URL", os.Getenv("LOCAL_URL"))

	//||------------------------------------------------------------------------------------------------||
	//|| Modal
	//||------------------------------------------------------------------------------------------------||

	sub := template.Create("private_locked", locale.Request(r))
	subHTML := sub.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| Submodal
	//||------------------------------------------------------------------------------------------------||

	tpl.Add("SUBMODAL", subHTML)

	//||------------------------------------------------------------------------------------------------||
	//|| Compile and Output
	//||------------------------------------------------------------------------------------------------||

	html := tpl.Compile()

	//||------------------------------------------------------------------------------------------------||
	//|| HTML Output
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(html))
}
