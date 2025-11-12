package pages

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"oauth/access"

	"github.com/complyage/base/encrypted"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Approve OAuth Handler
//||------------------------------------------------------------------------------------------------||

func PreviewShareHandler(w http.ResponseWriter, r *http.Request) {

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
		responses.Error(w, http.StatusBadRequest, app.Err("OAuth").Code("MISSING_SESSION"))
		return
	}
	if key == "" {
		responses.Error(w, http.StatusBadRequest, app.Err("OAuth").Code("MISSING_CHECKKEY"))
		return
	}
	if scope == "" {
		responses.Error(w, http.StatusBadRequest, app.Err("OAuth").Code("MISSING_SCOPE"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Redis Session
	//||------------------------------------------------------------------------------------------------||

	oa, err := access.LoadAccess(sessionId)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Key
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Preview Key:", oa.PreviewKey)
	app.Log.Data("Provided Key:", key)
	if oa.PreviewKey == "" || key != oa.PreviewKey {
		responses.Error(w, http.StatusBadRequest, app.Err("OAuth").Code("INVALID_CHECKKEY"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Verification
	//||------------------------------------------------------------------------------------------------||

	result, err := encrypted.LoadPreview(verification, oa.Enforcement.User.Private, scope)
	if err != nil {
		app.Log.Error(err.Error())
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	jsonBytes, _ := json.MarshalIndent(result, "", "  ")
	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonBytes)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Verification
	//||------------------------------------------------------------------------------------------------||

}
