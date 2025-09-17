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
	"github.com/ralphferrara/aria/auth/actions"
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

	verify.LogInfo("VerifyFaceSuccessHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest faceVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	if updateRequest.Identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier")
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

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch a Session
	//||------------------------------------------------------------------------------------------------||

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(updateRequest.Identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusBadRequest, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.ID {
		responses.Error(w, http.StatusBadRequest, "Session does not match account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt
	//||------------------------------------------------------------------------------------------------||

	encrypt, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get encryption keys")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], updateRequest.Identifier, encrypt.Private, encrypt.Public)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Face
	//||------------------------------------------------------------------------------------------------||

	if verifyRecord.Type != verify.DataTypeFACE {
		responses.Error(w, http.StatusBadRequest, "Verification record is not a Face verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| DOB
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Encrypted.Data.FACE.DOB = updateRequest.DOB

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Status to Pending Verification
	//||------------------------------------------------------------------------------------------------||

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
	//|| Save
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("QUEUED_L1"), "")
	verifyRecord.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, faceVerifyResponse{
		Identifier: verifyRecord.UUID,
		Status:     verify.StatusPendingVerification,
	})
}
