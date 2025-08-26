package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"base/verify"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type idenInitResponse struct {
	Identifier string `json:"identifier"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func IdentifierVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

	verify.LogInfo("IdentifierVerifyInitHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Account
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil || account == nil {
		responses.Error(w, http.StatusBadRequest, "Account not found for session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Init(verify.DataTypeIDEN, account.IDAccount, app.Storages["verifications"], app.SQLDB["main"], account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, idenInitResponse{
		Identifier: verifyRecord.UUID,
	})

}
