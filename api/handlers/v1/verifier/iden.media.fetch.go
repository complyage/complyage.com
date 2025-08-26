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
//|| Handler: IDVerifyStatusMediaHandler
//|| Endpoint: GET /api/verification/media?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func VerifyIDMediaFetch(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	which := r.URL.Query().Get("which")

	//||------------------------------------------------------------------------------------------------||
	//|| Check Which
	//||------------------------------------------------------------------------------------------------||

	if which != "front" && which != "back" && which != "selfie" {
		responses.Error(w, http.StatusBadRequest, "Invalid 'which' value")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Account
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session Match
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.IDAccount {
		responses.Error(w, http.StatusForbidden, "Session does not match verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Verification Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Data
	//||------------------------------------------------------------------------------------------------||

	verifyData := verifyRecord.Encrypted.Data.IDEN

	//||------------------------------------------------------------------------------------------------||
	//|| Add the Correct Media
	//||------------------------------------------------------------------------------------------------||

	mediaRecord := verify.Media{}

	if which == "front" {
		mediaRecord = verifyData.Front
	} else if which == "back" {
		mediaRecord = verifyData.Back
	} else if which == "selfie" {
		mediaRecord = verifyData.Selfie
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, mediaRecord)
}
