package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/verification"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"base/responses"
	"encoding/json"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDProgressHandler
//||------------------------------------------------------------------------------------------------||

func VerificationUpdate(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var req verification.VerificationIDProgressRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate UUID
	//||------------------------------------------------------------------------------------------------||

	if req.UUID == "" {
		responses.Error(w, http.StatusBadRequest, "Missing verification UUID")
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Step
	//||------------------------------------------------------------------------------------------------||

	if req.StepName == "" || req.StepStatus == "" {
		responses.Error(w, http.StatusBadRequest, "Missing stepName or stepStatus")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record & Account
	//||------------------------------------------------------------------------------------------------||

	updatedRecord := models.Verification{}
	verificationRecord, err := verification.LoadVerificationRecordByUUID(req.UUID)
	if err != nil || verificationRecord == nil {
		responses.Error(w, http.StatusNotFound, "Verification record not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build & Append New Step
	//||------------------------------------------------------------------------------------------------||

	newStep := interfaces.VerificationMetaStep{
		StepName:      req.StepName,
		StepStatus:    req.StepStatus,
		StepDetails:   req.StepDetails,
		StepTimestamp: helpers.UniversalNow(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Append the new step to existing steps
	//||------------------------------------------------------------------------------------------------||

	appendedSteps := append(verificationRecord.Meta.Steps, newStep)
	verificationRecord.Meta.Steps = appendedSteps

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	jsonMeta, err := json.Marshal(verificationRecord.Meta)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to marshal meta: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| What Status to Update?
	//||------------------------------------------------------------------------------------------------||

	verificationRecord.Meta.Step = verificationRecord.Meta.Step + 1
	updatedRecord.Status = verificationRecord.Status
	updatedRecord.Meta = string(jsonMeta)

	//||------------------------------------------------------------------------------------------------||
	//|| DB Update
	//||------------------------------------------------------------------------------------------------||

	dbErr := db.DB.Model(&models.Verification{}).Where("verification_uuid = ?", verificationRecord.UUID).Updates(updatedRecord).Error
	if dbErr != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to update verification record: "+dbErr.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   req.UUID,
		"status": req.StepStatus,
		"step":   req.Step,
	})
}
