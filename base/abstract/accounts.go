package abstract

import (
	"base/db"
	"base/models"

	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByID(id string) (*models.Account, error) {
	var account models.Account

	result := db.DB.First(&account, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Not found, not a DB error
		}
		return nil, result.Error
	}

	return &account, nil
}
