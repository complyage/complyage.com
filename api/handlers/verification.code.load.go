package handlers

import (
	"base/abstract"
	"base/constants"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"base/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| VerificationCoadeLoadRequest
//||------------------------------------------------------------------------------------------------||

type VerificationCodeLoadResponse struct {
	UUID    string `json:"uuid"`
	Status  string `json:"status"`
	Type    string `json:"type"`
	Details string `json:"details"`
	Message string `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Card Code Verification Attempt
//||------------------------------------------------------------------------------------------------||

func VerificationCodeLoad(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch a Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	var verificationRecord models.Verification
	if err := db.DB.Where("verification_uuid = ?", identifier).First(&verificationRecord).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Verification record not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	if verificationRecord.FidAccount != session.ID {
		fmt.Println("Verification Record Account:", verificationRecord.FidAccount, "Session Account:", session.ID)
		responses.Error(w, http.StatusForbidden, "Verification record does not match session account. Please re-login")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Status Check - Verified
	//||------------------------------------------------------------------------------------------------||

	if verificationRecord.Status == constants.VerificationStatuses.Verified {
		responses.Error(w, http.StatusBadRequest, "Verification already completed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Status Check - Not Pending Verification
	//||------------------------------------------------------------------------------------------------||

	if verificationRecord.Status != constants.VerificationStatuses.PendingVerification {
		responses.Error(w, http.StatusBadRequest, "Verification not in pending state")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal & Check Secret and Meta
	//||------------------------------------------------------------------------------------------------||

	var secret interfaces.VerificationSecret
	if err := json.Unmarshal([]byte(verificationRecord.Secret), &secret); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to decode secret: "+err.Error())
		return
	}

	var dbMeta interfaces.VerificationMeta
	if err := json.Unmarshal([]byte(verificationRecord.Meta), &dbMeta); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to decode meta: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update Attempts and Expiration
	//||------------------------------------------------------------------------------------------------||

	now := time.Now().UTC()
	expirationTime, _ := helpers.FromUniversalDate(secret.Expiration)
	if !expirationTime.IsZero() && now.After(expirationTime) {
		responses.Error(w, http.StatusBadRequest, "Verification code expired")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success! Update Status, Identity, and Save
	//||------------------------------------------------------------------------------------------------||

	if secret.Attempts > 5 {
		responses.Error(w, http.StatusBadRequest, "Too many attempts")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":    identifier,
		"status":  verificationRecord.Status,
		"type":    verificationRecord.Type,
		"details": "Verification code is valid",
	})
}
