package interfaces

//--------------------------------------------------------------------------------------------------
// VerificationCard - Update Request
//--------------------------------------------------------------------------------------------------

type VerificationAddressUpdateRequest struct {
	Amount        int64   `json:"amount"`
	Currency      string  `json:"currency"`
	UUID          string  `json:"uuid"`
	Address       Address `json:"address,omitempty"`
	ClientSecret  string  `json:"clientSecret"`
	TransactionID string  `json:"transactionId"`
	BillingZip    string  `json:"billingZip,omitempty"`
}
