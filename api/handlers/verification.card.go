package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/verification"
	"base/adapters"
	"base/helpers"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func CCVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

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

	var req interfaces.VerificationCardInitialRequest
	if err := json.Unmarshal(body, &req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload"+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate verification tuple (amount + 4-digit code)
	//||------------------------------------------------------------------------------------------------||

	code, err := helpers.Generate6DigitCode()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate code")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	verificationUUID, iErr := verification.CreateVerificationCRCD(session.ID, session.Public, code)
	if iErr != nil {
		fmt.Println("Error creating card verification:", iErr)
		responses.Error(w, http.StatusInternalServerError, "Failed to create verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification
	//||------------------------------------------------------------------------------------------------||

	intent, err := adapters.StripeIntent(verificationUUID, req.Amount, req.Currency, code)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create payment intent")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, interfaces.VerificationCardInitialResponse{
		UUID:         verificationUUID,
		Amount:       intent.Amount,
		Currency:     string(intent.Currency),
		ClientSecret: intent.ClientSecret,
	})

}

//||------------------------------------------------------------------------------------------------||
//|| (Optional) Trim utility for local use
//||------------------------------------------------------------------------------------------------||

func trim(s string) string { return strings.TrimSpace(s) }
