package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
	"oauth/access"
	"oauth/templates"
	"oauth/webhook"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/encrypted"
	"github.com/complyage/base/keeper"
	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/template"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Approve OAuth Handler
//||------------------------------------------------------------------------------------------------||

func VerifiedHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Querystring
	//||------------------------------------------------------------------------------------------------||

	sessionId := r.URL.Query().Get("session")
	key := r.URL.Query().Get("key")
	scope := r.URL.Query().Get("scope")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate
	//||------------------------------------------------------------------------------------------------||

	if sessionId == "" {
		templates.ErrorHTML(r, w, app.Err("OAuth").Code("MISSING_SESSION"))
		return
	}
	if key == "" {
		templates.ErrorHTML(r, w, app.Err("OAuth").Code("MISSING_CHECKKEY"))
		return
	}
	if scope == "" {
		templates.ErrorHTML(r, w, app.Err("OAuth").Code("MISSING_SCOPE"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Redis Session
	//||------------------------------------------------------------------------------------------------||

	oa, err := access.LoadAccess(sessionId)
	if err != nil {
		templates.ErrorHTML(r, w, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Key
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Authorize Key:", oa.AuthorizeKey)
	app.Log.Data("Provided Key:", key)
	if oa.AuthorizeKey == "" || key != oa.AuthorizeKey {
		templates.ErrorHTML(r, w, app.Err("OAuth").Code("INVALID_CHECKKEY"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Keeper Record for Cookie Storage
	//||------------------------------------------------------------------------------------------------||

	keep, err := keeper.Load(r)
	if err != nil {
		templates.ErrorHTML(r, w, "Load Keeper: "+err.Error())
		return
	}
	keep.Enforced = true
	keep.Verified = true
	keep.Age = oa.Enforcement.Age.DOB.Age()
	keep.UserId = oa.Enforcement.User.ID
	keep.IPAddress = oa.Enforcement.Zone.IPAddress
	keep.ClientId = oa.Enforcement.Site.ClientId
	keep.Status = "VERIFIED"
	keep.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Webhook
	//||------------------------------------------------------------------------------------------------||

	wh := webhook.Create(oa.Enforcement, keep, scope)

	for _, s := range oa.Enforcement.Scopes {
		if s.Code.String() == scope {
			switch s.Code {
			case types.DataTypeFACE:
				enc, err := encrypted.LoadFACE(s.Verification, oa.Enforcement.User.Private)
				if err != nil {
					templates.ErrorHTML(r, w, err.Error())
					return
				}
				wh.Selfie = enc.Selfie.Base64
			case types.DataTypeCRCD:
				enc, err := encrypted.LoadCRCD(s.Verification, oa.Enforcement.User.Private)
				if err != nil {
					templates.ErrorHTML(r, w, err.Error())
					return
				}
				wh.TransactionID = enc.TransactionId
				wh.LastFour = enc.LastFour
			case types.DataTypeIDEN:
				enc, err := encrypted.LoadIDEN(s.Verification, oa.Enforcement.User.Private)
				if err != nil {
					templates.ErrorHTML(r, w, err.Error())
					return
				}
				wh.IDFront = enc.Front.Base64
				wh.Selfie = enc.Selfie.Base64
			}
		}
	}

	log.PrettyPrint(wh)
	wh.Call(oa.Enforcement.Site.Webhook)

	//||------------------------------------------------------------------------------------------------||
	//|| Write the Shared Verifications to Data
	//||------------------------------------------------------------------------------------------------||

	sharedTx := abstract.SharedTransaction{
		AccountId: oa.Enforcement.User.ID,
		SiteId:    oa.Enforcement.Site.ID,
	}
	sharedVerifications := abstract.SharedVerifications{}
	for _, v := range oa.Enforcement.Scopes {
		if v.Code.String() == scope {
			sharedVerifications = append(sharedVerifications, abstract.SharedVerification{
				Type:         v.Code.String(),
				Verification: v.Verification,
			})
		}
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
	//|| Replace Zone Markers
	//||------------------------------------------------------------------------------------------------||

	tpl := template.Create("verified")
	tpl.Add("RETURN_URL", keep.ReturnURL)

	//||------------------------------------------------------------------------------------------------||
	//|| Delete old oauth key and save new access key
	//||------------------------------------------------------------------------------------------------||

	//_ = oa.RemoveAccess()

	//||------------------------------------------------------------------------------------------------||
	//|| Write the response
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write([]byte(tpl.Compile()))

}
