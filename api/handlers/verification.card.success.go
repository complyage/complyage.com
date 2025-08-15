//||------------------------------------------------------------------------------------------------||
//|| Handler: Update Card Verification
//|| POST /v1/api/verify/card/success
//||------------------------------------------------------------------------------------------------||

package handlers

import (
	"api/verification"
	"base/constants"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"net/http"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

//||------------------------------------------------------------------------------------------------||
//|| CCVerifyUpdateHandler
//||------------------------------------------------------------------------------------------------||

func CCVerifySuccessHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest interfaces.VerificationCardUpdateRequest
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
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	updErr := verification.UpdateStatusCRCDPendingVerification(updateRequest.UUID, updateRequest.TransactionID, cardType, lastFour, updateRequest.BillingZip)
	if updErr != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to update verification: "+updErr.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   updateRequest.UUID,
		"status": constants.VerificationStatuses.PendingVerification,
	})
}
