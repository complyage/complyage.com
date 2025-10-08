package verifier

import (
	"encoding/json"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/complyage/complyagent.com/publish"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/validate"
)

//||------------------------------------------------------------------------------------------------||
//|| Card Verification Success Request
//||------------------------------------------------------------------------------------------------||

type faceVerifyRequest struct {
	DOB        types.DOB `json:"dob"`
	Identifier string    `json:"identifier"`
}

type faceVerifyResponse struct {
	Identifier string            `json:"uuid"`
	Status     verify.StatusType `json:"status"`
}

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDSuccessHandler
//||------------------------------------------------------------------------------------------------||

func VerifyFaceSuccessHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Face Success")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest faceVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Date
	//||------------------------------------------------------------------------------------------------||

	if !validate.IsValidDate(updateRequest.DOB.Year, updateRequest.DOB.Month, updateRequest.DOB.Day) {
		responses.Error(w, http.StatusBadRequest, "Invalid date of birth")
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
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(updateRequest.Identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Face
	//||------------------------------------------------------------------------------------------------||

	if verifyRecord.Type != types.DataTypeFACE {
		responses.Error(w, http.StatusBadRequest, "Verification record is not a Face verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the DOB
	//||------------------------------------------------------------------------------------------------||

	verifyData := verifyRecord.Data.FACE
	verifyData.DOB = updateRequest.DOB
	verifyRecord.SetDataFACE(verifyData)
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("QUEUED_L1"), "")
	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	insert := publish.AgentVerifyFaceStart(verifyRecord)
	if insert != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to start Face verification: "+insert.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, faceVerifyResponse{
		Identifier: verifyRecord.UUID,
		Status:     verify.StatusPendingVerification,
	})
}
