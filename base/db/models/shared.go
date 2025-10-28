package models

import "time"

type Shared struct {
	IDShared           uint      `gorm:"column:id_shared;primaryKey;autoIncrement" json:"id_shared"`
	FidAccount         int64     `gorm:"column:fid_account" json:"fid_account"`
	FidSite            uint      `gorm:"column:fid_site" json:"fid_site"`
	SharedType         string    `gorm:"column:shared_type;size:4" json:"shared_type"`
	SharedVerification string    `gorm:"column:shared_verification;size:36" json:"shared_verification"`
	SharedTimestamp    time.Time `gorm:"column:shared_timestamp;default:CURRENT_TIMESTAMP" json:"shared_timestamp"`
}

// TableName overrides the default table name.
func (Shared) TableName() string {
	return "shared"
}
