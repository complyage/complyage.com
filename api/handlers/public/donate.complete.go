package public

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"base/db/abstract"

	"github.com/ralphferrara/aria/responses"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

//||------------------------------------------------------------------------------------------------||
//|| Request Payload
//||------------------------------------------------------------------------------------------------||

type DonateCompleteRequest struct {
	Method   string  `json:"method"`   // "CARD", "CRYPTO", "CHECK"
	Merchant string  `json:"merchant"` // "Stripe", "PayPal", "Coinbase"
	Amount   float64 `json:"amount"`   // donation amount in dollars
	TxID     string  `json:"txID"`     // reference/intent ID
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: DonateComplete
//||------------------------------------------------------------------------------------------------||

func DonateComplete(w http.ResponseWriter, r *http.Request) {
	var req DonateCompleteRequest

	// Parse incoming JSON
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	// For CARD payments, verify with Stripe
	if strings.ToUpper(req.Method) == "CARD" && strings.ToLower(req.Merchant) == "stripe" {
		stripe.Key = os.Getenv("MERCHANT_STRIPE_PRIVATE")
		if stripe.Key == "" {
			responses.Error(w, http.StatusInternalServerError, "Stripe secret key missing")
			return
		}

		pi, err := paymentintent.Get(req.TxID, nil)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, "Stripe lookup failed: "+err.Error())
			return
		}

		// Ensure status is succeeded
		if pi.Status != stripe.PaymentIntentStatusSucceeded {
			responses.Error(w, http.StatusBadRequest, "Payment not completed")
			return
		}

		// Validate amount matches (Stripe uses cents)
		if float64(pi.Amount)/100 != req.Amount {
			responses.Error(w, http.StatusBadRequest, "Amount mismatch")
			return
		}
	}

	// Save to DB
	if err := abstract.AddTransaction(req.Method, req.Merchant, req.Amount, req.TxID); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to record donation: "+err.Error())
		return
	}

	// Success response
	responses.Success(w, http.StatusOK, map[string]any{
		"message": "Donation recorded successfully",
		"txID":    req.TxID,
	})
}
