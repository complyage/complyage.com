package verification

import (
	"base/abstract"
	"base/helpers"
	"base/interfaces"
)

func UpdateStatusIDENInProgress(uuid string, step int, details string) error {
	account, err := abstract.GetAccountByVerificationUUID(uuid)
	if err != nil {
		return err
	}
	var data interfaces.VerificationData
	newStep := interfaces.VerificationMetaStep{
		StepName:      "INPG",
		StepStatus:    "INPG",
		StepDetails:   details,
		StepTimestamp: helpers.UniversalNow(),
	}
	meta := interfaces.VerificationMeta{
		Steps: []interfaces.VerificationMetaStep{newStep},
		Step:  step,
	}
	secret := interfaces.VerificationSecret{}
	return UpdateVerification(
		account.IDAccount,
		account.AccountPublic,
		uuid,
		"ID In Progress",
		data,
		meta,
		secret,
		"INPG",
	)
}
