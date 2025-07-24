package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"

	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler :: Processes the Signup Request
//||------------------------------------------------------------------------------------------------||

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Var
	//||------------------------------------------------------------------------------------------------||

	email := r.FormValue("email")
	password := r.FormValue("password")
	captcha := r.FormValue("captcha")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate
	//||------------------------------------------------------------------------------------------------||

	if !helpers.IsValidEmail(email) {
		responses.Error(w, http.StatusBadRequest, "Invalid email")
		return
	}

	if password == "" {
		responses.Error(w, http.StatusBadRequest, "Password is required")
		return
	}

	fmt.Println("[Login] Incoming -> email:", email, " password:", password, " captcha:", captcha)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Hashed Email
	//||------------------------------------------------------------------------------------------------||

	hashedEmail := helpers.GenerateEmailHash(email)

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Account Record
	//||------------------------------------------------------------------------------------------------||

	var account models.Account
	if err := db.DB.Where("account_email = ?", hashedEmail).First(&account).Error; err != nil {
		account.AccountEmail = "INVALID-EMAIL~!PREVENT-TIME~COM"
		account.AccountPassword = "DUMMY-PASSWORD-TEST"
		account.AccountSalt = "DUMMY-PREVENT-TIMING-DUMMY-PREVENT-TIMING-DUMMY-PREVENT-TIMING"
		account.AccountLevel = helpers.Int8Ptr(0)
	}

	fmt.Println("Account Level : ", account.AccountLevel)
	fmt.Println("Account Email : ", account.AccountEmail)
	fmt.Println("Account Salt : ", account.AccountSalt)

	//||------------------------------------------------------------------------------------------------||
	//|| Check the Password
	//||------------------------------------------------------------------------------------------------||

	if !helpers.VerifyPassword(password, account.AccountSalt, account.AccountPassword) {
		responses.Error(w, http.StatusUnauthorized, "Invalid email or password")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Account Record and redirect if completed
	//||------------------------------------------------------------------------------------------------||

	loginToken, err := helpers.SessionCreate(email, account)
	if err == nil {
		helpers.WriteSessionCookie(w, loginToken)
		responses.Success(w, http.StatusOK, map[string]any{
			"message": "Login Successful",
			"next":    "/members/",
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Store
	//||------------------------------------------------------------------------------------------------||

	responses.Error(w, http.StatusUnauthorized, "Error Writing Session Cookie, please try again later")

}
