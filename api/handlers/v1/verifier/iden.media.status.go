package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/verification"
	"base/abstract"
	"base/helpers"
	"base/responses"
	"fmt"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: VerifyIDStatusHandler
//|| Endpoint: GET /api/verification/status?identifier=...
//||------------------------------------------------------------------------------------------------||

func VerifyIDStatusHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	if identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier param")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Find Verification Record
	//||------------------------------------------------------------------------------------------------||

	verificationRecord, err := verification.LoadVerificationRecordByUUID(identifier)
	if err != nil || verificationRecord == nil {
		fmt.Println("Error loading verification record:", err)
		responses.Error(w, http.StatusNotFound, "Verification record not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Account for Verification Record
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Please login to continue")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session Account Match
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.IDAccount {
		responses.Error(w, http.StatusForbidden, "Session does not match verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Minimal Status Object
	//||------------------------------------------------------------------------------------------------||

	resp := verification.VerificationIDStatusProcess{
		UUID:   identifier,
		Status: verificationRecord.Status,
		Type:   verificationRecord.Type,
		Step:   verificationRecord.Meta.Step,
		Steps:  verificationRecord.Meta.Steps,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, resp)
}
