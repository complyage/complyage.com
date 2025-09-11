package abstract

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"time"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| AddTransaction - Inserts a new transaction into the DB
//||------------------------------------------------------------------------------------------------||
//|| Params:
//||    method   : string  (e.g. "CARD", "CRYPTO")
//||    merchant : string  (e.g. "Stripe", "PayPal")
//||    amount   : float64
//||    txID     : string  (transaction reference)
//|| Returns: error (nil if success)
//||------------------------------------------------------------------------------------------------||

func AddTransaction(method, merchant string, amount float64, txID string) error {
	trx := models.Transactions{
		Method:    method,
		Merchant:  merchant,
		Amount:    amount,
		TxID:      txID,
		Timestamp: time.Now(),
	}

	if err := app.SQLDB["main"].DB.Create(&trx).Error; err != nil {
		return err
	}
	return nil
}
