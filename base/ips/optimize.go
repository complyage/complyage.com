package ips

import (
	"base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| OptimizeSite
//||------------------------------------------------------------------------------------------------||

func OptimizeSite(site models.Site) OptimizedSite {
	return OptimizedSite{
		Name:        site.Name,
		Logo:        site.Logo,
		Description: site.Description,
		URL:         site.URL,
		Redirect:    site.Redirect,
		Permissions: site.Permissions,
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Optimize Zone
//||------------------------------------------------------------------------------------------------||

func OptimizeZone(z models.Zone) OptimizedZone {
	return OptimizedZone{
		State:       z.ZoneState,
		Country:     z.ZoneCountry,
		Law:         z.ZoneLaw,
		Description: z.ZoneLawDescription,
		Effective:   z.ZoneEffective,
	}
}
