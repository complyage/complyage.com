package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"

	"github.com/complyage/base/adapters"
	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type addressInitRequest struct {
	Base     float64       `json:"baseAmount"`
	Donation float64       `json:"donationAmount"`
	Total    float64       `json:"totalAmount"`
	Currency string        `json:"currency"`
	Address  types.Address `json:"address"`
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

	app.Log.Info("Handler: Address Init")

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

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord := verify.Create(types.DataTypeADDR, account)
	verifyRecord.SetDataADDR(updateRequest.Address)
	verifyRecord.Save()
	verifyRecord.DatabaseInsert()

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
