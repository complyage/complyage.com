package models

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Transaction Model (maps to `transactions` table)
//||------------------------------------------------------------------------------------------------||

type Transactions struct {
	ID        int       `gorm:"column:id_transaction;primaryKey;autoIncrement" json:"id_transaction"`
	Method    string    `gorm:"column:transaction_method;size:6"               json:"transaction_method,omitempty"`
	Merchant  string    `gorm:"column:transaction_merchant;size:16"            json:"transaction_merchant,omitempty"`
	Amount    float64   `gorm:"column:transaction_amount"                      json:"transaction_amount"`
	TxID      string    `gorm:"column:transaction_id;size:64"                  json:"transaction_id,omitempty"`
	Timestamp time.Time `gorm:"column:transaction_timestamp"                   json:"transaction_timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName override
//||------------------------------------------------------------------------------------------------||

func (Transactions) TableName() string {
	return "transactions"
}
