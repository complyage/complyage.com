package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/complyage/base/adapters"
	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type cardInitRequest struct {
	Base     float64 `json:"baseAmount"`
	Donation float64 `json:"donationAmount"`
	Total    float64 `json:"totalAmount"`
	Currency string  `json:"currency"`
}

type cardInitResponse struct {
	Identifier   string `json:"identifier"`
	Amount       int64  `json:"amount"`
	Currency     string `json:"currency"`
	ClientSecret string `json:"clientSecret"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func CCVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

	verify.LogInfo("CCVerifyInitHandler")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest cardInitRequest
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

	session, err := actions.FetchSession(cookie.Value)
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
	//|| Encrypt
	//||------------------------------------------------------------------------------------------------||

	encrypt, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get encryption keys")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Create(verify.DataTypeCRCD, account.ID, app.Storages["verifications"], app.SQLDB["main"], encrypt.Private, encrypt.Public)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	intent, err := adapters.StripeIntent(verifyRecord.UUID, updateRequest.Total, updateRequest.Currency, verifyRecord.TwoFactor.Code)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create payment intent")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, cardInitResponse{
		Identifier:   verifyRecord.UUID,
		Amount:       intent.Amount,
		Currency:     string(intent.Currency),
		ClientSecret: intent.ClientSecret,
	})

}
