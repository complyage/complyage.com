package models

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Verification represents a record in the `verifications` table.
//||------------------------------------------------------------------------------------------------||

type Verify struct {
	ID         int64     `gorm:"column:id_verify;primaryKey;autoIncrement" json:"id"`
	UUID       string    `gorm:"column:verify_uuid;type:varchar(64);uniqueIndex" json:"uuid"`
	FidAccount int64     `gorm:"column:fid_account" json:"fidAccount"`
	Display    string    `gorm:"column:verify_display;type:TEXT" json:"display"`
	Type       string    `gorm:"column:verify_type;type:varchar(4)" json:"type"`
	Status     string    `gorm:"column:verify_status;type:varchar(4)" json:"status"`
	CreatedAt  time.Time `gorm:"column:verify_created;autoCreateTime" json:"created"`
	UpdatedAt  time.Time `gorm:"column:verify_updated;autoUpdateTime" json:"updated"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName
//||------------------------------------------------------------------------------------------------||

func (Verify) TableName() string {
	return "verify"
}
