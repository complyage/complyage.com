package handlers

import (
	"fmt"
	"net/http"
	"os"

	"base/helpers"
	"base/loaders"
	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| In Memory
//||------------------------------------------------------------------------------------------------||

var UseInMemory = os.Getenv("ENV_MODE") == "production"

//||------------------------------------------------------------------------------------------------||
//|| Handle Loading of IP ranges into memory
//||------------------------------------------------------------------------------------------------||

func CheckClientEnforcement(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||
	apiKey := r.URL.Query().Get("apiKey")
	if apiKey == "" {
		responses.Error(w, http.StatusUnauthorized, "API key is required")
		return
	}
	ipAddress := helpers.GetClientIP(r)
	//||------------------------------------------------------------------------------------------------||
	//|| Get the API Key Data
	//||------------------------------------------------------------------------------------------------||
	site := loaders.GetSiteByPublic(apiKey)
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	location, err := helpers.GetLocationByIP(ipAddress)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get location by IP: "+err.Error())
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Get Zone
	//||------------------------------------------------------------------------------------------------||
	zone, zoneFound := loaders.FindZoneByLocation(location.City, location.State)
	//||------------------------------------------------------------------------------------------------||
	//|| Check if in a restricted territory
	//||------------------------------------------------------------------------------------------------||
	shouldEnforce := helpers.ShouldEnforce(location.Country, location.State, *site, *zone, zoneFound)
	fmt.Println(site.SiteURL)
	fmt.Println(zone.IDZone)
	//||------------------------------------------------------------------------------------------------||
	//|| Not allowed
	//||------------------------------------------------------------------------------------------------||
	if !shouldEnforce {
		responses.NoEnforce(w)
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Give the data
	//||------------------------------------------------------------------------------------------------||
	responses.Enforce(w, *site, *zone)
}
