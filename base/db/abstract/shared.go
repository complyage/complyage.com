package abstract

import (
	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Transaction
//||------------------------------------------------------------------------------------------------||

type SharedTransaction struct {
	AccountId     int64
	SiteId        uint
	Types         string
	Verifications SharedVerifications
}

type SharedVerifications []SharedVerification

type SharedVerification struct {
	Type         string
	Verification string
}

//||------------------------------------------------------------------------------------------------||
//|| Register
//||------------------------------------------------------------------------------------------------||

func RegisterShared(shared SharedTransaction) error {
	//||------------------------------------------------------------------------------------------------||
	//|| Start Transaction
	//||------------------------------------------------------------------------------------------------||
	tx := app.SQLDB["main"].DB.Begin()
	if tx.Error != nil {
		return tx.Error
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Loop Though Verifications
	//||------------------------------------------------------------------------------------------------||
	for _, v := range shared.Verifications {
		s := models.Shared{
			FidAccount:         shared.AccountId,
			FidSite:            shared.SiteId,
			SharedType:         v.Type,
			SharedVerification: v.Verification,
		}

		if err := tx.Create(&s).Error; err != nil {
			tx.Rollback()
			return err
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Commit
	//||------------------------------------------------------------------------------------------------||
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
