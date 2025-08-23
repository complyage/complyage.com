package verification

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

import (
	"base/constants"
	"base/helpers"
	"base/interfaces"
)

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

func CreateVerificationMAIL(accountId int64, publicKey string, email string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Credit Card Address
	//||------------------------------------------------------------------------------------------------||

	emailAddress := interfaces.EmailAddress{
		Email: email,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification Data
	//||------------------------------------------------------------------------------------------------||

	var data interfaces.VerificationData
	data.MAIL = emailAddress

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta Steps
	//||------------------------------------------------------------------------------------------------||

	step := interfaces.VerificationMetaStep{
		StepName:      "INIT",
		StepStatus:    constants.VerificationStatuses.Verified,
		StepDetails:   "Account Created - Email Verified",
		StepTimestamp: helpers.UniversalNow(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta Approval
	//||------------------------------------------------------------------------------------------------||

	approval := interfaces.VerificationMetaApproval{
		ApprovedBy:     "TWOFACTOR",
		ApprovedAt:     helpers.UniversalNow(),
		ApprovedMethod: "Email Verification",
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

	secret := interfaces.VerificationSecret{}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification
	//||------------------------------------------------------------------------------------------------||

	return CreateVerification(
		accountId,
		publicKey,
		constants.VerificationEmail,
		"APPR",
		helpers.MaskEmail(email),
		data,
		meta,
		secret,
	)
}
