package handlers

import (
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"net/http"
)

func AgentFatalHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	var req interfaces.AgentFatalRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload: "+err.Error())
		return
	}

	// Process the fatal error logic here...
}
