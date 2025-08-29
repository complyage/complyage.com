package verifier

import (
	"base/db/abstract"
	"base/verify"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Request Response
//||------------------------------------------------------------------------------------------------||

type codeLoadResponse struct {
	Identifier string            `json:"identifier"`
	Status     verify.StatusType `json:"status"`
	Type       verify.DataType   `json:"type"`
	Details    string            `json:"details"`
	Message    string            `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Card Code Verification Attempt
//||------------------------------------------------------------------------------------------------||

func VerificationCodeLoad(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

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
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.ID {
		responses.Error(w, http.StatusUnauthorized, "Session does not match account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, account.Private, account.Public)
	if err != nil {
		responses.Error(w, http.StatusNotFound, "Verification record not found."+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("VerificationCodeLoad: Loaded verification", verifyRecord.UUID, "Code:", verifyRecord.TwoFactor.Code)

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, codeLoadResponse{
		Identifier: identifier,
		Status:     verifyRecord.Status,
		Type:       verifyRecord.Type,
		Details:    "Verification code is valid",
	})
}
