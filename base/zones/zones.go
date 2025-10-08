package zones

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/db/models"
	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| In‑memory slice of zones (basic info)
//||------------------------------------------------------------------------------------------------||

var Zones []models.Zone

//||------------------------------------------------------------------------------------------------||
//|| Package init – automatically start refresher on import
//||------------------------------------------------------------------------------------------------||

func init() {
	StartZoneRefresher(30) // refresh every 10 minutes
}

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
		if Zones[i].ID == id {
			return &Zones[i], true
		}
	}
	return nil, false
}

//||------------------------------------------------------------------------------------------------||
//|| StartZoneRefresher – refresh zones in memory every X minutes
//||------------------------------------------------------------------------------------------------||

func StartZoneRefresher(minutes int) {
	ticker := time.NewTicker(time.Duration(minutes) * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := LoadZones(); err != nil {
					log.Printf("[ZONE] Failed to reload zones: %v\n", err)
				}
			}
		}
	}()
}

//||------------------------------------------------------------------------------------------------||
//|| FindZoneByLocation
//||------------------------------------------------------------------------------------------------||

func FindZoneByLocation(state, country string) (*models.Zone, bool) {
	// unknown
	if state == "" || country == "" {
		for i := range Zones {
			if Zones[i].ID == 9999 {
				return &Zones[i], true
			}
		}
		return nil, false
	}

	// exact state match
	for i := range Zones {
		z := &Zones[i]
		if z.Region != nil && *z.Region == state && z.Country != nil && *z.Country == country {
			return z, true
		}
	}

	// fallback: country‑wide match (state null but country matches)
	for i := range Zones {
		z := &Zones[i]
		if z.Region == nil && z.Country != nil && *z.Country == country {
			return z, true
		}
	}

	// no specific match, but caller treats zoneFound==true as “no enforcement needed”
	return nil, false
}

//||------------------------------------------------------------------------------------------------||
//|| Fetch Short Zone
//||------------------------------------------------------------------------------------------------||

func FetchShortZoneByLocation(state, country string) (*ShortZone, bool) {
	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Short Zone
	//||------------------------------------------------------------------------------------------------||
	zone, ok := FindZoneByLocation(state, country)
	if !ok {
		return nil, false
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Make Requirements Slice
	//||------------------------------------------------------------------------------------------------||
	if zone.Requirements == nil {
		return nil, false
	}
	raw := *zone.Requirements
	parts := strings.Split(raw, ",")
	requirements := make([]types.DataType, 0, len(parts))
	for _, part := range parts {
		dt, exists := types.StringToDataType(strings.TrimSpace(part))
		if exists {
			requirements = append(requirements, dt)
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Make Requirements Slice
	//||------------------------------------------------------------------------------------------------||
	shortZone := ShortZone{
		ID:           int(zone.ID),
		Law:          "",
		Requirements: requirements,
		Effective:    zone.Effective.String(),
		MinAge:       zone.MinAge,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Law
	//||------------------------------------------------------------------------------------------------||
	if zone.Law != nil {
		shortZone.Law = *zone.Law
	}
	return &shortZone, true
}
