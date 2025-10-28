package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: IDVerifyStatusMediaHandler
//|| Endpoint: GET /api/verification/media?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func VerifyFaceMediaFetch(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Face Media Fetch")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add the Correct Media
	//||------------------------------------------------------------------------------------------------||

	mediaRecord := verifyRecord.Data.FACE.Selfie

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, mediaRecord)
}
