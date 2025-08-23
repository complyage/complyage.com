package verification

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

import (
	"base/constants"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

func CreateVerificationIDEN(accountId int64, publicKey string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Credit Card Data
	//||------------------------------------------------------------------------------------------------||

	idCard := interfaces.Identification{}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification Data
	//||------------------------------------------------------------------------------------------------||

	var data interfaces.VerificationData
	data.IDEN = idCard

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta Steps
	//||------------------------------------------------------------------------------------------------||

	step := interfaces.VerificationMetaStep{
		StepName:      "INIT",
		StepStatus:    "INIT",
		StepDetails:   "ID Verification Initiated",
		StepTimestamp: helpers.UniversalNow(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta Approval
	//||------------------------------------------------------------------------------------------------||

	approval := interfaces.VerificationMetaApproval{
		ApprovedBy:     "",
		ApprovedAt:     "",
		ApprovedMethod: "",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta
	//||------------------------------------------------------------------------------------------------||

	meta := interfaces.VerificationMeta{
		Approval: approval,
		Steps:    []interfaces.VerificationMetaStep{step},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Secret
	//||------------------------------------------------------------------------------------------------||

	expiration := time.Now().Add(72 * time.Hour)
	secret := interfaces.VerificationSecret{
		Attempts:   0,
		Expiration: helpers.ToUniversalDate(expiration),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Base Name
	//||------------------------------------------------------------------------------------------------||

	displayName := helpers.MaskIDCard(idCard)

	//||------------------------------------------------------------------------------------------------||
	//|| Clean old Pending verifications
	//||------------------------------------------------------------------------------------------------||

	err := db.DB.Exec(`
		DELETE FROM verifications
		WHERE id_verification IN (
			SELECT id_verification FROM (
			SELECT
				id_verification,
				ROW_NUMBER() OVER (
					PARTITION BY fid_account
					ORDER BY created_at DESC
				) AS rn
			FROM verifications
				WHERE 
				verification_type = 'IDEN' AND 
				verification_status = 'PEND' AND 
				fid_account = ?				
			) AS ranked
			WHERE rn > 5
		)
	`, accountId).Error
	if err != nil {
		fmt.Println("Delete error:", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification
	//||------------------------------------------------------------------------------------------------||

	return CreateVerification(accountId, publicKey, constants.VerificationIdentification, constants.VerificationStatuses.Pending, displayName, data, meta, secret)

}
