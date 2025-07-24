package helpers

import (
	"base/interfaces"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| OptimizeSite
//||------------------------------------------------------------------------------------------------||

func OptimizeSite(site models.Site) interfaces.OptimizedSite {
	return interfaces.OptimizedSite{
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

func OptimizeZone(z models.Zone) interfaces.OptimizedZone {
	return interfaces.OptimizedZone{
		State:       z.ZoneState,
		Country:     z.ZoneCountry,
		Law:         z.ZoneLaw,
		Description: z.ZoneLawDescription,
		Effective:   z.ZoneEffective,
	}
}
