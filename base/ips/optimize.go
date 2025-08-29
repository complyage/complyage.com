package ips

import (
	"base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| OptimizeSite
//||------------------------------------------------------------------------------------------------||

func OptimizeSite(site models.Site) OptimizedSite {
	return OptimizedSite{
		Name:        site.SiteName,
		Logo:        site.SiteLogo,
		Description: site.SiteDescription,
		URL:         site.SiteURL,
		Redirect:    site.SiteRedirect,
		Permissions: site.SitePermissions,
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
