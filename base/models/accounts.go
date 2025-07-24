package models

type Account struct {
	IDAccount           string `gorm:"column:id_account;primaryKey;autoIncrement"`
	AccountType         string `gorm:"column:account_type;size:4"`
	AccountSalt         string `gorm:"column:account_salt;size:256"`
	AccountEmail        string `gorm:"column:account_email;size:160;index:idx_accounts_account_email"`
	AccountUsername     string `gorm:"column:account_username;size:64;index:idx_accounts_account_username"`
	AccountPassword     string `gorm:"column:account_password;size:256"`
	AccountPrivate      string `gorm:"column:account_private;type:text"`
	AccountPrivateCheck string `gorm:"column:account_private_hash;size:64"`
	AccountPublic       string `gorm:"column:account_public;type:text"`
	AccountStatus       string `gorm:"column:account_status;size:4"`
	AccountLevel        *int8  `gorm:"column:account_level"`
	AccountAdvanced     *int8  `gorm:"column:account_advanced"`
}

// TableName sets the table name for GORM.
func (Account) TableName() string {
	return "accounts"
}
