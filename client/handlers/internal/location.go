package internal

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/ips"
	"base/zones"
	"net/http"
	"os"
	"strings"

	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Types
//||------------------------------------------------------------------------------------------------||

type locationResponse struct {
	IPAddress string `json:"ipAddress"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Types     string `json:"types"`
}

//||------------------------------------------------------------------------------------------------||
//|| HandleInternalLocation
//||------------------------------------------------------------------------------------------------||

func HandleInternalLocation(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Internal Check
	//||------------------------------------------------------------------------------------------------||

	sharedKey := strings.TrimSpace(r.URL.Query().Get("internal"))
	checkKey := os.Getenv("API_SHARED_KEY")
	if sharedKey == "" || checkKey == "" || sharedKey != checkKey {
		responses.Error(w, http.StatusUnauthorized, "Unauthorized")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Resolve IP
	//||------------------------------------------------------------------------------------------------||

	ipAddress := strings.TrimSpace(r.URL.Query().Get("ip"))
	if ipAddress == "" {
		responses.Error(w, http.StatusBadRequest, "Missing IP address")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Geo Lookup
	//||------------------------------------------------------------------------------------------------||

	loc, err := ips.GetLocationByIP(ipAddress)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get location")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Zone Match (no early-return; Types may be empty)
	//||------------------------------------------------------------------------------------------------||

	var zoneTypes string
	var zoneMinAge int
	if zone, ok := zones.FindZoneByLocation(loc.State, loc.Country); ok && zone.ZoneRequirements != nil {
		zoneTypes = strings.TrimSpace(*zone.ZoneRequirements)
		zoneMinAge = zone.ZoneMinAge
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Respond
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, ips.IPLocationVerificationResponse{
		IPAddress: ipAddress,
		City:      loc.City,
		Region:    loc.State,
		Country:   loc.Country,
		Types:     zoneTypes,
		MinAge:    zoneMinAge,
	})
}
