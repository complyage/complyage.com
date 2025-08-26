package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/adapters"
	"base/helpers"
	"base/responses"
	"base/verify"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type addressInitRequest struct {
	Base     float64        `json:"baseAmount"`
	Donation float64        `json:"donationAmount"`
	Total    float64        `json:"totalAmount"`
	Currency string         `json:"currency"`
	Address  verify.Address `json:"address"`
}

type addressInitResponse struct {
	Identifier   string `json:"identifier"`
	Amount       int64  `json:"amount"`
	Currency     string `json:"currency"`
	ClientSecret string `json:"clientSecret"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func AddressVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

	verify.LogInfo("CCVerifyInitHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest addressInitRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
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

	verifyRecord, err := verify.Init(verify.DataTypeADDR, account.IDAccount, app.Storages["verifications"], app.SQLDB["main"], account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Address
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.SetDataADDR(updateRequest.Address)
	verifyRecord.Save()
	verifyRecord.DatabaseUpdate()

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	intent, err := adapters.StripeIntent(verifyRecord.UUID, updateRequest.Total, updateRequest.Currency, "") //No code for address verifications
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create payment intent")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, addressInitResponse{
		Identifier:   verifyRecord.UUID,
		Amount:       intent.Amount,
		Currency:     string(intent.Currency),
		ClientSecret: intent.ClientSecret,
	})

}
