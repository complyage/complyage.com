package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/send"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/locale"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
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

	app.Log.Info("Handler: Phone Resend")

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
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Verification Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(req.Identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
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

	phone := verifyRecord.Data.PHNE
	code := verifyRecord.TwoFactor.Code
	bodyTxt, sendErr := send.SendVerifyText(phone.CountryCode+phone.Number, code, locale.Request(r))
	if sendErr != nil {
		app.Log.Info("Error sending verification SMS:", sendErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to send verification SMS")
		return
	}
	app.Log.Data("Verification SMS sent successfully:", bodyTxt)

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("SENT_SMS"), verifyRecord.Status.Description())
	verifyRecord.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, phoneVerifyResendRequest{
		Identifier: verifyRecord.UUID,
	})

}
