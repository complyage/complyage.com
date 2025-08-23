package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/abstract"
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/models"
	"base/responses"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: IDVerifyStatusMediaHandler
//|| Endpoint: GET /api/verification/media?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func VerifyIDMediaFetch(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	which := r.URL.Query().Get("which")

	//||------------------------------------------------------------------------------------------------||
	//|| Check Which
	//||------------------------------------------------------------------------------------------------||

	if which != "front" && which != "back" && which != "selfie" {
		responses.Error(w, http.StatusBadRequest, "Invalid 'which' value")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Find Verification Record & Account
	//||------------------------------------------------------------------------------------------------||

	var verificationRecord models.Verification
	if err := db.DB.Where("verification_uuid = ?", identifier).First(&verificationRecord).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Verification record not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Account
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusNotFound, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session Match
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.IDAccount {
		responses.Error(w, http.StatusForbidden, "Session does not match verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make Key
	//||------------------------------------------------------------------------------------------------||

	bucket := os.Getenv("MINIO_PRIVATE_BUCKET")
	objectName := fmt.Sprintf("verifications/%s/%s", identifier, which)

	//||------------------------------------------------------------------------------------------------||
	//|| Download
	//||------------------------------------------------------------------------------------------------||

	data, contentType, err := helpers.MinioDownload(bucket, objectName)
	if err != nil || len(data) == 0 {
		fmt.Println("Media file not found media:", len(data))
		responses.Success(w, http.StatusOK, interfaces.VerificationMediaResponse{
			Exists: false,
			Blob:   "",
			Type:   which,
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch from Redis
	//||------------------------------------------------------------------------------------------------||

	encoded := base64.StdEncoding.EncodeToString(data)
	fmt.Println("Encoded Media Size:", len(encoded))
	fmt.Println("Mime Type:", contentType)

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, interfaces.VerificationMediaResponse{
		Exists: true,
		Blob:   encoded,
		Mime:   contentType,
		Type:   which,
	})
}
