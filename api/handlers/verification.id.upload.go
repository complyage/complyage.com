package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/responses"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: UploadVerificationMedia
//|| Endpoint: POST /api/verification/upload?identifier=...&which=front|back|selfie
//||------------------------------------------------------------------------------------------------||

func UploadVerificationIDMedia(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	which := r.URL.Query().Get("which")

	if identifier == "" || which == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier or which param")
		return
	}

	if which != "front" && which != "back" && which != "selfie" {
		responses.Error(w, http.StatusBadRequest, "Invalid 'which' value")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Missing session")
		return
	}

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Read Existing Process from Redis
	//||------------------------------------------------------------------------------------------------||

	key := "verify::iden::" + identifier
	raw, err := db.Redis.Get(context.Background(), key).Bytes()
	if err != nil {
		responses.Error(w, http.StatusNotFound, "Verification process not found")
		return
	}

	var process interfaces.VerificationProcessID
	if err := json.Unmarshal(raw, &process); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to parse verification data")
		return
	}

	if process.AccountID != session.ID {
		responses.Error(w, http.StatusForbidden, "Unauthorized access")
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

	mime := header.Header.Get("Content-Type")
	if mime == "" {
		mime = http.DetectContentType(content)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Media Struct
	//||------------------------------------------------------------------------------------------------||

	media := map[string]any{
		"type":    "image", // Assume "image" — you can improve this
		"section": which,
		"blob":    base64.StdEncoding.EncodeToString(content),
		"mime":    mime,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Assign Media to Correct Field
	//||------------------------------------------------------------------------------------------------||

	switch which {
	case "front":
		process.Front = media
	case "back":
		process.Back = media
	case "selfie":
		process.Selfie = media
	default:
		responses.Error(w, http.StatusBadRequest, "Unknown media target")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save Updated Process Back to Redis
	//||------------------------------------------------------------------------------------------------||

	updated, err := json.Marshal(process)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to serialize updated record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Rewrite Redis
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(r.Context(), key, updated, 60*time.Minute).Err()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to store updated record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("✅ Uploaded %s for verification: %s (size=%d)\n", which, identifier, len(content))
	responses.Success(w, http.StatusOK, map[string]string{
		"status":  "uploaded",
		"section": which,
		"mime":    mime,
	})
}
