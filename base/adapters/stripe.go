package adapters

import (
	"log"
	"os"
	"strings"

	"github.com/stripe/stripe-go/v82"
	"github.com/stripe/stripe-go/v82/paymentintent"
)

//||------------------------------------------------------------------------------------------------||
//|| Zero Decimal Currencies
//||------------------------------------------------------------------------------------------------||

var zeroDecimalCurrencies = map[string]bool{
	"BIF": true, "CLP": true, "DJF": true, "GNF": true,
	"JPY": true, "KMF": true, "KRW": true, "MGA": true,
	"PYG": true, "RWF": true, "UGX": true, "VND": true,
	"VUV": true, "XAF": true, "XOF": true, "XPF": true,
}

//||------------------------------------------------------------------------------------------------||
//|| Initalize and Set Stripe Key
//||------------------------------------------------------------------------------------------------||

func InitStripe() {
	key := os.Getenv("MERCHANT_STRIPE_PRIVATE") // e.g. sk_live_... or sk_test_...
	if key == "" {
		log.Fatal("MERCHANT_STRIPE_PRIVATE is not set")
	}
	stripe.Key = key
}

//||------------------------------------------------------------------------------------------------||
//|| stripeAmount converts a float amount into the correct Stripe int value
//||------------------------------------------------------------------------------------------------||

func StripeAmount(amount float64, currency string) int64 {
	currency = strings.ToUpper(currency)
	if zeroDecimalCurrencies[currency] {
		return int64(amount)
	}
	return int64(amount)
}

//||------------------------------------------------------------------------------------------------||
//|| stripeIntent creates a PaymentIntent with server-controlled parameters
//||------------------------------------------------------------------------------------------------||

func StripeIntent(UUID string, rawAmount float64, currency string, suffix string) (*stripe.PaymentIntent, error) {
	amt := StripeAmount(rawAmount, currency)

	params := &stripe.PaymentIntentParams{
		Amount:             stripe.Int64(amt),
		Currency:           stripe.String(currency),
		PaymentMethodTypes: stripe.StringSlice([]string{"card"}),

		StatementDescriptor:       stripe.String(os.Getenv("VITE_MERCHANT_DESCRIPTOR")),
		StatementDescriptorSuffix: stripe.String("CODE-" + suffix),

		Metadata: map[string]string{
			"uuid": UUID,
		},
	}

	return paymentintent.New(params)
}
