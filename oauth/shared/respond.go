package shared

import (
	"encoding/json"
	"net/http"
)

func (oas OAuthSharedAccess) Respond(r *http.Request, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(oas.Shared); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	return
}
