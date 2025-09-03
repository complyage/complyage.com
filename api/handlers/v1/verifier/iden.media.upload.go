package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/abstract"
	"base/verify"
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
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

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session Match
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.ID {
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
	//|| Load Verification Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, account.Private, account.Public)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Media
	//||------------------------------------------------------------------------------------------------||

	var mediaRecord verify.Media
	if len(content) == 0 {
		mediaRecord = verify.Media{
			Exists: false,
			Size:   0,
			Mime:   "",
			Base64: "",
		}
	} else {
		mediaRecord = verify.Media{
			Exists: true,
			Size:   int64(len(content)),
			Mime:   mime,
			Base64: base64.StdEncoding.EncodeToString(content),
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Data
	//||------------------------------------------------------------------------------------------------||

	verifyData := verifyRecord.Encrypted.Data.IDEN

	//||------------------------------------------------------------------------------------------------||
	//|| Add the Correct Media
	//||------------------------------------------------------------------------------------------------||

	switch which {
	case "front":
		verifyData.Front = mediaRecord
	case "back":
		verifyData.Back = mediaRecord
	case "selfie":
		verifyData.Selfie = mediaRecord
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Reassign the Data
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.SetDataIDEN(verifyData)

	//||------------------------------------------------------------------------------------------------||
	//|| Save All
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Save()
	verifyRecord.DatabaseUpdate()

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("✅ Uploaded %s for verification: %s (size=%d)\n", which, identifier, len(content))
	responses.Success(w, http.StatusOK, mediaRecord)
}
