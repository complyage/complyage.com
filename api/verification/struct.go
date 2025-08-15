//--------------------------------------------------------------------------------------------------
// VerificationCard
// Handles the Card Verification Process
//--------------------------------------------------------------------------------------------------

package verification

//--------------------------------------------------------------------------------------------------
// CurrencyAmount
//--------------------------------------------------------------------------------------------------

type CurrencyAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

//--------------------------------------------------------------------------------------------------
// VerificationCard
//--------------------------------------------------------------------------------------------------

type VerificationCardRequest struct {
	BaseAmount       CurrencyAmount `json:"baseAmount"`
	ChargeAmount     CurrencyAmount `json:"chargeAmount"`
	Donation         float64        `json:"donation"`
	Currency         string         `json:"currency"`
	CardNumber       *string        `json:"cardNumber,omitempty"`
	ExpMonth         *string        `json:"expMonth,omitempty"`
	ExpYear          *string        `json:"expYear,omitempty"`
	CVC              *string        `json:"cvc,omitempty"`
	BillingZip       *string        `json:"billingZip,omitempty"`
	LastFour         *string        `json:"lastFour,omitempty"`
	CardType         *string        `json:"cardType,omitempty"`
	TransactionID    *string        `json:"transactionId,omitempty"`
	VerificationUUID *string        `json:"verificationUUID,omitempty"`
}
