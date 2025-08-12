package abstract

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/interfaces"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| Get User Verifications
//||------------------------------------------------------------------------------------------------||

func GetUserVerifications(accountID int64) ([]interfaces.UserVerification, error) {
	var dbRecords []models.Verification
	err := db.DB.
		Where("verification_type IN ('PEND', 'APPR')").
		Where("fid_account = ?", accountID).
		Find(&dbRecords).Error
	if err != nil {
		return nil, err
	}

	verifications := make([]interfaces.UserVerification, len(dbRecords))
	for i, v := range dbRecords {
		verifications[i] = interfaces.UserVerification{
			ID:     v.ID,
			Type:   v.Type,
			Data:   string(v.Data),
			Meta:   v.Meta,
			Status: v.Status,
		}
	}
	return verifications, nil
}
