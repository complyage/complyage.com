package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: UploadVerificationIDMedia
//|| Endpoint: POST /api/verification/upload?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func VerifyFaceMediaUpload(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Face Media Upload")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

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
	//|| Get Verify Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
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
	//|| Create Media
	//||------------------------------------------------------------------------------------------------||

	var mediaRecord types.Media
	if len(content) == 0 {
		mediaRecord = types.Media{
			Exists: false,
			Size:   0,
			Mime:   "",
			Base64: "",
		}
	} else {
		mediaRecord = types.Media{
			Exists: true,
			Size:   int64(len(content)),
			Mime:   mime,
			Base64: base64.StdEncoding.EncodeToString(content),
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Data
	//||------------------------------------------------------------------------------------------------||

	verifyData := verifyRecord.Data.FACE
	verifyData.Selfie = mediaRecord

	//||------------------------------------------------------------------------------------------------||
	//|| Reassign the Data
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.SetDataFACE(verifyData)

	//||------------------------------------------------------------------------------------------------||
	//|| Save All
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Save()
	verifyRecord.DatabaseUpdate()

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("✅ Uploaded selfie for verification: %s (size=%d)\n", identifier, len(content))
	responses.Success(w, http.StatusOK, mediaRecord)
}
