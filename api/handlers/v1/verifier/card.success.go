package verifier

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"base/verify"
	"encoding/json"
	"net/http"

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

	verify.LogInfo("CCVerifySuccessHandler")

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

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	if updateRequest.Identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier")
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
	//|| Fetch a Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(updateRequest.Identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusBadRequest, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.IDAccount {
		responses.Error(w, http.StatusBadRequest, "Session does not match account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], updateRequest.Identifier, account.AccountPrivate, account.AccountPublic)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Log the Transaction
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.TransactionApproved(
		verify.TransactionTypeCredit,
		updateRequest.Base,
		updateRequest.Donation,
		updateRequest.Currency,
		cardType,
		lastFour,
		updateRequest.ClientSecret,
		verify.Address{
			Postal: updateRequest.BillingZip,
		},
		verify.Address{})

	//||------------------------------------------------------------------------------------------------||
	//|| Approve the Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, cardVerifyResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.String(),
	})
}
