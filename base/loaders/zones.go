package loaders

import (
	"fmt"

	"base/db"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| In‑memory slice of zones (basic info)
//||------------------------------------------------------------------------------------------------||

var Zones []models.Zone

//||------------------------------------------------------------------------------------------------||
//|| LoadZones
//||
//|| Queries the `zones` table for all rows (id, state, country, requirements),
//|| orders by `id_zone ASC`, and stores them in the in‑memory `Zones` slice.
//|| Returns an error if the DB query fails.
//||------------------------------------------------------------------------------------------------||

func LoadZones() error {
	var results []models.Zone
	if err := db.DB.
		Table("zones").
		Order("id_zone ASC").
		Find(&results).Error; err != nil {

		return fmt.Errorf("failed to load zones: %w", err)
	}

	Zones = make([]models.Zone, len(results))
	copy(Zones, results)

	fmt.Printf("Loaded %d zones into memory\n", len(Zones))
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
