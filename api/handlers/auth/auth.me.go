package auth

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"fmt"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func AuthMeHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Live Account Record
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil || account == nil {
		responses.Error(w, http.StatusInternalServerError, "Could not retrieve account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Responses
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id":       account.IDAccount,
		"status":   account.AccountStatus,
		"type":     account.AccountType,
		"email":    session.Email,
		"username": account.AccountUsername,
		"level":    helpers.DerefInt8(account.AccountLevel),
		"security": account.AccountSecurity,
		"created":  session.Created,
		"expires":  session.Expires,
		"identity": account.AccountIdentity,
	})

}
