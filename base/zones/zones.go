package zones

import (
	"fmt"

	"base/db/abstract"
	"base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| In‑memory slice of zones (basic info)
//||------------------------------------------------------------------------------------------------||

var Zones []models.Zone

//||------------------------------------------------------------------------------------------------||
//|| LoadZones
//||------------------------------------------------------------------------------------------------||

func LoadZones() error {
	var results []models.Zone
	results, err := abstract.ReturnAllZones()
	if err != nil {
		return err
	}
	Zones = make([]models.Zone, len(results))
	copy(Zones, results)
	fmt.Printf("\n\033[32m[LOAD] - Loaded %d zones into memory\033[0m\n", len(Zones))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| FindZoneByID
//||------------------------------------------------------------------------------------------------||

func FindZoneByID(id uint) (*models.Zone, bool) {
	for i := range Zones {
		if Zones[i].IDZone == id {
			return &Zones[i], true
		}
	}
	return nil, false
}

//||------------------------------------------------------------------------------------------------||
//|| FindZoneByLocation
//||------------------------------------------------------------------------------------------------||

func FindZoneByLocation(state, country string) (*models.Zone, bool) {
	// unknown
	if state == "" || country == "" {
		for i := range Zones {
			if Zones[i].IDZone == 9999 {
				return &Zones[i], true
			}
		}
		return nil, false
	}

	// exact state match
	for i := range Zones {
		z := &Zones[i]
		if z.ZoneState != nil && *z.ZoneState == state && z.ZoneCountry != nil && *z.ZoneCountry == country {
			return z, true
		}
	}

	// fallback: country‑wide match (state null but country matches)
	for i := range Zones {
		z := &Zones[i]
		if z.ZoneState == nil && z.ZoneCountry != nil && *z.ZoneCountry == country {
			return z, true
		}
	}

	// no specific match, but caller treats zoneFound==true as “no enforcement needed”
	return nil, false
}
