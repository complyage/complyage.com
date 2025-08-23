package verification

import (
	"base/db"
	"base/interfaces"
	"base/models"
	"encoding/json"
	"errors"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| LoadVerificationByUUID fetches a record and converts it
//||------------------------------------------------------------------------------------------------||

func LoadVerificationRecordByUUID(uuid string) (*VerificationRecord, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Pull from DB
	//||------------------------------------------------------------------------------------------------||

	var verif models.Verification
	if err := db.DB.Where("verification_uuid = ?", uuid).First(&verif).Error; err != nil {
		return nil, err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Convert Meta
	//||------------------------------------------------------------------------------------------------||

	var meta interfaces.VerificationMeta
	if verif.Meta != "" {
		fmt.Println("Meta JSON:", verif.Meta)
		if err := json.Unmarshal([]byte(verif.Meta), &meta); err != nil {
			return nil, errors.New("invalid meta JSON: " + err.Error())
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Convert Secret
	//||------------------------------------------------------------------------------------------------||

	var secret interfaces.VerificationSecret
	if verif.Secret != "" {
		if err := json.Unmarshal([]byte(verif.Secret), &secret); err != nil {
			return nil, errors.New("invalid secret JSON: " + err.Error())
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification Record
	//||------------------------------------------------------------------------------------------------||

	rec := &VerificationRecord{
		ID:         verif.ID,
		UUID:       verif.UUID,
		FidAccount: verif.FidAccount,
		Type:       verif.Type,
		Display:    verif.Display,
		Data:       verif.Data,
		Meta:       meta,
		Secret:     secret,
		Status:     verif.Status,
		CreatedAt:  verif.CreatedAt,
		UpdatedAt:  verif.UpdatedAt,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	return rec, nil
}
