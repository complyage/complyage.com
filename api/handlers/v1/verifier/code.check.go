package verifier

import (
	"encoding/json"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/encrypted"
	"github.com/complyage/base/identity"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| VerificationCardCodeRequest
//||------------------------------------------------------------------------------------------------||

type codeCheckRequest struct {
	Identifier string `json:"uuid"`
	Code       string `json:"code"`
}

type codeCheckResponse struct {
	Identifier string `json:"identifier"`
	Status     string `json:"status"`
	Type       string `json:"type"`
	Details    string `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Card Code Verification Attempt
//||------------------------------------------------------------------------------------------------||

func VerificationCodeCheck(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Code Check")

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	var updateRequest codeCheckRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, app.Err("API").Code("INVALID_JSON"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(updateRequest.Identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load
	//||------------------------------------------------------------------------------------------------||

	app.Log.Data("Verification :", verifyRecord.UUID, "Code:", verifyRecord.TwoFactor.Code)

	//||------------------------------------------------------------------------------------------------||
	//|| Check Code
	//||------------------------------------------------------------------------------------------------||

	err = verifyRecord.TwoFactorVerify(updateRequest.Code)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Identity
	//||------------------------------------------------------------------------------------------------||

	iden, err := identity.Load(account.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, app.Err("Identity").Code("IDENTITY_LOAD_FAILED"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Age for Credit Card
	//||------------------------------------------------------------------------------------------------||

	switch verifyRecord.Type {
	case types.DataTypeCRCD:
		iden.UpdateAge(types.DataTypeCRCD.String(), 18, verifyRecord.UUID)
		verifyRecord.IdentityUpdated = true
	case types.DataTypePHNE:
		iden.SetVerification(types.DataTypePHNE.String(), true, verifyRecord.Data.PHNE.Mask(), verifyRecord.UUID)
		verifyRecord.IdentityUpdated = true
	case types.DataTypeMAIL:
		iden.SetVerification(types.DataTypeMAIL.String(), true, verifyRecord.Data.PHNE.Mask(), verifyRecord.UUID)
		verifyRecord.IdentityUpdated = true
	case types.DataTypeADDR:
		iden.SetVerification(types.DataTypeADDR.String(), true, verifyRecord.Data.ADDR.Mask(), verifyRecord.UUID)
		verifyRecord.IdentityUpdated = true
	default:
		app.Log.Error("Unknown verification type for identity update: ", verifyRecord.Type.String())
		responses.Error(w, http.StatusInternalServerError, app.Err("Verify").Code("UNKNOWN_VERIFICATION_TYPE"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Encrypted
	//||------------------------------------------------------------------------------------------------||

	switch verifyRecord.Type {
	//||------------------------------------------------------------------------------------------------||
	//|| Credit Card
	//||------------------------------------------------------------------------------------------------||
	case types.DataTypeCRCD:
		err := encrypted.SaveCRCD(account.Public, verifyRecord.UUID, verifyRecord.Data.CRCD)
		if err != nil {
			app.Log.Error(err.Error())
			responses.Error(w, http.StatusInternalServerError, app.Err("API").Code("ENCRYPTED_FAILED"))
			return
		}
		verifyRecord.EncryptedSaved = true
	//||------------------------------------------------------------------------------------------------||
	//|| Phone
	//||------------------------------------------------------------------------------------------------||
	case types.DataTypePHNE:
		err := encrypted.SavePHNE(account.Public, verifyRecord.UUID, verifyRecord.Data.PHNE)
		if err != nil {
			app.Log.Error(err.Error())
			responses.Error(w, http.StatusInternalServerError, app.Err("API").Code("ENCRYPTED_FAILED"))
			return
		}
		verifyRecord.EncryptedSaved = true
	//||------------------------------------------------------------------------------------------------||
	//|| Email Address
	//||------------------------------------------------------------------------------------------------||
	case types.DataTypeMAIL:
		err := encrypted.SaveMAIL(account.Public, verifyRecord.UUID, verifyRecord.Data.MAIL)
		if err != nil {
			app.Log.Error(err.Error())
			responses.Error(w, http.StatusInternalServerError, app.Err("API").Code("ENCRYPTED_FAILED"))
			return
		}
		verifyRecord.EncryptedSaved = true
	//||------------------------------------------------------------------------------------------------||
	//|| Address
	//||------------------------------------------------------------------------------------------------||
	case types.DataTypeADDR:
		err := encrypted.SaveADDR(account.Public, verifyRecord.UUID, verifyRecord.Data.ADDR)
		if err != nil {
			app.Log.Error(err.Error())
			responses.Error(w, http.StatusInternalServerError, app.Err("API").Code("ENCRYPTED_FAILED"))
			return
		}
		verifyRecord.EncryptedSaved = true
	//||------------------------------------------------------------------------------------------------||
	//|| Fail
	//||------------------------------------------------------------------------------------------------||
	default:
		app.Log.Error("Unknown verification type for encrypted update: ", verifyRecord.Type.String())
		responses.Error(w, http.StatusInternalServerError, app.Err("Verify").Code("UNKNOWN_VERIFICATION_TYPE"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success - Update Status
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.UpdateStatusVerified("TWOFACTOR")

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, codeCheckResponse{
		Identifier: verifyRecord.UUID,
		Status:     verifyRecord.Status.String(),
		Type:       verifyRecord.Type.String(),
		Details:    "Verification code is valid",
	})
}
