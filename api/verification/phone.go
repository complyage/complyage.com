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
//|| CheckPhoneVerification. Checks if a phone verification record exists for the given fidAccount
//||------------------------------------------------------------------------------------------------||

func CheckPhoneVerification(fidAccount int64) (string, error) {
	var existing models.Verification
	err := db.DB.
		Where("fid_account = ? AND verification_type = ?", fidAccount, "PHNE").
		First(&existing).Error

	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return "", nil // No record found
		}
		return "", err // Return any other DB error
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decrypt the phone number
	//||------------------------------------------------------------------------------------------------||

	decrypted, err := helpers.DecryptWithPrivateKey(existing.Data, existing.Meta)
	if err != nil {
		return "", err
	}

	return string(decrypted), nil
}

//||------------------------------------------------------------------------------------------------||
//|| CreatePhoneVerification inserts a new phone verification record
//||------------------------------------------------------------------------------------------------||

func CreatePhoneVerification(fidAccount int64, phone string, publicKeyPEM string) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt the email address
	//||------------------------------------------------------------------------------------------------||

	encrypted, err := helpers.EncryptWithPublicKey([]byte(phone), publicKeyPEM)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build meta JSON
	//||------------------------------------------------------------------------------------------------||

	meta := map[string]interface{}{
		"created":    time.Now().UTC().Format(time.RFC3339),
		"approvedBy": "twoFactor",
		"info":       "phone verification created",
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
		Where("fid_account = ? AND verification_type = ?", fidAccount, "PHNE").
		First(&existing).Error

	if err != nil && err != gorm.ErrRecordNotFound {
		return err // return if there's any other DB error
	}

	if err == gorm.ErrRecordNotFound {
		record := models.Verification{
			FidAccount: fidAccount,
			Type:       "PHNE",
			Data:       encrypted,
			Meta:       string(metaJSON),
			Status:     "APPR",
		}
		return db.DB.Create(&record).Error
	}

	existing.Data = encrypted
	existing.Meta = string(metaJSON)
	existing.Status = "APPR"

	//||------------------------------------------------------------------------------------------------||
	//|| Error Handling
	//||------------------------------------------------------------------------------------------------||

	return db.DB.Save(&existing).Error

}
