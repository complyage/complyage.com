//--------------------------------------------------------------------------------------------------
// VerificationCard
// Handles the Card Verification Process
//--------------------------------------------------------------------------------------------------

package verification

import (
	"base/interfaces"
	"time"
)

//--------------------------------------------------------------------------------------------------
// CurrencyAmount
//--------------------------------------------------------------------------------------------------

type CurrencyAmount struct {
	Amount   string `json:"amount"`
	Currency string `json:"currency"`
}

//--------------------------------------------------------------------------------------------------
// VerificationRecord
//--------------------------------------------------------------------------------------------------

type VerificationRecord struct {
	ID         int64                         `json:"id"`
	UUID       string                        `json:"uuid"`
	FidAccount int64                         `json:"fidAccount"`
	Type       string                        `json:"type"`
	Display    string                        `json:"display"`
	Data       []byte                        `json:"data"`
	Meta       interfaces.VerificationMeta   `json:"meta"`
	Secret     interfaces.VerificationSecret `json:"secret"`
	Status     string                        `json:"status"`
	CreatedAt  time.Time                     `json:"created"`
	UpdatedAt  time.Time                     `json:"updated"`
}

//--------------------------------------------------------------------------------------------------
// Address Verification Request
//--------------------------------------------------------------------------------------------------

type VerificationIDStatusProcess struct {
	Status string                            `json:"status"`
	Type   string                            `json:"type,omitempty"`
	UUID   string                            `json:"verificationUUID"`
	Step   int                               `json:"step"`
	Steps  []interfaces.VerificationMetaStep `json:"steps,omitempty"`
}

//--------------------------------------------------------------------------------------------------
// Address Verification Initial Request
//--------------------------------------------------------------------------------------------------

type VerificationIDProgressRequest struct {
	UUID        string                            `json:"uuid"`
	Step        int                               `json:"step"`
	StepName    string                            `json:"stepName"`
	StepStatus  string                            `json:"stepStatus"`
	StepDetails string                            `json:"stepDetails,omitempty"`
	Steps       []interfaces.VerificationMetaStep `json:"steps,omitempty"`
}

//--------------------------------------------------------------------------------------------------
// Address Verification Initial Request
//--------------------------------------------------------------------------------------------------

type VerificationAddressInitialRequest struct {
	Address  interfaces.Address `json:"address"`
	Amount   float64            `json:"amount"`
	Currency string             `json:"currency"`
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
