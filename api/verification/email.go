package verification

//||------------------------------------------------------------------------------------------------||
//|| CreateEmailVerification inserts a new email verification record
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/models"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| CreateEmailVerification inserts a new email verification record
//||------------------------------------------------------------------------------------------------||

func CreateEmailVerification(fidAccount string, email string, publicKeyPEM string) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt the email address
	//||------------------------------------------------------------------------------------------------||

	encrypted, err := helpers.EncryptWithPublicKey([]byte(email), publicKeyPEM)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build meta JSON
	//||------------------------------------------------------------------------------------------------||

	meta := map[string]interface{}{
		"created":    time.Now().UTC().Format(time.RFC3339),
		"approvedBy": "twoFactor",
		"info":       "email verification created",
	}

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Upsert Record
	//||------------------------------------------------------------------------------------------------||

	var existing models.Verification
	err = db.DB.
		Where("fid_account = ? AND verification_type = ?", fidAccount, "MAIL").
		First(&existing).Error

	if err == nil {
		existing.Data = string(encrypted)
		existing.Meta = string(metaJSON)
		existing.Status = "APPR"
		return db.DB.Save(&existing).Error
	}

	if err == gorm.ErrRecordNotFound {
		record := models.Verification{
			FidAccount: fidAccount,
			Type:       "MAIL",
			Data:       string(encrypted),
			Meta:       string(metaJSON),
			Status:     "APPR",
		}
		return db.DB.Create(&record).Error
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Error Handling
	//||------------------------------------------------------------------------------------------------||

	return err // any other DB error

}
