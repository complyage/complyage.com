package models

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Verification represents a record in the `verifications` table.
//||------------------------------------------------------------------------------------------------||

type Verification struct {
	ID         int64     `gorm:"column:id_verification;primaryKey;autoIncrement"`
	FidAccount int64     `gorm:"column:fid_account"`
	Type       string    `gorm:"column:verification_type;type:varchar(4)"`
	Data       []byte    `gorm:"column:verification_data;type:longblob"`
	Meta       string    `gorm:"column:verification_meta;type:text"`
	Status     string    `gorm:"column:verification_status;type:varchar(4)"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName
//||------------------------------------------------------------------------------------------------||

// TableName overrides the default table name.
func (Verification) TableName() string {
	return "verifications"
}
