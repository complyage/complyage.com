package verifier

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"base/verify"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| VerificationCardCodeRequest
//||------------------------------------------------------------------------------------------------||

type codeCheckRequest struct {
	Identifier string `json:"uuid"`
	Code       string `json:"code"`
}

type codeCheckResponse struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Details    string `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Card Code Verification Attempt
//||------------------------------------------------------------------------------------------------||

func VerificationCodeCheck(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	var updateRequest codeCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	if updateRequest.Identifier == "" || updateRequest.Code == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier or code")
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

	account, err := abstract.GetAccountByVerificationUUID(updateRequest.Identifier)
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

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], updateRequest.Identifier, account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("VerificationCodeLoad: Loaded verification", verifyRecord.UUID, "Code:", verifyRecord.TwoFactor.Code)

	//||------------------------------------------------------------------------------------------------||
	//|| Check Code
	//||------------------------------------------------------------------------------------------------||

	err = verifyRecord.TwoFactorVerify(updateRequest.Code)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	fmt.Println(verifyRecord.Type, "VerificationCodeCheck: Code verified for", verifyRecord.UUID, "Status now", verifyRecord.Status)
	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, codeCheckResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.String(),
		Type:       verifyRecord.Type.String(),
		Details:    "Verification code is valid",
	})
}
