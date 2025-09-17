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
	"github.com/ralphferrara/aria/auth/actions"
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

	session, err := actions.FetchSession(cookie.Value)
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

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, encrypt.Private, encrypt.Public)
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
