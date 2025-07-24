package models

type Account struct {
	IDAccount       string `gorm:"column:id_account;primaryKey;autoIncrement"`
	AccountType     string `gorm:"column:account_type;size:4"`   // e.g. "USER" or "VNDR"
	AccountSalt     string `gorm:"column:account_salt;size:256"` // salt for password hashing
	AccountEmail    string `gorm:"column:account_email;size:160;index:idx_accounts_account_email"`
	AccountUsername string `gorm:"column:account_username;size:64;index:idx_accounts_account_username"`
	AccountPassword string `gorm:"column:account_password;size:256"` // final password hash with PASSWORD_PEPPER
	AccountPrivate  string `gorm:"column:account_private;type:text"` // user-owned private key (if stored — usually empty)
	AccountPublic   string `gorm:"column:account_public;type:text"`  // public key for data encryption
	AccountStatus   string `gorm:"column:account_status;size:4"`     // e.g. "VERF", "ACTV"
	AccountLevel    *int8  `gorm:"column:account_level"`             // numeric account tier
	AccountAdvanced *int8  `gorm:"column:account_advanced"`          // 1 if advanced, else 0
}

// TableName sets the table name for GORM.
func (Account) TableName() string {
	return "accounts"
}
