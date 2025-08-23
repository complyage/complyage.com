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
	"io"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: UploadVerificationIDMedia
//|| Endpoint: POST /api/verification/upload?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func VerifyIDMediaUpload(w http.ResponseWriter, r *http.Request) {

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
	//|| Read Uploaded File
	//||------------------------------------------------------------------------------------------------||

	file, header, err := r.FormFile("media")
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Missing media file")
		return
	}
	defer file.Close()

	content, err := io.ReadAll(file)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to read file")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get MIME
	//||------------------------------------------------------------------------------------------------||

	mime := header.Header.Get("Content-Type")
	if mime == "" {
		mime = http.DetectContentType(content)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make Key
	//||------------------------------------------------------------------------------------------------||

	bucket := os.Getenv("MINIO_PRIVATE_BUCKET")
	objectName := fmt.Sprintf("verifications/%s/%s", identifier, which)

	//||------------------------------------------------------------------------------------------------||
	//|| Check File
	//||------------------------------------------------------------------------------------------------||

	if len(content) == 0 {
		if err := helpers.MinioDelete(bucket, objectName); err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to delete media")
			return
		}
		responses.Success(w, http.StatusOK, map[string]string{
			"status":  "deleted",
			"section": which,
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Upload the file to MinIO
	//||------------------------------------------------------------------------------------------------||

	url, err := helpers.MinioUpload(bucket, objectName, mime, content)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to upload to MinIO")
		return
	}
	fmt.Println("✅ Uploaded to MinIO:", url)

	//||------------------------------------------------------------------------------------------------||
	//|| If content is empty, delete from Redis and return status
	//||------------------------------------------------------------------------------------------------||

	if len(content) == 0 {
		if err := helpers.MinioDelete(bucket, objectName); err != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to delete media from MinIO")
			return
		}

		responses.Success(w, http.StatusOK, interfaces.VerificationMediaResponse{
			Exists: true,
			Blob:   "",
			Mime:   "",
			Type:   which,
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch from Redis
	//||------------------------------------------------------------------------------------------------||

	encoded := base64.StdEncoding.EncodeToString(content)

	//||------------------------------------------------------------------------------------------------||
	//|| Build Media Struct
	//||------------------------------------------------------------------------------------------------||

	media := interfaces.VerificationMediaResponse{
		Exists: true,
		Blob:   encoded,
		Mime:   mime,
		Type:   which,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("✅ Uploaded %s for verification: %s (size=%d)\n", which, identifier, len(content))
	responses.Success(w, http.StatusOK, media)
}
