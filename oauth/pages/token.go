package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
	"oauth/access"
	"oauth/templates"

	"github.com/complyage/base/keeper"
	"github.com/ralphferrara/aria/base/template"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type TokenRequest struct {
	ClientId string `json:"client_id"`,
	ClientSecret     string `json:"client_secret"`,
	AuthorizationCode string `json:"authorization_code"`,
}

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

type TokenResponse struct {
	AccessToken string `json:"access_token"`,
}

//||------------------------------------------------------------------------------------------------||
//|| Approve OAuth Handler
//||------------------------------------------------------------------------------------------------||

func TokenOAuthHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Context. OAuth
	//||------------------------------------------------------------------------------------------------||

	sessionId := r.URL.Query().Get("session")
	key := r.URL.Query().Get("key")
	if sessionId == "" {
		templates.ErrorHTML(r, w, "Missing OAuth session ID")
		return
	}
	if key == "" {
		templates.ErrorHTML(r, w, "Missing Action key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Redis Session
	//||------------------------------------------------------------------------------------------------||

	oa, err := access.LoadAccess(sessionId)
	if err != nil {
		templates.ErrorHTML(r, w, "LoadAccess: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Key
	//||------------------------------------------------------------------------------------------------||

	if oa.BypassKey == "" || key != oa.BypassKey {
		templates.ErrorHTML(r, w, "Invalid Bypass key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Delete old oauth key and save new access key
	//||------------------------------------------------------------------------------------------------||

	_ = oa.RemoveAccess()

	//||------------------------------------------------------------------------------------------------||
	//|| Delete old oauth key and save new access key
	//||------------------------------------------------------------------------------------------------||

	keep, err := keeper.Load(r)
	if err != nil {
		templates.ErrorHTML(r, w, "Load Keeper: "+err.Error())
		return
	}
	keep.Verified = true
	keep.Status = "BYPASS"
	keep.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	tpl := template.Create("verified")
	tpl.Add("RETURN_URL", keep.ReturnURL)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl.Compile()))
}
