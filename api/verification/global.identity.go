package verification

import (
	"base/constants"
	"base/db"
	"base/interfaces"
	"base/models"
	"encoding/json"
	"errors"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| UpdateVerifiedEmail : Updates account identity email by account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdateEmail(accountId int64, display string) error {
	return IdentityUpdateField(
		accountId,
		constants.VerificationEmail,
		func(identity *interfaces.Identity) {
			identity.Email = display
		},
	)
}

//||------------------------------------------------------------------------------------------------||
//|| Update Credit Card by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdateCreditCard(accountId int64, display string) error {
	fmt.Println(`--------------------| Identity Update | -------------------`)
	return IdentityUpdateField(
		accountId,
		constants.VerificationCreditCard,
		func(identity *interfaces.Identity) {
			identity.CreditCard = display
		},
	)
}

//||------------------------------------------------------------------------------------------------||
//|| Update Phone by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdateAddress(accountId int64, display string) error {
	fmt.Println(`--------------------| Identity Update | -------------------`)
	return IdentityUpdateField(
		accountId,
		constants.VerificationAddress,
		func(identity *interfaces.Identity) {
			identity.Phone = display
		},
	)
}

//||------------------------------------------------------------------------------------------------||
//|| Update Phone by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdatePhone(accountId int64, display string) error {
	fmt.Println(`--------------------| Identity Update | -------------------`)
	return IdentityUpdateField(
		accountId,
		constants.VerificationPhone,
		func(identity *interfaces.Identity) {
			identity.Phone = display
		},
	)
}

//||------------------------------------------------------------------------------------------------||
//|| Fetch Identity by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityFetch(accountId int64) (interfaces.Identity, error) {

	var account models.Account
	if err := db.DB.Where("id_account = ?", accountId).First(&account).Error; err != nil {
		return interfaces.Identity{}, errors.New("account not found")
	}

	var identity interfaces.Identity
	if err := json.Unmarshal([]byte(account.AccountIdentity), &identity); err != nil {
		return interfaces.Identity{}, errors.New("failed to unmarshal identity: " + err.Error())
	}

	return identity, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Update Identity by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdate(accountId int64, identity interfaces.Identity) error {

	identityJSON, err := json.Marshal(identity)
	if err != nil {
		return errors.New("failed to marshal updated identity: " + err.Error())
	}

	if err := db.DB.Model(&models.Account{}).
		Where("id_account = ?", accountId).
		Update("account_identity", identityJSON).Error; err != nil {
		return errors.New("failed to update identity in DB: " + err.Error())
	}

	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Update a Field in Identity and Mark as Approved by Account ID
//||------------------------------------------------------------------------------------------------||

func IdentityUpdateField(accountId int64, fieldType constants.VerificationType, setField func(*interfaces.Identity)) error {
	identity, fetchErr := IdentityFetch(accountId)
	if fetchErr != nil {
		return fetchErr
	}
	fmt.Println("Updating field:", fieldType, "for account ID:", accountId)
	setField(&identity)
	fmt.Println("Set field:", fieldType, "for account ID:", accountId)
	already := false
	for _, v := range identity.Approved {
		if v == fieldType {
			already = true
			break
		}
	}
	if !already {
		identity.Approved = append(identity.Approved, fieldType)
	}

	return IdentityUpdate(accountId, identity)
}
