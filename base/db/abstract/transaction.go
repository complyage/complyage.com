package abstract

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"time"

	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| AddTransaction - Inserts a new transaction into the DB
//||------------------------------------------------------------------------------------------------||
//|| Params:
//||    method   : string   (e.g. "CARD", "CRYPTO")
//||    merchant : string   (e.g. "Stripe", "PayPal")
//||    amount   : float64  (converted to USD)
//||    origAmt  : float64  (original currency amount)
//||    origCurr : string   (e.g. "EUR", "JPY")
//||    txID     : string   (transaction reference)
//|| Returns: error (nil if success)
//||------------------------------------------------------------------------------------------------||

func AddTransaction(method, merchant string, amount float64, origAmt float64, origCurr string, txID string) error {
	trx := models.Transactions{
		Method:           method,
		Merchant:         merchant,
		Amount:           amount,
		OriginalAmount:   origAmt,
		OriginalCurrency: origCurr,
		TxID:             txID,
		Timestamp:        time.Now(),
	}

	if err := app.SQLDB["main"].DB.Create(&trx).Error; err != nil {
		return err
	}
	return nil
}
