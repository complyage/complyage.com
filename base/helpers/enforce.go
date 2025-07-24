package helpers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/models"
	"fmt"
	"strconv"

	"gorm.io/datatypes"
)

//||------------------------------------------------------------------------------------------------||
//|| Enforce
//||------------------------------------------------------------------------------------------------||

func ShouldEnforce(city string, state string, site models.Site, zone models.Zone, zoneFound bool) bool {

	//||------------------------------------------------------------------------------------------------||
	//|| All Traffic
	//||------------------------------------------------------------------------------------------------||

	if site.SiteEnforcement == "ALLZ" {
		return true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Regulated Zones
	//||------------------------------------------------------------------------------------------------||

	if site.SiteEnforcement == "REGZ" && zoneFound {
		if zone.IDZone == 9999 {
			return false
		}
		return true
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Custom Zones
	//||------------------------------------------------------------------------------------------------||

	if site.SiteEnforcement == "CSTM" && zoneFound {
		// stringify the zone ID to look it up in the JSON map
		key := strconv.Itoa(int(zone.IDZone))

		// Zones is datatypes.JSONMap (map[string]interface{})
		var zonesMap datatypes.JSONMap
		if zm, ok := any(site.SiteZones).(datatypes.JSONMap); ok {
			zonesMap = zm
		} else {
			// parsing error or wrong type: enforce by default
			return true
		}

		// handle the “unknown” zone default
		if zone.IDZone == 9999 {
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
