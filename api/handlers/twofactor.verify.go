package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/constants"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"base/responses"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/redis/go-redis/v9"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

func TwoFactorHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	ctx := context.Background()
	token := r.FormValue("token")
	code := r.FormValue("code")

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("[TwoFactor] Incoming -> token:", token, " code:", code)

	//||------------------------------------------------------------------------------------------------||
	//|| Basic Validation
	//||------------------------------------------------------------------------------------------------||

	if token == "" || code == "" {
		responses.Error(w, http.StatusBadRequest, "Missing or invalid code/token/type")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Record from Redis
	//||------------------------------------------------------------------------------------------------||

	val, err := db.Redis.Get(ctx, "verify::"+token).Result()
	if err != nil {
		if err == redis.Nil {
			responses.Error(w, http.StatusBadRequest, "Invalid or expired token")
			return
		}
		responses.Error(w, http.StatusInternalServerError, "Redis error")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Convert to struct
	//||------------------------------------------------------------------------------------------------||

	var record interfaces.VerificationRecord
	if err := json.Unmarshal([]byte(val), &record); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Invalid stored record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Record from Redis
	//||------------------------------------------------------------------------------------------------||

	if record.Attempts >= 5 {
		db.Redis.Del(ctx, fmt.Sprintf("verify:%s", token))
		responses.Error(w, http.StatusTooManyRequests, "Too many attempts")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Record from Redis
	//||------------------------------------------------------------------------------------------------||

	if code != record.Code {
		record.Attempts++
		newData, _ := json.Marshal(record)
		db.Redis.Set(ctx, fmt.Sprintf("verify:%s", token), newData, time.Until(record.Expires))
		responses.Error(w, http.StatusUnauthorized, "Invalid code")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check if the token is expired
	//||------------------------------------------------------------------------------------------------||

	if time.Now().After(record.Expires) {
		db.Redis.Del(ctx, fmt.Sprintf("verify:%s", token))
		responses.Error(w, http.StatusBadRequest, "Token expired")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success! Delete
	//||------------------------------------------------------------------------------------------------||

	db.Redis.Del(ctx, fmt.Sprintf("verify:%s", token))

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Hashed Email
	//||------------------------------------------------------------------------------------------------||

	hashedEmail := helpers.GenerateEmailHash(record.Email)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Account Record
	//||------------------------------------------------------------------------------------------------||

	var account models.Account
	if err := db.DB.Where("account_email = ?", hashedEmail).First(&account).Error; err != nil {
		fmt.Println("[Session] Account not found for email:", record.Email, " creating new account")
	} else {
		existsToken, err := helpers.SessionCreate(record.Email, account)
		if err == nil {
			helpers.WriteSessionCookie(w, existsToken)
			responses.Success(w, http.StatusOK, map[string]any{
				"message": "Two-factor authentication successful. Redirecting to /members",
				"next":    "/members/",
			})
			return
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Struct
	//||------------------------------------------------------------------------------------------------||

	account.AccountType = record.Type
	account.AccountEmail = hashedEmail
	account.AccountStatus = constants.AccountStatus.Verified
	account.AccountLevel = helpers.Int8Ptr(1)

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Record
	//||------------------------------------------------------------------------------------------------||

	if err := db.DB.Create(&account).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Record
	//||------------------------------------------------------------------------------------------------||

	newToken, err := helpers.SessionCreate(record.Email, account)
	if err == nil {
		helpers.WriteSessionCookie(w, newToken)
		responses.Success(w, http.StatusOK, map[string]any{
			"message": "Account Created",
			"next":    "/complete",
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Account Record
	//||------------------------------------------------------------------------------------------------||

	responses.Error(w, http.StatusInternalServerError, "Failed to create account - Unknown")
}
