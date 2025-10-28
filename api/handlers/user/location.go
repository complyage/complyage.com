package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/complyage/base/ips"

	"github.com/ralphferrara/aria/auth/actions"
	ariaHTTP "github.com/ralphferrara/aria/http"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Envelope
//||------------------------------------------------------------------------------------------------||

type apiEnvelope struct {
	Success bool                               `json:"success"`
	Data    ips.IPLocationVerificationResponse `json:"data"`
	Message string                             `json:"message,omitempty"`
	Error   string                             `json:"error,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerifications - Location
//||------------------------------------------------------------------------------------------------||

func LocationUserDashboard(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	ipAddress := ariaHTTP.GetClientIP(r)
	if ipAddress == "127.0.0.1" || ipAddress == "::1" {
		ipAddress = "192.110.165.116"
	}
	fmt.Println("[Location] ip:", ipAddress)

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	_, _, _, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Internal URL
	//||------------------------------------------------------------------------------------------------||

	sharedKey := os.Getenv("API_SHARED_KEY")
	baseURL := os.Getenv("VITE_COMPLYAGE_CLIENT_URL") // ⚠️ if this points to your FRONTEND, consider using API base instead
	locationURL := fmt.Sprintf("%s/v1/internal/location?ip=%s&internal=%s", baseURL, ipAddress, sharedKey)
	fmt.Println("[Location] URL:", locationURL)

	//||------------------------------------------------------------------------------------------------||
	//|| Request
	//||------------------------------------------------------------------------------------------------||
	req, err := http.NewRequestWithContext(r.Context(), "GET", locationURL, nil)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create request")
		return
	}
	req.Header.Set("Accept", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch location")
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		responses.Error(w, resp.StatusCode, "Location service error")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decode (NOTE: response is wrapped in {success,data})
	//||------------------------------------------------------------------------------------------------||

	var env apiEnvelope
	if err := json.NewDecoder(resp.Body).Decode(&env); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to decode location response")
		return
	}

	if !env.Success {
		msg := env.Message
		if msg == "" {
			msg = "Location lookup failed"
		}
		responses.Error(w, http.StatusBadGateway, msg)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Debug
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("[LocationResp] %+v\n", env.Data)

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response (pass through)
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, ips.IPLocationVerificationResponse{
		IPAddress: env.Data.IPAddress,
		City:      env.Data.City,
		Region:    env.Data.Region,
		Country:   env.Data.Country,
		Types:     env.Data.Types,
		MinAge:    env.Data.MinAge,
	})
}
