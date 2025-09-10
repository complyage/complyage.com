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
//|| Request / Response
//||------------------------------------------------------------------------------------------------||

type handlerRequest struct {
	CountryCode string `json:"countryCode"`
	Phone       string `json:"phone"`
}

type handlerResponse struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func PhoneVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

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

	var req handlerRequest
	if err := json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload"+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Create(verify.DataTypePHNE, session.ID, app.Storages["verifications"], app.SQLDB["main"], session.Private, session.Public)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| We are in progress
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("STATUS_CHANGE"), verifyRecord.Status.Description())

	//||------------------------------------------------------------------------------------------------||
	//|| Send the Data
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.SetDataPhone(verify.PhoneNumber{CountryCode: req.CountryCode, Number: req.Phone})

	//||------------------------------------------------------------------------------------------------||
	//|| Send the verification SMS
	//||------------------------------------------------------------------------------------------------||

	bodyTxt, sendErr := send.SendVerifyText(req.CountryCode+req.Phone, verifyRecord.TwoFactor.Code, locale.Request(r))
	if sendErr != nil {
		fmt.Println("Error sending verification SMS:", sendErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to send verification SMS")
		return
	}
	fmt.Println("Verification SMS sent successfully:", bodyTxt, verifyRecord.TwoFactor.Code)
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("SENT_SMS"), verifyRecord.Status.Description())

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, handlerResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.Code(),
	})

}
