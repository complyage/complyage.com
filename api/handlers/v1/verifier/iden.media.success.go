package verifier

import (
	"api/agent"
	"api/verification"
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
	"os"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDSuccessHandler
//||------------------------------------------------------------------------------------------------||

func VerifyIDSuccessHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var updateRequest interfaces.VerificationIDSuccessRequest
	if err := json.NewDecoder(r.Body).Decode(&updateRequest); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update Request
	//||------------------------------------------------------------------------------------------------||

	if updateRequest.UUID == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(updateRequest.UUID)
	if err != nil || account == nil {
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	var verificationRecord models.Verification
	if err := db.DB.Where("verification_uuid = ?", updateRequest.UUID).First(&verificationRecord).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Verification record not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate All Three Images Exist in MinIO and Are Valid Images
	//||------------------------------------------------------------------------------------------------||

	bucket := os.Getenv("MINIO_PRIVATE_BUCKET")
	images := []string{"front", "back", "selfie"}
	for _, which := range images {
		objectName := fmt.Sprintf("verifications/%s/%s", updateRequest.UUID, which)
		data, contentType, err := helpers.MinioDownload(bucket, objectName)

		if err != nil || len(data) == 0 {
			fmt.Println("Object : ", objectName, "not found or empty")
			fmt.Println("Error downloading image:", err)
			fmt.Println("Media Size:", len(data))
			responses.Error(w, http.StatusBadRequest, fmt.Sprintf("Missing or invalid file: %s", which))
			return
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Validate it's an image (contentType should start with "image/")
		//||------------------------------------------------------------------------------------------------||
		if !strings.HasPrefix(contentType, "image/") {
			responses.Error(w, http.StatusBadRequest, fmt.Sprintf("File is not an image: %s (%s)", which, contentType))
			return
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	insert := agent.AgentVerifyIDStart(updateRequest.UUID)
	if insert != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	updErr := verification.UpdateStatusIDENPendingVerification(updateRequest.UUID)
	if updErr != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to update verification: "+updErr.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   updateRequest.UUID,
		"status": constants.VerificationStatuses.PendingVerification,
	})
}
