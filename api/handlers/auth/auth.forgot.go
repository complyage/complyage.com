package auth

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler :: Initiates Forgot Password Flow
//||------------------------------------------------------------------------------------------------||

func ForgotPasswordHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	email := r.FormValue("email")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Email
	//||------------------------------------------------------------------------------------------------||

	if !helpers.IsValidEmail(email) {
		responses.Error(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Reset Token & Code
	//||------------------------------------------------------------------------------------------------||

	key, err := helpers.GenerateRandom()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not generate reset token")
		return
	}

	keyEncoded := base64.URLEncoding.EncodeToString(key)

	code, err := helpers.Generate6DigitCode()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not generate reset code")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Forgot Password Record
	//||------------------------------------------------------------------------------------------------||

	record := interfaces.AuthVerification{
		Code:     code,
		Key:      key,
		Email:    email,
		Type:     "RESET", // distinguish from signup
		Attempts: 0,
		Created:  time.Now(),
		Expires:  time.Now().Add(15 * time.Minute),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serialize Record for Storage
	//||------------------------------------------------------------------------------------------------||

	data, err := json.Marshal(record)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to serialize reset request")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save Reset Request in Redis (15 min expiry)
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(context.Background(), "reset::"+keyEncoded, data, 15*time.Minute).Err()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to store reset request")
		return
	}

	fmt.Printf("✅ ForgotPassword :: email=%s key=%s code=%s\n", email, keyEncoded, code)

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"token": keyEncoded,
		"email": email,
	})
}
