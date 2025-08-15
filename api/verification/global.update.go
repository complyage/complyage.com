package verification

//||------------------------------------------------------------------------------------------------||
//|| UpdateVerification updates an existing verification record by UUID
//||------------------------------------------------------------------------------------------------||

import (
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
//|| UpdateVerification
//|| Updates an existing verification record, appending new steps and replacing approval meta
//||------------------------------------------------------------------------------------------------||

func UpdateVerification(
	accountId int64,
	publicKey string,
	uuid string,
	display string,
	data interfaces.VerificationData,
	meta interfaces.VerificationMeta,
	secret interfaces.VerificationSecret,
	status string,
) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Existing Record
	//||------------------------------------------------------------------------------------------------||

	var existing models.Verification
	if err := db.DB.Where("verification_uuid = ? AND fid_account = ?", uuid, accountId).First(&existing).Error; err != nil {
		return errors.New("update verification: could not find record: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create New Verification Meta/Secret
	//||------------------------------------------------------------------------------------------------||

	updatedMeta := interfaces.VerificationMeta{
		Steps:    []interfaces.VerificationMetaStep{},
		Approval: interfaces.VerificationMetaApproval{},
	}

	updatedSecret := interfaces.VerificationSecret{
		CheckCode: "",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal Existsing Meta, Secret
	//||------------------------------------------------------------------------------------------------||

	var oldMeta interfaces.VerificationMeta
	if err := json.Unmarshal([]byte(existing.Meta), &oldMeta); err != nil {
		return errors.New("update verification: failed to unmarshal meta: " + err.Error())
	}

	var oldSecret interfaces.VerificationSecret
	if err := json.Unmarshal([]byte(existing.Secret), &oldSecret); err != nil {
		return errors.New("update verification: failed to unmarshal secret: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Append Steps & Replace Approval
	//||------------------------------------------------------------------------------------------------||

	updatedMeta.Steps = append(oldMeta.Steps, meta.Steps...)
	updatedMeta.Approval = meta.Approval

	//||------------------------------------------------------------------------------------------------||
	//|| Replace Secret if not empty
	//||------------------------------------------------------------------------------------------------||

	if secret.CheckCode != "" {
		updatedSecret.CheckCode = secret.CheckCode
	} else {
		updatedSecret.CheckCode = oldSecret.CheckCode
	}
	if secret.Expiration != "" {
		updatedSecret.Expiration = secret.Expiration
	} else {
		updatedSecret.Expiration = oldSecret.Expiration
	}
	if secret.Attempts > 0 {
		updatedSecret.Attempts = secret.Attempts
	} else {
		updatedSecret.Attempts = oldSecret.Attempts
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal Updated Meta
	//||------------------------------------------------------------------------------------------------||

	metaJSON, err := json.Marshal(updatedMeta)
	if err != nil {
		return errors.New("update verification: failed to marshal updated meta: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal & Encrypt New Data, Marshal Secret
	//||------------------------------------------------------------------------------------------------||

	dataJSON, err := json.Marshal(data)
	if err != nil {
		return errors.New("update verification: failed to marshal data: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt the Data
	//||------------------------------------------------------------------------------------------------||

	encrypted, err := helpers.EncryptWithPublicKey(dataJSON, publicKey)
	if err != nil {
		fmt.Println("Error encrypting data:", publicKey, err)
		return errors.New("update verification: failed to encrypt data: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Secret only if not empty
	//||------------------------------------------------------------------------------------------------||

	secretJSON, err := json.Marshal(updatedSecret)
	if err != nil {
		return errors.New("update verification: failed to marshal secret: " + err.Error())
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update Record in DB
	//||------------------------------------------------------------------------------------------------||

	verification := models.Verification{
		Display:   display,
		Data:      encrypted,
		Meta:      string(metaJSON),
		Secret:    string(secretJSON),
		Status:    status,
		UpdatedAt: time.Now(),
		// Set other fields as needed (e.g., UUID or ID if required for update)
	}

	result := db.DB.Model(&models.Verification{}).
		Where("verification_uuid = ? AND fid_account = ?", uuid, accountId).
		Updates(verification)

	if result.Error != nil {
		return errors.New("update verification: database error: " + result.Error.Error())
	}

	if result.RowsAffected == 0 {
		return errors.New("update verification: no record updated")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||

	return nil
}
