package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"base/verify"
	"net/http"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

type VerifyStatusResponse struct {
	UUID   string        `json:"uuid"`
	Status string        `json:"status"`
	Type   string        `json:"type"`
	Step   int           `json:"step"`
	Steps  []verify.Step `json:"steps"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: VerifyIDStatusHandler
//||------------------------------------------------------------------------------------------------||

func VerifyStatusHandler(w http.ResponseWriter, r *http.Request) {

	verify.LogInfo("VerifyStatusHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	if identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier param")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	if identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier")
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

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusBadRequest, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.IDAccount {
		responses.Error(w, http.StatusBadRequest, "Session does not match account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Minimal Status Object
	//||------------------------------------------------------------------------------------------------||

	resp := VerifyStatusResponse{
		UUID:   verifyRecord.UUID,
		Status: verifyRecord.Status.String(),
		Type:   verifyRecord.Type.String(),
		Step:   verifyRecord.Step,
		Steps:  verifyRecord.Steps,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, resp)
}
