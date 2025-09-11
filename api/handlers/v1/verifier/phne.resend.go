package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/send"
	"base/verify"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/ralphferrara/aria/locale"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type phoneVerifyResendRequest struct {
	Identifier string `json:"identifier"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func PhoneVerifyResendHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse JSON
	//||------------------------------------------------------------------------------------------------||

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate verification tuple (amount + 4-digit code)
	//||------------------------------------------------------------------------------------------------||

	var req phoneVerifyResendRequest
	if err := json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload"+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], req.Identifier, session.Private, session.Public)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Count
	//||------------------------------------------------------------------------------------------------||

	sendCount := verifyRecord.CountStepsOfType(app.Constants("VERIFY_STEP_TYPES").Get("SENT_SMS"))
	if sendCount > 3 {
		responses.Error(w, http.StatusBadRequest, "Maximum number of resend attempts reached")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Send the verification SMS
	//||------------------------------------------------------------------------------------------------||

	phone := verifyRecord.Encrypted.Data.PHNE
	code := verifyRecord.TwoFactor.Code
	bodyTxt, sendErr := send.SendVerifyText(phone.CountryCode+phone.Number, code, locale.Request(r))
	if sendErr != nil {
		fmt.Println("Error sending verification SMS:", sendErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to send verification SMS")
		return
	}
	fmt.Println("Verification SMS sent successfully:", bodyTxt)

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("SENT_SMS"), verifyRecord.Status.Description())

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, phoneVerifyResendRequest{
		Identifier: verifyRecord.UUID,
	})

}
