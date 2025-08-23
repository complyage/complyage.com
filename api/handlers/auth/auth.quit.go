package auth

import (
	"base/abstract"
	"base/helpers"
	"base/responses"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| generateQuitToken – creates an HMAC hash from AccountPrivateHash
//||------------------------------------------------------------------------------------------------||

func generateQuitToken(accountHash string) string {
	secret := []byte(os.Getenv("CSRF_SECRET"))
	h := hmac.New(sha256.New, secret)
	h.Write([]byte(accountHash))
	return hex.EncodeToString(h.Sum(nil))
}

//||------------------------------------------------------------------------------------------------||
//|| QuitHandler – serves token for delete confirmation
//||------------------------------------------------------------------------------------------------||

func QuitHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	account, err := abstract.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Account not found")
		return
	}

	token := generateQuitToken(account.AccountPrivateHash)

	responses.Success(w, http.StatusOK, map[string]any{
		"quitToken": token,
	})
}

// ||------------------------------------------------------------------------------------------------||
// || DeleteAccountHandler – validates token and deletes account
// ||------------------------------------------------------------------------------------------------||

func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	clientToken := r.FormValue("quitToken")

	account, err := abstract.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Account not found")
		return
	}

	expected := generateQuitToken(account.AccountPrivateHash)
	if clientToken != expected {
		responses.Error(w, http.StatusForbidden, "Invalid token")
		return
	}

	if err := abstract.DeleteAccount(account.IDAccount); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to delete account")
		return
	}

	responses.Success(w, http.StatusOK, map[string]string{
		"message": "Account deleted successfully",
	})
}
