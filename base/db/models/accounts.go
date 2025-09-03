package models

//||------------------------------------------------------------------------------------------------||
//|| Email Verification
//||------------------------------------------------------------------------------------------------||

type Account struct {
	ID          int64  `gorm:"column:id_account;primaryKey;autoIncrement"`
	Type        string `gorm:"column:account_type;size:4"`
	Salt        string `gorm:"column:account_salt;size:256"`
	Username    string `gorm:"column:account_username;size:64;index:idx_accounts_account_username"`
	Email       string `gorm:"column:account_email;size:160;index:idx_accounts_account_email"`
	Password    string `gorm:"column:account_password;size:256"`
	Security    int    `gorm:"column:account_security;default:1"`
	Public      string `gorm:"column:account_public;type:text"`
	Private     string `gorm:"column:account_private;type:text"`
	PrivateHash string `gorm:"column:account_private_hash;size:64"`
	Status      string `gorm:"column:account_status;size:4"`
	Level       int    `gorm:"column:account_level"`
	Identity    string `gorm:"column:account_identity;"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (Account) TableName() string {
	return "accounts"
}
