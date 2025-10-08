//||------------------------------------------------------------------------------------------------||
//|| ModelAccountKeys
//||------------------------------------------------------------------------------------------------||

package models

//||------------------------------------------------------------------------------------------------||
//|| ModelAccountKeys
//||------------------------------------------------------------------------------------------------||

type ModelAccountKeys struct {
	ID         int64  `gorm:"column:id_account"`
	Salt       string `gorm:"column:account_salt"`
	Username   string `gorm:"column:account_username"`
	Identifier string `gorm:"column:account_identifier"`
	Status     string `gorm:"column:account_status"`
	Level      int    `gorm:"column:account_level"`
	Private    string `gorm:"column:secret_private"`
	Public     string `gorm:"column:secret_public"`
	CheckKey   string `gorm:"column:secret_check"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name Hint (anchors this struct to accounts)
//||------------------------------------------------------------------------------------------------||

func (ModelAccountKeys) TableName() string {
	return "accounts"
}
