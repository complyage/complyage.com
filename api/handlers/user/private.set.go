package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/complyage/base/db/abstract"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/base/bip39"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Request / Response envelopes
//||------------------------------------------------------------------------------------------------||

type userPrivateSetRequest struct {
	Words      bip39.BIP39WordList `json:"words,omitempty"`
	PrivateKey string              `json:"privateKey,omitempty"`
	Minutes    int                 `json:"minutes,omitempty"` // how long to store in redis (minutes). 0 => indefinite
}

type userPrivateSetResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserPrivateSet - Store / update an account's private key (encrypted + transient redis copy)
//||------------------------------------------------------------------------------------------------||

func UserPrivateSet(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse request body
	//||------------------------------------------------------------------------------------------------||

	var req userPrivateSetRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, app.Err("JSON").Code("BAD_PAYLOAD"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate session / account
	//||------------------------------------------------------------------------------------------------||

	_, account, _, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, app.Err("AUTH").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get user's key metadata (to determine level)
	//||------------------------------------------------------------------------------------------------||

	keys, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, app.Err("DB").Code("KEY_LOOKUP"))
		return
	}
	level := keys.Level

	//||------------------------------------------------------------------------------------------------||
	//|| Level 1 : We manage the key
	//||------------------------------------------------------------------------------------------------||

	if level == 1 {
		responses.Success(w, http.StatusBadRequest, app.Err("KEY").Code("LEVEL1_MANAGED"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Levels 2-5 : Words List
	//||------------------------------------------------------------------------------------------------||

	var privateKeyPEM string
	fmt.Println("Level:", level)
	if level >= 2 && level <= 5 {
		expected := abstract.KeyLevelToWordCount(level)
		bip39.ValidateBIP39List(req.Words, expected)

		priv, _, genErr := bip39.GenerateBIP39Keys(req.Words)
		if genErr != nil {
			responses.Error(w, http.StatusInternalServerError, app.Err("CRYPTO").Code("BIP39_GEN_FAIL"))
			return
		}
		privateKeyPEM = strings.TrimSpace(priv)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level 6 : Private Key
	//||------------------------------------------------------------------------------------------------||

	if level == 6 {
		req.PrivateKey = strings.TrimSpace(req.PrivateKey)
		if req.PrivateKey == "" {
			responses.Error(w, http.StatusBadRequest, app.Err("KEY").Code("EMPTY"))
			return
		}

		if encrypt.CheckPrivateKey(req.PrivateKey, keys.CheckKey) != nil {
			responses.Error(w, http.StatusBadRequest, app.Err("KEY").Code("MISMATCH"))
			return
		}

		privateKeyPEM = req.PrivateKey
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	if privateKeyPEM == "" {
		responses.Error(w, http.StatusBadRequest, app.Err("KEY").Code("NO_KEY"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Use Check Key to ensure we have a key match
	//||------------------------------------------------------------------------------------------------||

	if err := encrypt.CheckPrivateKey(privateKeyPEM, keys.CheckKey); err != nil {
		responses.Error(w, http.StatusBadRequest, app.Err("KEY").Code("KEY_MISMATCH"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Transient store: save plaintext in Redis (for X minutes) and set cookie with token
	//||------------------------------------------------------------------------------------------------||

	minutes := req.Minutes
	if minutes == 0 {
		minutes = 60 * 24 * 365 * 10 // default: 10 years in minutes
	}

	seconds := minutes * 60

	token, err := abstract.SetPrivateKey(r, privateKeyPEM, seconds)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, app.Err("REDIS").Code("SAVE_FAIL"))
		return
	}

	cookie := &http.Cookie{
		Name:     "private",
		Value:    token,
		Path:     "/",
		HttpOnly: true,
		SameSite: http.SameSiteLaxMode,
		Secure:   true,
		MaxAge:   seconds,
		Expires:  time.Now().Add(time.Duration(minutes) * time.Minute),
	}

	http.SetCookie(w, cookie)

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, userPrivateSetResponse{
		Success: true,
		Message: "OK",
	})
}
