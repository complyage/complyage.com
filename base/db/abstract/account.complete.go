package abstract

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/send"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/db"
	"github.com/ralphferrara/aria/base/crypto"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/base/validate"
)

func OnAccountComplete(r *http.Request, accountID int64, accountIdentifier string) error {

	//||------------------------------------------------------------------------------------------------||
	//|| DB Account
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Fetching the created account, ID : ", accountID)
	dbAccount, err := db.GetAccountByID(fmt.Sprintf("%d", accountID))
	if err != nil {
		return app.Err("Auth").Error("ACCOUNT_LOOKUP_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Account Status
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("dbAccount.Status:", dbAccount.Status)
	fmt.Println("Expected Status:", app.Constants("AccountStatus").Code("Verified"))
	if dbAccount.Status != app.Constants("AccountStatus").Code("Verified") {
		return app.Err("auth").Error("ACCOUNT_NOT_PENDING")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	rawEncrypt := r.FormValue("encryptionLevel")
	privateKeyInput := r.FormValue("privateKey")
	publicKeyInput := r.FormValue("publicKey")
	wordListJSON := r.FormValue("wordList")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Encryption Level
	//||------------------------------------------------------------------------------------------------||

	encryptionLevel, err := strconv.Atoi(rawEncrypt)
	if err != nil || encryptionLevel < 1 || encryptionLevel > 3 {
		return app.Err("Auth").Error("INVALID_ACCOUNT_TYPE")
	}

	fmt.Println("Encryption Level = ", encryptionLevel)

	//||------------------------------------------------------------------------------------------------||
	//|| Validate and Generate Private/Public Key
	//||------------------------------------------------------------------------------------------------||

	var privateKey, publicKey string

	//||------------------------------------------------------------------------------------------------||
	//|| Level 1 - We handle the keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 1 {
		genPrivateKey, genPublicKey, err := crypto.GenerateKeyPair()
		if err != nil {
			return app.Err("Auth").Error("PRIVPUB_FAILED")
		}
		privateKey = genPrivateKey
		publicKey = genPublicKey
	}
	fmt.Println("Private Key:", privateKey)
	fmt.Println("Public Key:", publicKey)

	//||------------------------------------------------------------------------------------------------||
	//|| Level 2 - BIPList
	//||------------------------------------------------------------------------------------------------||

	var BIPList []string

	if encryptionLevel == 2 {
		BIPList, err := validate.ValidateBIP39(wordListJSON)
		if err != nil {
			return app.Err("Auth").Error("BAD_BIP39")
		}
		genPrivate, genPublic, err := crypto.GenerateBIP39Keys(BIPList)
		if err != nil {
			return app.Err("Auth").Error("BIP39_GEN_FAILED")
		}
		privateKey = genPrivate
		publicKey = genPublic
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 3 requires both keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 3 {
		privateKey = privateKeyInput
		publicKey = publicKeyInput
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Key Pair
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Validating Key Pair")
	err = validate.ValidateKeyPair(privateKey, publicKey)
	if err != nil {
		fmt.Println("Error validating key pair:", err)
		return app.Err("Auth").Error("PRIVPUB_MISMATCH")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Private Key Hash
	//||------------------------------------------------------------------------------------------------||

	privateKeyHash, err := encrypt.GenerateCheckKey(privateKey)
	if err != nil {
		fmt.Println("Error generating private key hash:", err)
		return app.Err("Auth").Error("PRIVPUB_CHECKEY_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Contact the User with the keys if needed
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Sending keys to user")
	if encryptionLevel == 1 {
		_ = send.EmailPrivateKeyToUser(dbAccount.Identifier, privateKey)
	}

	if encryptionLevel == 2 {
		_ = send.EmailBIPListToUser(dbAccount.Identifier, BIPList, privateKey)
	}

	if encryptionLevel == 3 {
		_ = send.EmailPrivateKeyToUser(dbAccount.Identifier, privateKey)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Key to DB
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Creating Encryption Type")
	if encryptionLevel == 1 {
		CreateKey(&models.ModelKey{
			FidAccount: uint(dbAccount.ID),
			Level:      encryptionLevel,
			Private:    privateKey,
			Public:     publicKey,
			CheckKey:   privateKeyHash,
			CreatedAt:  time.Now().UTC(),
		})
	} else {
		CreateKey(&models.ModelKey{
			FidAccount: uint(dbAccount.ID),
			Level:      encryptionLevel,
			Private:    "",
			Public:     publicKey,
			CheckKey:   privateKeyHash,
			CreatedAt:  time.Now().UTC(),
		})
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Identifier Type
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Getting identifier type")
	identifierType := validate.IsEmailOrPhone(accountIdentifier)
	fmt.Println("Identifier Type:", identifierType)
	if identifierType == "unknown" {
		return app.Err("Auth").Error("INVALID_IDENTIFIER")
	}

	verifyType := verify.DataTypeMAIL
	if identifierType == "phone" {
		verifyType = verify.DataTypePHNE
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Email Verification
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Creating Verification Record")
	verifyRecord, err := verify.Create(verifyType, dbAccount.ID, app.Storages["verifications"], app.SQLDB["main"], privateKey, publicKey)
	if err != nil {
		return app.Err("Verify").Error("CREATE_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Determine Email or Phone
	//||------------------------------------------------------------------------------------------------||

	if verifyType == verify.DataTypeMAIL {
		var email types.EmailAddress
		email.Email = accountIdentifier
		verifyRecord.SetDataEmail(email)
		verifyRecord.Identity.SetVerification(verifyType.String(), true, email.Mask(), verifyRecord.UUID)
	} else {
		phone, err := types.PhoneFromString(accountIdentifier)
		if err != nil {
			return app.Err("Auth").Error("INVALID_PHONE")
		}
		verifyRecord.SetDataPhone(phone)
		verifyRecord.Identity.SetVerification(verifyType.String(), true, phone.Mask(), verifyRecord.UUID)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Determine Email or Phone
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.DatabaseSaveIdentity()
	verifyRecord.UpdateStatusVerified("TWOFACTOR")
	verifyRecord.Save()
	verifyRecord.Lock()

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	return nil
}
