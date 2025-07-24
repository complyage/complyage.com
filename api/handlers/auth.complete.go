package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import (
	"api/verification"
	"base/abstract"
	"base/constants"
	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
	"fmt"
	"log"
	"net/http"
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

	dbAccount, err := abstract.GetAccountByID(session.ID)
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
	advanced := (r.FormValue("advanced") == "advanced")

	//||------------------------------------------------------------------------------------------------||
	//|| Password/Salt
	//||------------------------------------------------------------------------------------------------||

	passwordHash, saltHash := helpers.GeneratePassword(password)

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Private / Public Key
	//||------------------------------------------------------------------------------------------------||

	privateKey, publicKey, err := helpers.GenerateKeyPair()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate keys")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Email the Private Key to the User if Advanced
	//||------------------------------------------------------------------------------------------------||

	if advanced {
		_ = helpers.EmailPrivateKeyToUser(session.Email, privateKey)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Record
	//||------------------------------------------------------------------------------------------------||

	account := models.Account{}
	account.IDAccount = dbAccount.IDAccount
	account.AccountType = dbAccount.AccountType
	account.AccountEmail = dbAccount.AccountEmail
	account.AccountPublic = publicKey
	account.AccountPassword = passwordHash
	account.AccountSalt = saltHash
	account.AccountLevel = helpers.Int8Ptr(1) // Default level
	account.AccountStatus = "ACTV"

	if advanced {
		account.AccountPrivate = ""
		account.AccountAdvanced = helpers.BoolToInt8Ptr(true)
	} else {
		account.AccountPrivate = privateKey
		account.AccountAdvanced = helpers.BoolToInt8Ptr(false)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Email Verification
	//||------------------------------------------------------------------------------------------------||

	dbErr := verification.CreateEmailVerification(account.IDAccount, account.AccountEmail, account.AccountPublic)
	if dbErr != nil {
		log.Println("❌ Failed to create email verification:", err)
	} else {
		log.Println("✅ Email verification record created or updated")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Database Record
	//||------------------------------------------------------------------------------------------------||

	if err := db.DB.Save(&account).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not update account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"message": "Signup complete",
		"next":    "/members",
	})
}
