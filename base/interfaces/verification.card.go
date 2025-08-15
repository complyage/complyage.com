//--------------------------------------------------------------------------------------------------
// VerificationCard
// Handles the Card Verification Process
//--------------------------------------------------------------------------------------------------

package interfaces

//--------------------------------------------------------------------------------------------------
// CurrencyAmount
//--------------------------------------------------------------------------------------------------

type CurrencyAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

//--------------------------------------------------------------------------------------------------
// VerificationCard - Initial Request
//--------------------------------------------------------------------------------------------------

type VerificationCardInitialRequest struct {
	Amount   float64 `json:"amount"`
	Currency string  `json:"currency"`
}

//--------------------------------------------------------------------------------------------------
// VerificationCard - Initial Response
//--------------------------------------------------------------------------------------------------

type VerificationCardInitialResponse struct {
	Amount       int64  `json:"amount"`
	Currency     string `json:"currency"`
	UUID         string `json:"uuid"`
	ClientSecret string `json:"clientSecret"`
}

//--------------------------------------------------------------------------------------------------
// VerificationCard - Update Request
//--------------------------------------------------------------------------------------------------

type VerificationCardUpdateRequest struct {
	Amount        int64  `json:"amount"`
	Currency      string `json:"currency"`
	UUID          string `json:"uuid"`
	BillingZip    string `json:"billingZip,omitempty"`
	ClientSecret  string `json:"clientSecret"`
	TransactionID string `json:"transactionId"`
}

//--------------------------------------------------------------------------------------------------
// VerificationCard - Update Response
//--------------------------------------------------------------------------------------------------

type VerificationCardUpdateResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status"`
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

//--------------------------------------------------------------------------------------------------
// VerificationCard
//--------------------------------------------------------------------------------------------------

type VerificationCardResponse struct {
	Success bool   `json:"success"`
	Last4   string `json:"last4"`
	TxID    string `json:"txId"`
}
