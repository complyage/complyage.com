package responses

import (
	"base/helpers"
	"base/models"
	"encoding/json"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Success Response
//||------------------------------------------------------------------------------------------------||

func Enforce(w http.ResponseWriter, site models.Site, zone models.Zone) {
	//||------------------------------------------------------------------------------------------------||
	//|| Convert to Optimized Structures
	//||------------------------------------------------------------------------------------------------||
	optimizedSite := helpers.OptimizeSite(site)
	optimizedZone := helpers.OptimizeZone(zone)
	//||------------------------------------------------------------------------------------------------||
	//|| Headers
	//||------------------------------------------------------------------------------------------------||
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//||------------------------------------------------------------------------------------------------||
	//|| Response Data
	//||------------------------------------------------------------------------------------------------||
	json.NewEncoder(w).Encode(map[string]any{
		"enforce": true,
		"site":    optimizedSite,
		"zone":    optimizedZone,
	})
}

//||------------------------------------------------------------------------------------------------||
//|| Error Response
//||------------------------------------------------------------------------------------------------||

func NoEnforce(w http.ResponseWriter) {
	//||------------------------------------------------------------------------------------------------||
	//|| Headers
	//||------------------------------------------------------------------------------------------------||
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//||------------------------------------------------------------------------------------------------||
	//|| Response Data
	//||------------------------------------------------------------------------------------------------||
	json.NewEncoder(w).Encode(map[string]any{
		"enforce": false,
	})
}
