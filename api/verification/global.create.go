package verification

//||------------------------------------------------------------------------------------------------||
//|| CreateEmailVerification inserts a new email verification record
//||------------------------------------------------------------------------------------------------||

import (
	"base/constants"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| CreateEmailVerification inserts a new email verification record
//||------------------------------------------------------------------------------------------------||

func CreateVerification(accountId int64, publicKey string, vType constants.VerificationType, vStatus string, data interfaces.VerificationData, meta interfaces.VerificationMeta, secret interfaces.VerificationSecret) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build meta JSON
	//||------------------------------------------------------------------------------------------------||

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return "", errors.New("Failed to marshal data: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build meta JSON
	//||------------------------------------------------------------------------------------------------||

	metaJSON, err := json.Marshal(meta)
	if err != nil {
		return "", errors.New("Failed to marshal meta: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build secret JSON
	//||------------------------------------------------------------------------------------------------||

	secretJSON, err := json.Marshal(secret)
	if err != nil {
		return "", errors.New("Failed to marshal secret: " + err.Error())
	}
	fmt.Printf("DEBUG secretJSON: %s\n", secretJSON)

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt the email address
	//||------------------------------------------------------------------------------------------------||

	encrypted, err := helpers.EncryptWithPublicKey(dataJSON, publicKey)
	if err != nil {
		return "", errors.New("Failed to encrypt data: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| UUID
	//||------------------------------------------------------------------------------------------------||

	uuid := helpers.GenerateUUID()

	//||------------------------------------------------------------------------------------------------||
	//|| Create Model
	//||------------------------------------------------------------------------------------------------||

	verification := models.Verification{
		UUID:       uuid,
		FidAccount: accountId,
		Type:       string(vType),
		Data:       encrypted,
		Meta:       string(metaJSON),
		Secret:     string(secretJSON),
		Status:     vStatus,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Insert
	//||------------------------------------------------------------------------------------------------||

	dbResult := db.DB.Debug().Create(&verification)
	if dbResult.Error != nil {
		return "", errors.New("Failed to create verification record: " + dbResult.Error.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||

	return uuid, nil

}
