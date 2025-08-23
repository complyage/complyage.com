package auth

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import (
	"base/abstract"
	"base/constants"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"base/responses"
	"base/verify"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func CompleteHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		responses.Error(w, http.StatusUnauthorized, "Missing or invalid session")
		return
	}
	fmt.Println("Complete Cookie:", cookie.Value)

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Database Account
	//||------------------------------------------------------------------------------------------------||

	dbAccount, err := abstract.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil || dbAccount == nil {
		responses.Error(w, http.StatusInternalServerError, "Could not retrieve account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Account Status
	//||------------------------------------------------------------------------------------------------||

	if dbAccount.AccountStatus != constants.AccountStatus.Verified {
		responses.Error(w, http.StatusForbidden, "Account is already created")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	password := r.FormValue("password")
	rawEncrypt := r.FormValue("encryptionLevel")
	privateKeyInput := r.FormValue("privateKey")
	publicKeyInput := r.FormValue("publicKey")
	wordListJSON := r.FormValue("wordList")

	//||------------------------------------------------------------------------------------------------||
	//||
	//|| Sanitize and Validate
	//|| Also generate the private/public key if needed
	//||
	//||------------------------------------------------------------------------------------------------||

	if password == "" || len(password) < 8 {
		responses.Error(w, http.StatusBadRequest, "Password must be at least 8 characters long")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Encryption Level
	//||------------------------------------------------------------------------------------------------||

	encryptionLevel, err := strconv.Atoi(rawEncrypt)
	if err != nil || encryptionLevel < 1 || encryptionLevel > 3 {
		responses.Error(w, http.StatusBadRequest, "Invalid or missing encryption level")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate and Generate Private/Public Key
	//||------------------------------------------------------------------------------------------------||

	var privateKey, publicKey string
	var BIPList interfaces.BIPList

	//||------------------------------------------------------------------------------------------------||
	//|| Level 1 - We handle the keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 1 {
		genPrivateKey, genPublicKey, err := helpers.GenerateKeyPair()
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to generate keys")
			return
		}
		privateKey = genPrivateKey
		publicKey = genPublicKey
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 2 - BIPList
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 2 {
		BIPList, err := helpers.ValidateBIP39(wordListJSON)
		if err != nil {
			responses.Error(w, http.StatusBadRequest, "Invalid BIP39 word list: "+err.Error())
			return
		}
		genPrivate, genPublic, err := helpers.GenerateBIP39Keys(BIPList)
		if err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to generate BIP39 keys")
			return
		}
		privateKey = genPrivate
		publicKey = genPublic
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 3 requires both keys
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 3 {
		err := helpers.ValidateKeyPair(privateKeyInput, publicKeyInput)
		if err != nil {
			responses.Error(w, http.StatusBadRequest, "Invalid key pair: "+err.Error())
			return
		}
		privateKey = privateKeyInput
		publicKey = publicKeyInput
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Private Key Hash
	//||------------------------------------------------------------------------------------------------||

	privateKeyHash, err := helpers.GenerateCheckKey(privateKey)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate private key hash")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Password/Salt
	//||------------------------------------------------------------------------------------------------||

	passwordHash, saltHash := helpers.GeneratePassword(password, "")
	if passwordHash == "" {
		responses.Error(w, http.StatusBadRequest, "Could not generate password")
	}

	//||------------------------------------------------------------------------------------------------||
	//||
	//|| Contact the User with the keys if needed
	//||
	//||------------------------------------------------------------------------------------------------||

	if encryptionLevel == 1 {
		_ = helpers.EmailPrivateKeyToUser(session.Email, privateKey)
	}

	if encryptionLevel == 2 {
		_ = helpers.EmailBIPListToUser(session.Email, BIPList, privateKey)
	}

	if encryptionLevel == 3 {
		_ = helpers.EmailPrivateKeyToUser(session.Email, privateKey)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Random Usenrame
	//||------------------------------------------------------------------------------------------------||

	randomUsername, err := helpers.GenerateUsername()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate username")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Record
	//||------------------------------------------------------------------------------------------------||

	account := models.Account{}
	account.IDAccount = dbAccount.IDAccount
	account.AccountUsername = dbAccount.AccountUsername
	account.AccountType = dbAccount.AccountType
	account.AccountEmail = dbAccount.AccountEmail
	account.AccountUsername = randomUsername
	account.AccountPublic = publicKey
	account.AccountPassword = passwordHash
	account.AccountPrivateHash = privateKeyHash
	account.AccountSalt = saltHash
	account.AccountLevel = helpers.Int8Ptr(1) // Default level
	account.AccountStatus = "ACTV"
	account.AccountSecurity = encryptionLevel // Default security level
	account.AccountIdentity = string("{}")
	account.AccountPrivate = privateKey
	app.SQLDB["main"].DB.Save(&account)

	//||------------------------------------------------------------------------------------------------||
	//|| Email Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Init(verify.DataTypeMAIL, dbAccount.IDAccount, app.Storages["verifications"], app.SQLDB["main"], privateKey, publicKey)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to initialize verification: "+err.Error())
		return
	}
	verifyRecord.UpdateStatusVerified(0) // Automatic Moderator

	//||------------------------------------------------------------------------------------------------||
	//|| Refetch the User Data
	//||------------------------------------------------------------------------------------------------||

	updatedAccount, err := abstract.GetAccountByID(fmt.Sprintf("%d", account.IDAccount))
	if err != nil || updatedAccount == nil {
		responses.Error(w, http.StatusInternalServerError, "Could not re-fetch account after update")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Session
	//||------------------------------------------------------------------------------------------------||

	sessionToken, err := helpers.SessionCreate(updatedAccount.AccountEmail, *updatedAccount)
	if err != nil || sessionToken == "" {
		responses.Error(w, http.StatusInternalServerError, "Failed to create session: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Write the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	helpers.WriteSessionCookie(w, sessionToken)

	//||------------------------------------------------------------------------------------------------||
	//|| Delete the Old Session Cookie
	//||------------------------------------------------------------------------------------------------||

	if cookie.Value != "" && cookie.Value != sessionToken {
		_ = helpers.DeleteSession(cookie.Value)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"message": "Signup complete",
		"next":    "/members",
	})
}
