package verification

//||------------------------------------------------------------------------------------------------||
//|| Update CRCD Verification Record
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/constants"
	"base/helpers"
	"base/interfaces"
)

//||------------------------------------------------------------------------------------------------||
//|| UpdateVerificationCRCD
//|| Updates a credit card verification record by UUID
//||------------------------------------------------------------------------------------------------||

func UpdateStatusIDENPendingVerification(uuid string) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Account
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(uuid)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Data
	//||------------------------------------------------------------------------------------------------||

	var data interfaces.VerificationData

	//||------------------------------------------------------------------------------------------------||
	//|| Meta Step (to append)
	//||------------------------------------------------------------------------------------------------||

	step := interfaces.VerificationMetaStep{
		StepName:      constants.VerificationStatuses.PendingVerification,
		StepStatus:    constants.VerificationStatuses.PendingVerification,
		StepDetails:   "Placed in AI-Agent Queue for verification",
		StepTimestamp: helpers.UniversalNow(),
	}
	steps := []interfaces.VerificationMetaStep{step}

	//||------------------------------------------------------------------------------------------------||
	//|| Meta Approval (to replace)
	//||------------------------------------------------------------------------------------------------||

	approval := interfaces.VerificationMetaApproval{}

	//||------------------------------------------------------------------------------------------------||
	//|| Meta
	//||------------------------------------------------------------------------------------------------||

	meta := interfaces.VerificationMeta{
		Steps:    steps,
		Step:     1,
		Approval: approval,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Secret (only updated if provided)
	//||------------------------------------------------------------------------------------------------||

	secret := interfaces.VerificationSecret{}

	//||------------------------------------------------------------------------------------------------||
	//|| Call UpdateVerification
	//||------------------------------------------------------------------------------------------------||

	return UpdateVerification(
		account.IDAccount,
		account.AccountPublic,
		uuid,
		"ID Uploaded",
		data,
		meta,
		secret,
		constants.VerificationStatuses.PendingVerification,
	)
}
