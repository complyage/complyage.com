package verify

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification represents a record in the `verifications` table.
//||------------------------------------------------------------------------------------------------||

type ModelVerification struct {
	ID         int64     `gorm:"column:id_verification;primaryKey;autoIncrement" json:"id"`
	UUID       string    `gorm:"column:verification_uuid;type:varchar(64);uniqueIndex" json:"uuid"`
	FidAccount int64     `gorm:"column:fid_account" json:"fidAccount"`
	Type       string    `gorm:"column:verification_type;type:varchar(4)" json:"type"`
	Display    string    `gorm:"column:verification_display;type:TEXT" json:"display"`
	Data       []byte    `gorm:"column:verification_data;type:longblob" json:"data"`
	Meta       string    `gorm:"column:verification_meta;type:text" json:"meta"`
	Secret     string    `gorm:"column:verification_secret;type:text" json:"secret"`
	Status     string    `gorm:"column:verification_status;type:varchar(4)" json:"status"`
	CreatedAt  time.Time `gorm:"column:created_at;autoCreateTime" json:"created"`
	UpdatedAt  time.Time `gorm:"column:updated_at;autoUpdateTime" json:"updated"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName
//||------------------------------------------------------------------------------------------------||

// TableName overrides the default table name.
func (ModelVerification) TableName() string {
	return "verifications"
}
