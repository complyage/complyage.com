package enforce

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/ips"
	"github.com/complyage/base/types"
	"github.com/complyage/base/zones"
	ariaHTTP "github.com/ralphferrara/aria/http"
)

//||------------------------------------------------------------------------------------------------||
//|| EnforcedZone
//||------------------------------------------------------------------------------------------------||

type EnforcedZone struct {
	ID           int64            `json:"id"`
	IPAddress    string           `json:"ipAddress"`
	Region       string           `json:"region"`
	Country      string           `json:"country"`
	Requirements []types.DataType `json:"requirements"`
	Enforced     bool             `json:"enforced"`
}

//||------------------------------------------------------------------------------------------------||
//|| Load Zone
//||------------------------------------------------------------------------------------------------||

func LoadZone(r *http.Request, site Site) EnforcedZone {

	setZone := EnforcedZone{
		ID:        0,
		Region:    "Unknown",
		Country:   "Location",
		IPAddress: "0.0.0.0",
		Enforced:  ZoneIsEnforced(zones.ShortZone{ID: 9999}, site),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get IP
	//||------------------------------------------------------------------------------------------------||

	ipAddress := ariaHTTP.GetClientIP(r)
	fmt.Println("Client IP Address:", ipAddress)
	if ipAddress == "" {
		return setZone
	}
	setZone.IPAddress = ipAddress

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch IP Location
	//||------------------------------------------------------------------------------------------------||

	location, err := ips.GetLocationByIP(ipAddress)
	if err != nil {
		return setZone
	}
	setZone.Region = location.Region
	setZone.Country = location.Country

	//||------------------------------------------------------------------------------------------------||
	//|| Get Zone
	//||------------------------------------------------------------------------------------------------||

	zone, ok := zones.FetchShortZoneByLocation(location.Region, location.Country)
	if !ok {
		return setZone
	}
	setZone.ID = int64(zone.ID)
	setZone.Requirements = zone.Requirements
	setZone.Enforced = ZoneIsEnforced(*zone, site)
	return setZone
}

//||------------------------------------------------------------------------------------------------||
//|| Zone is Enforced
//||------------------------------------------------------------------------------------------------||

func (zone EnforcedZone) AllowsMethod(method types.DataType) bool {
	for _, m := range zone.Requirements {
		if m == method {
			return true
		}
	}
	return false
}

//||------------------------------------------------------------------------------------------------||
//|| Zone is Enforced
//||------------------------------------------------------------------------------------------------||

func ZoneIsEnforced(zone zones.ShortZone, site Site) bool {
	//||------------------------------------------------------------------------------------------------||
	//|| Force All
	//||------------------------------------------------------------------------------------------------||
	if site.Enforcement == "ALLZ" {
		return true
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Regulated Only
	//||------------------------------------------------------------------------------------------------||
	if site.Enforcement == "REGU" {
		if zone.ID == 9999 {
			return false
		} else {
			return true
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Custom
	//||------------------------------------------------------------------------------------------------||
	for _, z := range site.Zones {
		if z.Zone == zone.ID {
			return z.Enforced
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Default True
	//||------------------------------------------------------------------------------------------------||
	return true
}
