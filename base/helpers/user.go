package helpers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/interfaces"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| Helper: OAuthVerifications
//||------------------------------------------------------------------------------------------------||

func OAuthVerifications(fidAccount string) ([]interfaces.OAuthVerification, error) {
	// load raw verification records
	var verifs []models.Verification
	if err := db.DB.
		Where("fid_account = ?", fidAccount).
		Find(&verifs).Error; err != nil {
		return nil, err
	}

	// map to OAuthVerification
	out := make([]interfaces.OAuthVerification, len(verifs))
	for i, v := range verifs {
		out[i] = interfaces.OAuthVerification{
			Type:   v.Type,
			Status: v.Status,
			Data:   v.Data,
		}
	}

	return out, nil
}
