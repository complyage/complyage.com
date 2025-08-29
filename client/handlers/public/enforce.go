package public

import (
	"fmt"
	"net/http"
	"os"

	"base/ips"
	"base/sites"
	"base/zones"

	ariaHTTP "github.com/ralphferrara/aria/http"
	"github.com/ralphferrara/aria/responses"
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
	lang := r.URL.Query().Get("lang")
	ipAddress := ariaHTTP.GetClientIP(r)
	//||------------------------------------------------------------------------------------------------||
	//|| Get the API Key Data
	//||------------------------------------------------------------------------------------------------||
	if apiKey == "" {
		responses.Error(w, http.StatusUnauthorized, "API key is required")
		return
	}
	if lang == "" {
		lang = "en"
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Get the API Key Data
	//||------------------------------------------------------------------------------------------------||
	site := sites.FetchSiteByPublic(apiKey)
	if site == nil {
		// 401 because we didn’t find a matching API key
		responses.Error(w, http.StatusUnauthorized, "invalid apiKey")
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	location, err := ips.GetLocationByIP(ipAddress)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get location by IP: "+err.Error())
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Get Zone
	//||------------------------------------------------------------------------------------------------||
	zone, zoneFound := zones.FindZoneByLocation(location.City, location.State)
	if !zoneFound || zone == nil {
		// no enforcement zone for this location → no enforcement
		ips.NoEnforce(w)
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Check if in a restricted territory
	//||------------------------------------------------------------------------------------------------||
	shouldEnforce := ips.ShouldEnforce(location.Country, location.State, *site, *zone, zoneFound)
	fmt.Println(site.SiteURL)
	fmt.Println(zone.IDZone)
	//||------------------------------------------------------------------------------------------------||
	//|| Not allowed
	//||------------------------------------------------------------------------------------------------||
	if !shouldEnforce {
		ips.NoEnforce(w)
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Give the data
	//||------------------------------------------------------------------------------------------------||
	ips.Enforce(w, *site, *zone)
}
