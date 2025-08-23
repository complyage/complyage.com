package handlers

// import (
// 	"api/verification"
// 	"base/abstract"
// 	"base/constants"
// 	"base/db"
// 	"base/helpers"
// 	"base/interfaces"
// 	"base/models"
// 	"base/responses"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"time"
// )

// //||------------------------------------------------------------------------------------------------||
// //|| VerificationCardCodeRequest
// //||------------------------------------------------------------------------------------------------||

// type VerificationCodeCheckRequest struct {
// 	UUID string `json:"uuid"`
// }

// type VerificationCodeCheckResponse struct {
// 	UUID    string `json:"uuid"`
// 	Status  string `json:"status"`
// 	Message string `json:"message"`
// }

// //||------------------------------------------------------------------------------------------------||
// //|| Handler: Card Code Verification Attempt
// //||------------------------------------------------------------------------------------------------||

// func VerificationCodeCheck(w http.ResponseWriter, r *http.Request) {

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Parse Request
// 	//||------------------------------------------------------------------------------------------------||

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Get Params
// 	//||------------------------------------------------------------------------------------------------||

// 	uuid := r.URL.Query().Get("identifier")
// 	identifier := r.URL.Query().Get("identifier")

// 	fmt.Println("[CCVerifyCheck] Incoming -> UUID:", req.UUID, " Code:", req.Code)
// 	if req.Code == "" {
// 		responses.Error(w, http.StatusBadRequest, "Missing code")
// 		return
// 	}

// 	if req.UUID == "" {
// 		responses.Error(w, http.StatusBadRequest, "Missing identifier")
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Fetch Verification Record (with account public for decrypt)
// 	//||------------------------------------------------------------------------------------------------||

// 	account, err := abstract.GetAccountByVerificationUUID(req.UUID)
// 	if err != nil || account == nil {
// 		responses.Error(w, http.StatusNotFound, "Account not found for verification")
// 		return
// 	}

// 	var verificationRecord models.Verification
// 	if err := db.DB.Where("verification_uuid = ?", req.UUID).First(&verificationRecord).Error; err != nil {
// 		responses.Error(w, http.StatusNotFound, "Verification record not found")
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Unmarshal & Check Secret and Meta
// 	//||------------------------------------------------------------------------------------------------||

// 	var secret interfaces.VerificationSecret
// 	if err := json.Unmarshal([]byte(verificationRecord.Secret), &secret); err != nil {
// 		responses.Error(w, http.StatusInternalServerError, "Failed to decode secret: "+err.Error())
// 		return
// 	}
// 	var dbMeta interfaces.VerificationMeta
// 	if err := json.Unmarshal([]byte(verificationRecord.Meta), &dbMeta); err != nil {
// 		responses.Error(w, http.StatusInternalServerError, "Failed to decode meta: "+err.Error())
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Update Attempts and Expiration
// 	//||------------------------------------------------------------------------------------------------||

// 	now := time.Now().UTC()
// 	expirationTime, _ := helpers.FromUniversalDate(secret.Expiration)
// 	if !expirationTime.IsZero() && now.After(expirationTime) {
// 		// Too late!
// 		secret.Attempts++

// 		meta := interfaces.VerificationMeta{
// 			Steps: []interfaces.VerificationMetaStep{
// 				{
// 					StepName:      "ATTEMPT",
// 					StepStatus:    "EXPIRED",
// 					StepDetails:   "Verification code expired",
// 					StepTimestamp: now.Format(time.RFC3339),
// 				},
// 			},
// 			Approval: dbMeta.Approval,
// 		}

// 		saveErr := verification.UpdateVerification(
// 			account.IDAccount, account.AccountPublic, req.UUID, verificationRecord.Display,
// 			interfaces.VerificationData{}, meta, secret, constants.VerificationStatuses.Rejected,
// 		)
// 		if saveErr != nil {
// 			responses.Error(w, http.StatusInternalServerError, "Failed to update attempts: "+saveErr.Error())
// 			return
// 		}
// 		responses.Error(w, http.StatusBadRequest, "Verification code expired")
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Check the code
// 	//||------------------------------------------------------------------------------------------------||

// 	if req.Code != secret.CheckCode {
// 		secret.Attempts++

// 		meta := interfaces.VerificationMeta{
// 			Steps: []interfaces.VerificationMetaStep{
// 				{
// 					StepName:      "ATTEMPT",
// 					StepStatus:    "FAILED",
// 					StepDetails:   "Incorrect code",
// 					StepTimestamp: now.Format(time.RFC3339),
// 				},
// 			},
// 			Approval: dbMeta.Approval,
// 		}

// 		saveErr := verification.UpdateVerification(
// 			account.IDAccount, account.AccountPublic, req.UUID, verificationRecord.Display,
// 			interfaces.VerificationData{}, meta, secret, constants.VerificationStatuses.Pending,
// 		)
// 		if saveErr != nil {
// 			responses.Error(w, http.StatusInternalServerError, "Failed to update attempts: "+saveErr.Error())
// 			return
// 		}
// 		responses.Error(w, http.StatusBadRequest, "Incorrect code")
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Success! Update Status, Identity, and Save
// 	//||------------------------------------------------------------------------------------------------||

// 	secret.Attempts++

// 	meta := interfaces.VerificationMeta{
// 		Steps: []interfaces.VerificationMetaStep{
// 			{
// 				StepName:      "ATTEMPT",
// 				StepStatus:    "VERIFIED",
// 				StepDetails:   "Code verified successfully",
// 				StepTimestamp: now.Format(time.RFC3339),
// 			},
// 		},
// 		Approval: dbMeta.Approval,
// 	}

// 	status := constants.VerificationStatuses.Verified

// 	saveErr := verification.UpdateVerification(
// 		account.IDAccount, account.AccountPublic, req.UUID, verificationRecord.Display,
// 		interfaces.VerificationData{}, meta, secret, status,
// 	)
// 	if saveErr != nil {
// 		responses.Error(w, http.StatusInternalServerError, "Failed to update verification: "+saveErr.Error())
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Update Identity
// 	//||------------------------------------------------------------------------------------------------||

// 	idErr := verification.IdentityUpdateCreditCard(account.IDAccount, verificationRecord.Display)
// 	if idErr != nil {
// 		responses.Error(w, http.StatusInternalServerError, "Failed to update identity: "+idErr.Error())
// 		return
// 	}

// 	//||------------------------------------------------------------------------------------------------||
// 	//|| Success
// 	//||------------------------------------------------------------------------------------------------||

// 	responses.Success(w, http.StatusOK, map[string]interface{}{
// 		"uuid":   req.UUID,
// 		"status": status,
// 	})
// }
