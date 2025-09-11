package models

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Bill Model (maps to `bills` table)
//||------------------------------------------------------------------------------------------------||

type Bills struct {
	ID          int       `gorm:"column:id_bill;primaryKey;autoIncrement" json:"id_bill"`
	Transaction string    `gorm:"column:bill_transaction;size:64"         json:"bill_transaction,omitempty"`
	Type        string    `gorm:"column:bill_type;size:16"                json:"bill_type,omitempty"`
	Vendor      string    `gorm:"column:bill_vendor;size:16"              json:"bill_vendor,omitempty"`
	Amount      float64   `gorm:"column:bill_amount"                      json:"bill_amount"`
	Timestamp   time.Time `gorm:"column:bill_timestamp"                   json:"bill_timestamp"`
	Meta        string    `gorm:"column:bill_meta;type:text"              json:"bill_meta,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName override
//||------------------------------------------------------------------------------------------------||

func (Bills) TableName() string {
	return "bills"
}
