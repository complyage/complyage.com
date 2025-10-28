package verify

import (
	"time"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| Transaction
//||------------------------------------------------------------------------------------------------||

type TransactionPrivate struct {
	ID           string            `json:"id"`
	Type         TransactionType   `json:"type"`
	Status       TransactionStatus `json:"status"`
	Base         int64             `json:"baseAmount"`
	Donation     int64             `json:"domationAmount,omitempty"`
	Total        int64             `json:"totalAmount"`
	Currency     string            `json:"currency"`
	Timestamp    time.Time         `json:"timestamp,omitempty"`
	CardType     string            `json:"cardType,omitempty"`
	LastFour     string            `json:"lastFour,omitempty"`
	ClientSecret string            `json:"clientSecret,omitempty"`
	Billing      types.Address     `json:"billing,omitempty"`
	Shipping     types.Address     `json:"shipping,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Transaction Public
//||------------------------------------------------------------------------------------------------||

type TransactionPublic struct {
	Display   string            `json:"display"`
	Type      TransactionType   `json:"type"`
	Status    TransactionStatus `json:"status"`
	Amount    int64             `json:"amount"`
	Currency  string            `json:"currency"`
	Timestamp time.Time         `json:"timestamp,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Transaction Public
//||------------------------------------------------------------------------------------------------||

func (v *Verification) TransactionPublicInit() {
	//||------------------------------------------------------------------------------------------------||
	//|| Only add for Credit Card and Address
	//||------------------------------------------------------------------------------------------------||
	if v.Type != DataTypeCRCD && v.Type != DataTypeADDR {
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Add Public Transaction
	//||------------------------------------------------------------------------------------------------||
	v.Transaction = TransactionPublic{
		Status:   TransactionPending,
		Currency: "USD",
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Transaction Private
//||------------------------------------------------------------------------------------------------||

func (e *Encrypted) TransactionPrivateInit() {
	//||------------------------------------------------------------------------------------------------||
	//|| Only add for Credit Card and Address
	//||------------------------------------------------------------------------------------------------||
	if e.Type != DataTypeCRCD && e.Type != DataTypeADDR {
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| AddPrivate Transaction
	//||------------------------------------------------------------------------------------------------||
	e.Transaction = TransactionPrivate{
		Status:   TransactionPending,
		Currency: "USD",
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Transaction Private
//||------------------------------------------------------------------------------------------------||

func (v *Verification) TransactionApproved(transactionType TransactionType, base int64, donation int64, currency string, cardType string, lastFour string, clientSecret string, billing types.Address, shipping types.Address) {
	LogInfo("Verify: Transaction Approval")
	//||------------------------------------------------------------------------------------------------||
	//|| Only add for Credit Card and Address
	//||------------------------------------------------------------------------------------------------||
	if v.Type != DataTypeCRCD && v.Type != DataTypeADDR {
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Total Amount
	//||------------------------------------------------------------------------------------------------||
	total := base + donation
	//||------------------------------------------------------------------------------------------------||
	//|| Add Public Transaction
	//||------------------------------------------------------------------------------------------------||
	v.Transaction = TransactionPublic{
		Display:  "CC-" + types.CreditCard{LastFour: lastFour, CardType: cardType}.Mask(),
		Type:     transactionType,
		Status:   TransactionApproved,
		Amount:   total,
		Currency: currency,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Add Private Transaction
	//||------------------------------------------------------------------------------------------------||
	v.Encrypted.Transaction = TransactionPrivate{
		ID:           v.UUID,
		Type:         transactionType,
		Status:       TransactionApproved,
		Base:         base,
		Donation:     donation,
		Total:        total,
		Currency:     currency,
		Timestamp:    time.Now().UTC(),
		CardType:     cardType,
		LastFour:     lastFour,
		ClientSecret: clientSecret,
		Billing:      billing,
		Shipping:     shipping,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
}
