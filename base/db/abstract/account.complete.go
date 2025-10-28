package abstract

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/encrypted"
	"github.com/complyage/base/identity"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/db"
	"github.com/ralphferrara/aria/base/validate"
)

func OnAccountComplete(r *http.Request, accountId int64, accountIdentifier string) error {

	//||------------------------------------------------------------------------------------------------||
	//|| DB Account
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Fetching the created account, ID : ", accountId)
	dbAccount, err := db.GetAccountByID(fmt.Sprintf("%d", accountId))
	if err != nil {
		return app.Err("Auth").Error("ACCOUNT_LOOKUP_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Account Status
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("dbAccount.Status:", dbAccount.Status)
	app.Log.Info("Expected Status:", app.Constants("AccountStatus").Code("Verified"))
	if dbAccount.Status != app.Constants("AccountStatus").Code("Verified") {
		return app.Err("auth").Error("ACCOUNT_NOT_PENDING")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Key Generation
	//||------------------------------------------------------------------------------------------------||

	var encryptionLevel int
	var privateKey, publicKey, privateCheck string
	keysExist := false

	//||------------------------------------------------------------------------------------------------||
	//|| Set the existing keys if they exist
	//||------------------------------------------------------------------------------------------------||

	existingKey, err := GetKeyByAccount(uint(accountId)) // write this helper if not present
	if err == nil && existingKey != nil {
		app.Log.Info("Keys already exist, loading them instead of generating new")
		encryptionLevel = existingKey.Level
		privateKey = existingKey.Private
		publicKey = existingKey.Public
		privateCheck = existingKey.CheckKey
		keysExist = true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Keys
	//||------------------------------------------------------------------------------------------------||

	if !keysExist {
		secrets, secretsErr := CreateAccountSecrets(r, accountId, accountIdentifier)
		if secretsErr != nil {
			return secretsErr
		}
		privateKey = secrets.PrivateKey
		publicKey = secrets.PublicKey
		privateCheck = secrets.CheckKey
		encryptionLevel = secrets.EncryptionLevel
	}

	app.Log.Info("Encryption Level:", encryptionLevel)
	app.Log.Info("Private Key:", len(privateKey), " characters")
	app.Log.Info("Public Key:", len(publicKey), " characters")
	app.Log.Info("Private Key Hash:", privateCheck)

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Identity
	//||------------------------------------------------------------------------------------------------||

	iden, err := identity.Create(dbAccount.ID)
	if err != nil {
		return app.Err("Auth").Error("IDENTITY_CREATE_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Identifier Type
	//||------------------------------------------------------------------------------------------------||

	identifierType := validate.IsEmailOrPhone(accountIdentifier)
	app.Log.Info("Getting identifier type -", accountIdentifier)
	app.Log.Info("Identifier Type:", identifierType)
	if identifierType == "unknown" {
		return app.Err("Auth").Error("INVALID_IDENTIFIER")
	}
	verifyType := types.DataTypeMAIL
	if identifierType == "phone" {
		verifyType = types.DataTypePHNE
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Verification Record
	//||------------------------------------------------------------------------------------------------||

	app.Log.Info("Creating Verification Record")
	verifyRecord := verify.Create(verifyType, types.AccountSessionChecked{
		ID:         dbAccount.ID,
		Salt:       dbAccount.Salt,
		Username:   dbAccount.Username,
		Identifier: dbAccount.Identifier,
		Status:     dbAccount.Status,
		Level:      dbAccount.Level,
		KeysLoaded: true,
		Private:    privateKey,
		Public:     publicKey,
		CheckKey:   privateCheck,
	})
	app.Log.Info("Creating Verification Record with UUID:", verifyRecord.UUID)

	//||------------------------------------------------------------------------------------------------||
	//|| Set the Encrypted/Identity Data
	//||------------------------------------------------------------------------------------------------||

	var encErr error
	if verifyType == types.DataTypeMAIL {
		//||------------------------------------------------------------------------------------------------||
		//|| Create the type
		//||------------------------------------------------------------------------------------------------||
		var email types.EmailAddress
		email.Email = accountIdentifier
		//||------------------------------------------------------------------------------------------------||
		//|| Set the Verify Record
		//||------------------------------------------------------------------------------------------------||
		verifyRecord.SetDataEmail(email)
		//||------------------------------------------------------------------------------------------------||
		//|| Set the Identity
		//||------------------------------------------------------------------------------------------------||
		iden.SetVerification(types.DataTypeMAIL.String(), true, email.Mask(), verifyRecord.UUID)
		//||------------------------------------------------------------------------------------------------||
		//|| Encrypt the Data
		//||------------------------------------------------------------------------------------------------||
		encErr = encrypted.SaveMAIL(publicKey, verifyRecord.UUID, email)
		verifyRecord.EncryptedSaved = true
	} else {
		//||------------------------------------------------------------------------------------------------||
		//|| Create the type
		//||------------------------------------------------------------------------------------------------||
		phone, err := types.PhoneFromString(accountIdentifier)
		if err != nil {
			return app.Err("Types").Error("INVALID_PHONE")
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Set the Verify Record
		//||------------------------------------------------------------------------------------------------||
		verifyRecord.SetDataPhone(phone)
		//||------------------------------------------------------------------------------------------------||
		//|| Set the Identity
		//||------------------------------------------------------------------------------------------------||
		iden.SetVerification(types.DataTypePHNE.String(), true, phone.Mask(), verifyRecord.UUID)
		//||------------------------------------------------------------------------------------------------||
		//|| Encrypt the Data
		//||------------------------------------------------------------------------------------------------||
		encErr = encrypted.SavePHNE(publicKey, verifyRecord.UUID, phone)
		verifyRecord.EncryptedSaved = true
	}
	if encErr != nil {
		return app.Err("Encrypted").Error("ENCRYPT_CREATE_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save the Identity
	//||------------------------------------------------------------------------------------------------||

	iErr := iden.Save()
	if iErr != nil {
		return app.Err("Identity").Error("IDENTITY_UPDATE_FAILED")
	}
	verifyRecord.IdentityUpdated = true

	//||------------------------------------------------------------------------------------------------||
	//|| Save Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.UpdateStatusVerified("TWOFACTOR")
	verifyRecord.Save()
	verifyRecord.DatabaseUpdate()

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	return nil
}
