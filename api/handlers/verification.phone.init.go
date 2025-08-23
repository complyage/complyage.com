package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/send"
	"api/verification"
	"base/helpers"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func PhoneVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

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

	//||------------------------------------------------------------------------------------------------||
	//|| Parse JSON
	//||------------------------------------------------------------------------------------------------||

	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Failed to read request body")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate verification tuple (amount + 4-digit code)
	//||------------------------------------------------------------------------------------------------||

	var req interfaces.VerificationPhoneInitialRequest
	if err := json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload"+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate verification code
	//||------------------------------------------------------------------------------------------------||

	code, err := helpers.Generate6DigitCode()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate code")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Phone Structure
	//||------------------------------------------------------------------------------------------------||

	phone := interfaces.PhoneNumber{
		CountryCode: req.CountryCode,
		Number:      req.Phone,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	verificationUUID, iErr := verification.CreateVerificationPHNE(session.ID, session.Public, phone, code)
	if iErr != nil {
		fmt.Println("Error creating phone verification:", iErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to create verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Send the verification SMS
	//||------------------------------------------------------------------------------------------------||

	bodyTxt, sendErr := send.SendVerifyText(phone.CountryCode+phone.Number, code)
	if sendErr != nil {
		fmt.Println("Error sending verification SMS:", sendErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to send verification SMS")
		return
	}
	fmt.Println("Verification SMS sent successfully:", bodyTxt)

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, interfaces.VerificationBasicInitialResponse{
		UUID: verificationUUID,
	})

}
