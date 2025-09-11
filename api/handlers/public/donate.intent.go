package public

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"

	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type PaymentIntentRequest struct {
	Amount   int64  `json:"amount"`
	Currency string `json:"currency"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func CreatePaymentIntentHandler(w http.ResponseWriter, r *http.Request) {
	stripe.Key = os.Getenv("MERCHANT_STRIPE_PRIVATE")
	if stripe.Key == "" {
		responses.Error(w, http.StatusInternalServerError, "Stripe key missing")
		return
	}

	var req PaymentIntentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	pi, err := paymentintent.New(&stripe.PaymentIntentParams{
		Amount:   stripe.Int64(req.Amount),
		Currency: stripe.String(strings.ToLower(strings.TrimSpace(req.Currency))),
		AutomaticPaymentMethods: &stripe.PaymentIntentAutomaticPaymentMethodsParams{
			Enabled: stripe.Bool(true),
		},
	})
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Stripe error: "+err.Error())
		return
	}

	responses.Success(w, http.StatusOK, map[string]any{
		"clientSecret": pi.ClientSecret,
		"id":           pi.ID,
	})
}
