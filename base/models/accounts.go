package models

type Account struct {
	IDAccount       int64  `gorm:"column:id_account;primaryKey;autoIncrement"`
	AccountType     string `gorm:"column:account_type;size:4"`
	AccountSalt     string `gorm:"column:account_salt;size:256"`
	AccountUsername string `gorm:"column:account_username;size:64;index:idx_accounts_account_username"`
	AccountEmail    string `gorm:"column:account_email;size:160;index:idx_accounts_account_email"`
	AccountPassword string `gorm:"column:account_password;size:256"`

	AccountSecurity int    `gorm:"column:account_security;default:1"`
	AccountPublic   string `gorm:"column:account_public;type:text"`
	AccountPrivate  string `gorm:"column:account_private;type:text"`

	AccountPrivateHash string `gorm:"column:account_private_hash;size:64"`
	AccountStatus      string `gorm:"column:account_status;size:4"`
	AccountLevel       *int8  `gorm:"column:account_level"`
	AccountIdentity    string `gorm:"column:account_identity;"`
}

// TableName sets the table name for GORM.
func (Account) TableName() string {
	return "accounts"
}
