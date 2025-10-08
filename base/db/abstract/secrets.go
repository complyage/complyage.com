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

func CreateKey(secret *models.ModelSecrets) (*models.ModelSecrets, error) {
	result := app.SQLDB["main"].DB.Create(secret)
	if result.Error != nil {
		return nil, result.Error
	}
	return secret, nil
}

//||------------------------------------------------------------------------------------------------||
//|| GetKeyByID
//||------------------------------------------------------------------------------------------------||

func GetKeyByID(id uint) (*models.ModelSecrets, error) {
	var secret models.ModelSecrets
	result := db.AuthDB().First(&secret, id)
	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return &secret, nil
}

//||------------------------------------------------------------------------------------------------||
//|| GetKeysByAccount
//||------------------------------------------------------------------------------------------------||

func GetKeyByAccount(fidAccount uint) (*models.ModelSecrets, error) {
	var secret models.ModelSecrets
	result := db.AuthDB().
		Where("fid_account = ?", fidAccount).
		First(&secret)

	if result.Error != nil {
		return nil, result.Error
	}
	return &secret, nil
}

//||------------------------------------------------------------------------------------------------||
//|| UpdateKeyCheck
//||------------------------------------------------------------------------------------------------||

func UpdateKeyCheck(id uint, newCheck string) error {
	result := db.AuthDB().
		Model(&models.ModelSecrets{}).
		Where("id_secret = ?", id).
		Update("check_secret", newCheck)

	return result.Error
}

//||------------------------------------------------------------------------------------------------||
//|| DeleteKey
//||------------------------------------------------------------------------------------------------||

func DeleteKey(id uint) error {
	result := db.AuthDB().Delete(&models.ModelSecrets{}, id)
	return result.Error
}
