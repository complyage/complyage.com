package ips

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"strconv"

	"github.com/complyage/base/db/models"

	"gorm.io/datatypes"
)

//||------------------------------------------------------------------------------------------------||
//|| Enforce
//||------------------------------------------------------------------------------------------------||

func ShouldEnforce(city string, state string, site models.Site, zone models.Zone, zoneFound bool) bool {

	//||------------------------------------------------------------------------------------------------||
	//|| All Traffic
	//||------------------------------------------------------------------------------------------------||

	if site.Enforcement == "ALLZ" {
		return true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Regulated Zones
	//||------------------------------------------------------------------------------------------------||

	if site.Enforcement == "REGZ" && zoneFound {
		if zone.ID == 9999 {
			return false
		}
		return true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Custom Zones
	//||------------------------------------------------------------------------------------------------||

	if site.Enforcement == "CSTM" && zoneFound {
		// stringify the zone ID to look it up in the JSON map
		key := strconv.Itoa(int(zone.ID))

		// Zones is datatypes.JSONMap (map[string]interface{})
		var zonesMap datatypes.JSONMap
		if zm, ok := any(site.Zones).(datatypes.JSONMap); ok {
			zonesMap = zm
		} else {
			// parsing error or wrong type: enforce by default
			return true
		}

		// handle the “unknown” zone default
		if zone.ID == 9999 {
			if raw, exists := zonesMap[key]; exists {
				// raw will typically be float64 (JSON numbers) or string
				val := fmt.Sprintf("%v", raw)
				return val != "0"
			}
			// not defined → enforce by default
			return true
		}

		// for other zones: if key exists and value == "0", skip enforcement
		if raw, exists := zonesMap[key]; exists {
			val := fmt.Sprintf("%v", raw)
			return val != "0"
		}
		// not defined → enforce by default
		return true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Broken
	//||------------------------------------------------------------------------------------------------||

	return false

}
