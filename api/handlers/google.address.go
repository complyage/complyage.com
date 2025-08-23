package handlers

import (
	"base/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Address Verification Request/Response
//||------------------------------------------------------------------------------------------------||

type AddressVerificationRequest struct {
	Line1   string `json:"line1"`
	Line2   string `json:"line2"`
	City    string `json:"city"`
	State   string `json:"state"`
	Postal  string `json:"postal"`
	Country string `json:"country"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Google Address Verification
//||------------------------------------------------------------------------------------------------||

func GoogleAddressVerifyHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Request
	//||------------------------------------------------------------------------------------------------||

	var req AddressVerificationRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON payload: "+err.Error(), http.StatusBadRequest)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Address
	//||------------------------------------------------------------------------------------------------||

	addressParts := []string{req.Line1}
	if strings.TrimSpace(req.Line2) != "" {
		addressParts = append(addressParts, req.Line2)
	}
	addressParts = append(addressParts, req.City, req.State, req.Postal, req.Country)
	fullAddress := strings.Join(addressParts, ", ")

	//||------------------------------------------------------------------------------------------------||
	//|| API Key
	//||------------------------------------------------------------------------------------------------||

	googleKey := os.Getenv("GOOGLE_API_PRIVATE")
	if googleKey == "" {
		http.Error(w, "Google API key not set on server", http.StatusInternalServerError)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Google
	//||------------------------------------------------------------------------------------------------||

	googleURL := fmt.Sprintf(
		"https://maps.googleapis.com/maps/api/geocode/json?address=%s&key=%s",
		url.QueryEscape(fullAddress),
		googleKey,
	)

	//||------------------------------------------------------------------------------------------------||
	//|| HTTP
	//||------------------------------------------------------------------------------------------------||

	client := &http.Client{Timeout: 7 * time.Second}
	resp, err := client.Get(googleURL)
	if err != nil {
		http.Error(w, "Failed to reach Google Maps API: "+err.Error(), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	//||------------------------------------------------------------------------------------------------||
	//|| Process Response
	//||------------------------------------------------------------------------------------------------||

	var apiResponse map[string]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&apiResponse); err != nil {
		http.Error(w, "Failed to decode Google response: "+err.Error(), http.StatusInternalServerError)
		return
	}

	status := fmt.Sprintf("%v", apiResponse["status"])
	if status != "OK" {
		msg := fmt.Sprintf("Google could not verify this address (%s)", status)
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Process Response
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "application/json")
	responses.Success(w, http.StatusOK, map[string]interface{}{
		"status":  status,
		"results": apiResponse["results"],
	})
}
