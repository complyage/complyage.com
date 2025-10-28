package verifier

import (
	"encoding/json"
	"net/http"

	"github.com/complyage/complyagent.com/publish"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Card Verification Success Request
//||------------------------------------------------------------------------------------------------||

type idenVerifyRequest struct {
	Identifier string `json:"identifier"`
}

type idenVerifyResponse struct {
	Identifier string            `json:"uuid"`
	Status     verify.StatusType `json:"status"`
}

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDSuccessHandler
//||------------------------------------------------------------------------------------------------||

func VerifyIDSuccessHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: ID Success")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest idenVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

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

	verifyRecord, err := verify.CheckLoad(updateRequest.Identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	insert := publish.AgentVerifyIDStart(verifyRecord)
	if insert != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("QUEUED_L1"), "")
	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, idenVerifyResponse{
		Identifier: verifyRecord.UUID,
		Status:     verify.STATUSES.PendingVerification,
	})
}
