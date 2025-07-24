package models

//||------------------------------------------------------------------------------------------------||
//|| Verification represents a record in the `verifications` table.
//||------------------------------------------------------------------------------------------------||

type Verification struct {
	IDVerification int    `gorm:"column:id_verification;primaryKey;autoIncrement:false"`
	FidAccount     string `gorm:"column:fid_account;type:varchar(256)"`
	Type           string `gorm:"column:verification_type;type:char(4)"`
	Data           string `gorm:"column:verification_data;type:text"`
	Meta           string `gorm:"column:verification_meta;type:text"`
	Status         string `gorm:"column:verifcation_status;type:char(4)"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName
//||------------------------------------------------------------------------------------------------||

// TableName overrides the default table name.
func (Verification) TableName() string {
	return "verifications"
}
