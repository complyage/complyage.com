package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: InitVerification
//|| Endpoint: /api/verification/init
//|| Description: Starts a new verification session for type "iden"
//||------------------------------------------------------------------------------------------------||

func VerificationInit(w http.ResponseWriter, r *http.Request) {

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
	//|| Parse Type from Query
	//||------------------------------------------------------------------------------------------------||

	vtype := strings.ToLower(r.URL.Query().Get("type"))
	if vtype != "iden" {
		responses.Error(w, http.StatusBadRequest, "Unsupported verification type")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Identity and Store in Redis
	//|| Key: verify:{identity} => TTL 60 minutes
	//||------------------------------------------------------------------------------------------------||

	identifier := uuid.NewString()

	record := interfaces.VerificationProcessID{
		Identifier: identifier,
		Status:     "pending",
		Level:      "basic",
		IDType:     "other",
		Error:      nil,
		AccountID:  session.ID,
		Front:      nil,
		Back:       nil,
		Selfie:     nil,
		Steps: interfaces.VerificationSteps{
			ParsedTextFront: false,
			ParsedTextBack:  false,
			DataParsed:      false,
			HasDOB:          false,
			HasName:         false,
			HasAddress:      false,
			FaceMatch:       false,
			Verified:        false,
		},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal Session Data
	//||------------------------------------------------------------------------------------------------||

	jsonBytes, err := json.Marshal(record)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to marshal verification record")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal Session Data
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(r.Context(), "verify::iden::"+identifier, jsonBytes, 60*time.Minute).Err()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to store session in Redis")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Identity to Frontend
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]string{
		"type":       vtype,
		"identifier": identifier,
	})
}
