package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/complyage/complyagent.com/publish"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDProgressHandler
//||------------------------------------------------------------------------------------------------||

func TestingResetVerification(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Util Reset")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
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
	//|| Load Verification Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Steps
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Steps = []verify.Step{}
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("AGENT_L1"), "")
	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	if verifyRecord.Type == types.DataTypeIDEN {
		app.Log.Data("Starting ID Verification Reset for: " + verifyRecord.UUID)
		insert := publish.AgentVerifyIDStart(verifyRecord)
		if insert != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
			return
		}
	} else {
		app.Log.Data("Starting Facial Verification Reset for: " + verifyRecord.UUID)
		insert := publish.AgentVerifyFaceStart(verifyRecord)
		if insert != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
			return
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   verifyRecord.UUID,
		"status": verifyRecord.Status,
	})
}
