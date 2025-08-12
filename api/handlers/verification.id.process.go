//||------------------------------------------------------------------------------------------------||
//|| GET /v1/api/verification/id/process
//||------------------------------------------------------------------------------------------------||

package handlers

import (
	"base/db"
	"base/interfaces"
	"base/responses"
	"context"
	"encoding/json"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification ID Process
//||------------------------------------------------------------------------------------------------||

func VerificationIDProcess(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Query Parameter
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")
	if identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing identifier")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load from Redis
	//||------------------------------------------------------------------------------------------------||

	key := "verify::iden::" + identifier
	raw, err := db.Redis.Get(context.Background(), key).Bytes()
	if err != nil {
		responses.Error(w, http.StatusNotFound, "Verification process not found")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load from Redis
	//||------------------------------------------------------------------------------------------------||

	var process interfaces.VerificationProcessID
	if err := json.Unmarshal(raw, &process); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to decode verification process")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, process)
}
