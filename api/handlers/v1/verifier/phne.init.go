package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/send"
	"github.com/complyage/base/types"

	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/locale"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
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

	app.Log.Info("Handler: Phone Init")

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
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord := verify.Create(types.DataTypePHNE, account)
	verifyRecord.Save()
	verifyRecord.DatabaseInsert()

	//||------------------------------------------------------------------------------------------------||
	//|| We are in progress
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("STATUS_CHANGE"), verifyRecord.Status.Description())
	verifyRecord.SetDataPhone(types.PhoneNumber{CountryCode: req.CountryCode, Number: req.Phone})

	//||------------------------------------------------------------------------------------------------||
	//|| Send the verification SMS
	//||------------------------------------------------------------------------------------------------||

	bodyTxt, sendErr := send.SendVerifyText(req.CountryCode+req.Phone, verifyRecord.TwoFactor.Code, locale.Request(r))
	if sendErr != nil {
		fmt.Println("Error sending verification SMS:", sendErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to send verification SMS")
		return
	}
	app.Log.Data("Verification SMS sent successfully:", bodyTxt, verifyRecord.TwoFactor.Code)

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("SENT_SMS"), verifyRecord.Status.Description())
	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, handlerResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.Code(),
	})

}
