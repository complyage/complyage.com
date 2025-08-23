package verification

//||------------------------------------------------------------------------------------------------||
//|| Update CRCD Verification Record
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/constants"
	"base/helpers"
	"base/interfaces"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| UpdateVerificationCADDR
//|| Updates a address verification record by UUID
//||------------------------------------------------------------------------------------------------||

func UpdateStatusADDRPendingVerification(uuid string, transactionId string, cardType string, lastFour string, postalCode string, address interfaces.Address) error {

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
	data.CRCD = interfaces.CreditCard{
		LastFour: lastFour,
		CardType: cardType,
		Address: interfaces.Address{
			Postal: postalCode,
		},
		TransactionId: transactionId,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Address
	//||------------------------------------------------------------------------------------------------||

	data.ADDR = address

	//||------------------------------------------------------------------------------------------------||
	//|| Meta Step (to append)
	//||------------------------------------------------------------------------------------------------||

	step := interfaces.VerificationMetaStep{
		StepName:      constants.VerificationStatuses.PendingVerification,
		StepStatus:    constants.VerificationStatuses.PendingVerification,
		StepDetails:   "Postage Transaction",
		StepTimestamp: helpers.UniversalNow(),
	}
	steps := []interfaces.VerificationMetaStep{step}

	//||------------------------------------------------------------------------------------------------||
	//|| Meta Approval (to replace)
	//||------------------------------------------------------------------------------------------------||

	approval := interfaces.VerificationMetaApproval{
		ApprovedBy:     os.Getenv("MERCHANT_NAME"),
		ApprovedAt:     helpers.UniversalNow(),
		ApprovedMethod: "MERCHANT",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Meta
	//||------------------------------------------------------------------------------------------------||

	meta := interfaces.VerificationMeta{
		Steps:    steps,
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
		helpers.MaskAddress(address.Line1, address.City, address.Country),
		data,
		meta,
		secret,
		constants.VerificationStatuses.PendingVerification,
	)
}
