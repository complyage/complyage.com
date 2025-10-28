package abstract

import (
	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Account With Keys
//||------------------------------------------------------------------------------------------------||

func GetAccountWithKeys(id int64) (*models.ModelAccountKeys, error) {
	var result models.ModelAccountKeys
	if err := app.SQLDB["main"].DB.Table("accounts").
		Select(`accounts.id_account,
                accounts.account_salt,
                accounts.account_username,
                accounts.account_identifier,
                accounts.account_status,
                accounts.account_level,
                secrets.id_secret,
                secrets.secret_private,
                secrets.secret_public,
                secrets.secret_check`).
		Joins("JOIN secrets ON secrets.fid_account = accounts.id_account").
		Where("accounts.id_account = ?", id).
		First(&result).Error; err != nil {
		return nil, err
	}
	return &result, nil
}
