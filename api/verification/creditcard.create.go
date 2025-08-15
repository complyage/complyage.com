package verification

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

import (
	"base/constants"
	"base/helpers"
	"base/interfaces"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Create CRCD
//||------------------------------------------------------------------------------------------------||

func CreateVerificationCRCD(accountId int64, publicKey string, checkCode string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Credit Card Data
	//||------------------------------------------------------------------------------------------------||

	creditCard := interfaces.CreditCard{
		LastFour:      "",
		CardType:      "",
		Address:       interfaces.Address{},
		TransactionId: "",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification Data
	//||------------------------------------------------------------------------------------------------||

	var data interfaces.VerificationData
	data.CRCD = creditCard

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Meta Steps
	//||------------------------------------------------------------------------------------------------||

	step := interfaces.VerificationMetaStep{
		StepName:      "INIT",
		StepStatus:    "INIT",
		StepDetails:   "Credit Card Verification Initiated",
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
		CheckCode:  checkCode,
		Attempts:   0,
		Expiration: helpers.ToUniversalDate(expiration),
	}
	fmt.Printf("DEBUG expiration: %s\n", expiration)
	fmt.Printf("DEBUG secret: %+v\n", secret)
	fmt.Printf("DEBUG ToUniversalDate: %s\n", helpers.ToUniversalDate(expiration))

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification
	//||------------------------------------------------------------------------------------------------||

	return CreateVerification(accountId, publicKey, constants.VerificationCreditCard, constants.VerificationStatuses.Pending, data, meta, secret)

}
