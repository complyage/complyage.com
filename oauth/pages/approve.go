package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
	"net/url"
	"oauth/access"
	"oauth/shared"
	"oauth/templates"

	"github.com/complyage/base/db/abstract"
)

//||------------------------------------------------------------------------------------------------||
//|| Approve OAuth Handler
//||------------------------------------------------------------------------------------------------||

func ApproveOAuthHandler(w http.ResponseWriter, r *http.Request) {

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

	if oa.AuthorizeKey == "" || key != oa.AuthorizeKey {
		templates.ErrorHTML(r, w, "Invalid Action key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Shared Payload
	//||------------------------------------------------------------------------------------------------||

	shareWrap, sErr := shared.Create(oa.Enforcement)
	if sErr != nil {
		templates.ErrorHTML(r, w, "shared.Create: "+sErr.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Write the Shared Verifications to Data
	//||------------------------------------------------------------------------------------------------||

	sharedTx := abstract.SharedTransaction{
		AccountId: oa.Enforcement.User.ID,
		SiteId:    oa.Enforcement.Site.ID,
	}
	sharedVerifications := abstract.SharedVerifications{}
	for _, v := range oa.Enforcement.Scopes {
		sharedVerifications = append(sharedVerifications, abstract.SharedVerification{
			Type:         v.Code.String(),
			Verification: v.Verification,
		})
	}
	sharedTx.Verifications = sharedVerifications

	//||------------------------------------------------------------------------------------------------||
	//|| Write to the Database
	//||------------------------------------------------------------------------------------------------||

	err = abstract.RegisterShared(sharedTx)
	if err != nil {
		templates.ErrorHTML(r, w, "RegisterShared: Could not store verifications")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Delete old oauth key and save new access key
	//||------------------------------------------------------------------------------------------------||

	_ = oa.RemoveAccess()

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Redirect
	//||------------------------------------------------------------------------------------------------||

	redirectURL := oa.Enforcement.Site.Redirect
	if redirectURL == "" {
		redirectURL = "/"
	}

	parsedURL, err := url.Parse(redirectURL)
	if err != nil {
		templates.ErrorHTML(r, w, "Invalid redirect URI")
		return
	}

	query := parsedURL.Query()
	query.Set("code", shareWrap.Token)
	if oa.Enforcement.State != "" {
		query.Set("state", oa.Enforcement.State)
	}
	parsedURL.RawQuery = query.Encode()

	//||------------------------------------------------------------------------------------------------||
	//|| Redirect to original client with query params
	//||------------------------------------------------------------------------------------------------||

	http.Redirect(w, r, parsedURL.String(), http.StatusSeeOther)

}
