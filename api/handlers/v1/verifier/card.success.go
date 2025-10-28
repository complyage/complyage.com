package verifier

import (
	"api/handlers/utils"
	"encoding/json"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

//||------------------------------------------------------------------------------------------------||
//|| Card Verification Success Request
//||------------------------------------------------------------------------------------------------||

type cardVerifyRequest struct {
	Identifier    string `json:"identifier"`
	Base          int64  `json:"base"`
	Donation      int64  `json:"donation"`
	Currency      string `json:"currency"`
	Amount        int64  `json:"amount"`
	BillingZip    string `json:"billingZip,omitempty"`
	TransactionID string `json:"transactionId"`
	ClientSecret  string `json:"clientSecret"`
}

type cardVerifyResponse struct {
	Identifier string `json:"uuid"`
	Status     string `json:"status"`
}

//||------------------------------------------------------------------------------------------------||
//|| CCVerifyUpdateHandler
//||------------------------------------------------------------------------------------------------||

func CCVerifySuccessHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: CC Success")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest cardVerifyRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Stripe Payment Intent (expand payment_method & latest_charge)
	//||------------------------------------------------------------------------------------------------||

	params := &stripe.PaymentIntentParams{}
	params.AddExpand("payment_method")
	params.AddExpand("latest_charge")
	params.AddExpand("latest_charge.payment_method_details")
	intent, err := paymentintent.Get(updateRequest.TransactionID, params)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Failed to retrieve payment intent: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Extract Card Details (brand/last4)
	//||------------------------------------------------------------------------------------------------||

	lastFour := ""
	cardType := ""
	if intent.PaymentMethod != nil && intent.PaymentMethod.Card != nil {
		lastFour = intent.PaymentMethod.Card.Last4
		cardType = string(intent.PaymentMethod.Card.Brand)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fallback if we don't have lastFour/cardType
	//||------------------------------------------------------------------------------------------------||

	if (lastFour == "" || cardType == "") && intent.LatestCharge != nil &&
		intent.LatestCharge.PaymentMethodDetails != nil &&
		intent.LatestCharge.PaymentMethodDetails.Card != nil {
		if lastFour == "" {
			lastFour = intent.LatestCharge.PaymentMethodDetails.Card.Last4
		}
		if cardType == "" {
			cardType = string(intent.LatestCharge.PaymentMethodDetails.Card.Brand)
		}
	}
	app.Log.Info("Card Type:", cardType, "Last4:", lastFour)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the USD
	//||------------------------------------------------------------------------------------------------||

	usdAmount, cErr := utils.ConvertToUSD(updateRequest.Amount, updateRequest.Currency)
	if cErr != nil {
		usdAmount = -1
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save Transaction
	//||------------------------------------------------------------------------------------------------||

	if err := abstract.AddTransaction(
		"CARD",                        // method
		"STRIPE",                      // merchant
		(usdAmount / 100),             // USD value
		float64(updateRequest.Amount), // original amount in major units
		updateRequest.Currency,        // original currency
		updateRequest.TransactionID,   // transaction reference
	); err != nil {
		app.Log.Error("Failed to insert transaction: ", err.Error())
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
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(updateRequest.Identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Approve the Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.TransactionSaved = true
	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, cardVerifyResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.String(),
	})
}
