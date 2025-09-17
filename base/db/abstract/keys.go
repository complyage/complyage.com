package abstract

import (
	"errors"

	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/db"
	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| CreateKey
//||------------------------------------------------------------------------------------------------||

func CreateKey(key *models.ModelKey) (*models.ModelKey, error) {
	result := app.SQLDB["main"].DB.Create(key)
	if result.Error != nil {
		return nil, result.Error
	}
	return key, nil
}

//||------------------------------------------------------------------------------------------------||
//|| GetKeyByID
//||------------------------------------------------------------------------------------------------||

func GetKeyByID(id uint) (*models.ModelKey, error) {
	var key models.ModelKey
	result := db.AuthDB().First(&key, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &key, nil
}

//||------------------------------------------------------------------------------------------------||
//|| GetKeysByAccount
//||------------------------------------------------------------------------------------------------||

func GetKeyByAccount(fidAccount uint) (models.ModelKey, error) {
	var key models.ModelKey
	result := db.AuthDB().
		Where("fid_account = ?", fidAccount).
		First(&key)

	if result.Error != nil {
		return models.ModelKey{}, result.Error
	}
	return key, nil
}

//||------------------------------------------------------------------------------------------------||
//|| UpdateKeyCheck
//||------------------------------------------------------------------------------------------------||

func UpdateKeyCheck(id uint, newCheck string) error {
	result := db.AuthDB().
		Model(&models.ModelKey{}).
		Where("id_key = ?", id).
		Update("check_key", newCheck)

	return result.Error
}

//||------------------------------------------------------------------------------------------------||
//|| DeleteKey
//||------------------------------------------------------------------------------------------------||

func DeleteKey(id uint) error {
	result := db.AuthDB().Delete(&models.ModelKey{}, id)
	return result.Error
}
