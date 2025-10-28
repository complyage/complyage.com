package abstract

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/base/encrypt"
)

//||------------------------------------------------------------------------------------------------||
//|| Session Load
//||------------------------------------------------------------------------------------------------||

func AccountCheckLogin(r *http.Request, withKeys bool, level int) (types.AccountSessionChecked, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		app.Log.Info("LoadSessionAccount: Missing session cookie", err.Error())
		return types.AccountSessionChecked{}, app.Err("Auth").Error("MISSING_SESSION_COOKIE")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		return types.AccountSessionChecked{}, app.Err("Auth").Error("SESSION_LOOKUP_FAILED")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Account
	//||------------------------------------------------------------------------------------------------||

	var account types.AccountSessionChecked
	account.KeysLoaded = false
	account.Private = ""
	account.Public = ""
	account.CheckKey = ""

	//||------------------------------------------------------------------------------------------------||
	//|| Get Database Account
	//||------------------------------------------------------------------------------------------------||

	if withKeys {
		accKeys, err := GetAccountWithKeys(int64(session.ID))
		if err != nil {
			return types.AccountSessionChecked{}, app.Err("Auth").Error("ACCOUNT_LOOKUP_FAILED")
		}
		account.ID = accKeys.ID
		account.Salt = accKeys.Salt
		account.Username = accKeys.Username
		account.Identifier = session.Identifier
		account.Status = accKeys.Status
		account.Level = accKeys.Level
		account.Private = accKeys.Private
		account.Public = accKeys.Public
		account.CheckKey = accKeys.CheckKey
	} else {
		acc, err := GetAccountByID(fmt.Sprintf("%d", session.ID))
		if err != nil {
			return types.AccountSessionChecked{}, app.Err("Auth").Error("ACCOUNT_LOOKUP_FAILED")
		}
		account.ID = acc.ID
		account.Salt = acc.Salt
		account.Username = acc.Username
		account.Identifier = session.Identifier
		account.Status = acc.Status
		account.Level = acc.Level
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Account Session Check
	//||------------------------------------------------------------------------------------------------||

	if account.ID == 0 || account.ID != session.ID {
		app.Log.Info("LoadSessionAccount: Account/Session mismatch", fmt.Sprintf("Account ID: %d | Session ID: %d", account.ID, session.ID))
		return types.AccountSessionChecked{}, app.Err("Auth").Error("ACCOUNT_MISMATCH")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Status Check
	//||------------------------------------------------------------------------------------------------||

	if account.Status != "ACTV" {
		app.Log.Info("LoadSessionAccount: Invalid account status", fmt.Sprintf("Account ID: %d | Status: %s", account.ID, account.Status))
		return types.AccountSessionChecked{}, app.Err("Auth").Error("ACCOUNT_STATUS_" + account.Status)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level Check
	//||------------------------------------------------------------------------------------------------||

	if account.Level < level {
		app.Log.Info("LoadSessionAccount: Insufficient account level", fmt.Sprintf("Account ID: %d | Level: %d | Required: %d", account.ID, account.Level, level))
		return types.AccountSessionChecked{}, app.Err("Auth").Error("INSUFFICIENT_ACCOUNT_LEVEL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Account Session Check
	//||------------------------------------------------------------------------------------------------||

	if withKeys && account.Private != "" && account.Public != "" && account.CheckKey != "" {
		app.Log.Debug("LoadSessionAccount: Attempting to load keys from DB")
		checkErr := encrypt.CheckPrivateKey(account.Private, account.CheckKey)
		pairErr := encrypt.CheckKeyPair(account.Private, account.Public)
		if checkErr == nil && pairErr == nil {
			account.KeysLoaded = true
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Try Session Load
	//||------------------------------------------------------------------------------------------------||

	if withKeys && !account.KeysLoaded {
		app.Log.Debug("LoadSessionAccount: Attempting to load keys from Session")
		account.Private, err = StoredPrivateKey(r)
		if err == nil {
			checkErr := encrypt.CheckPrivateKey(account.Private, account.CheckKey)
			pairErr := encrypt.CheckKeyPair(account.Private, account.Public)
			if checkErr == nil && pairErr == nil {
				account.KeysLoaded = true
			}
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	return account, nil
}
