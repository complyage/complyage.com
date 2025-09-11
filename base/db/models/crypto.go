package models

//||------------------------------------------------------------------------------------------------||
//|| Crypto Model (maps to `crypto` table)
//||------------------------------------------------------------------------------------------------||

type Crypto struct {
	IDCrypto uint   `gorm:"column:id_crypto;primaryKey;autoIncrement" json:"id"`
	Name     string `gorm:"column:crypto_name;size:60"                json:"name"`
	Prefix   string `gorm:"column:crypto_prefix;size:10"               json:"prefix"`
	Address  string `gorm:"column:crypto_address;size:255"            json:"address"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (Crypto) TableName() string {
	return "crypto"
}
