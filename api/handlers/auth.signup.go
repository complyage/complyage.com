package handlers

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
//|| Handler :: Processes the Signup Request
//||------------------------------------------------------------------------------------------------||

func SignupHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	email := r.FormValue("email")
	accountType := r.FormValue("type") // USER or VNDR

	//||------------------------------------------------------------------------------------------------||
	//|| Validate
	//||------------------------------------------------------------------------------------------------||

	if !helpers.IsValidEmail(email) {
		responses.Error(w, http.StatusBadRequest, "Invalid email")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Values
	//||------------------------------------------------------------------------------------------------||

	key, err := helpers.GenerateRandom()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not generate key")
		return
	}

	keyEncoded := base64.URLEncoding.EncodeToString(key) // URL-safe Base64

	code, err := helpers.Generate6DigitCode()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Could not generate code")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create record
	//||------------------------------------------------------------------------------------------------||

	record := interfaces.VerificationRecord{
		Code:     code,
		Key:      key,
		Email:    email,
		Type:     accountType,
		Attempts: 0,
		Created:  time.Now(),
		Expires:  time.Now().Add(15 * time.Minute),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serialize to Save
	//||------------------------------------------------------------------------------------------------||

	data, err := json.Marshal(record)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to serialize")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save to Redis with expiry
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(context.Background(), "verify::"+keyEncoded, data, 15*time.Minute).Err()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to save verification")
		return
	}

	fmt.Printf("✅ TwoFactor %s :: key=%s code=%s\n", accountType, keyEncoded, code)

	//||------------------------------------------------------------------------------------------------||
	//|| Store
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"token": keyEncoded,
		"type":  accountType,
	})
}
