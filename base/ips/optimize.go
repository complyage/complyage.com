package ips

import (
	"github.com/complyage/base/db/models"
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
		Scopes:      site.Scopes,
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Optimize Zone
//||------------------------------------------------------------------------------------------------||

func OptimizeZone(z models.Zone) OptimizedZone {
	return OptimizedZone{
		Region:      z.Region,
		Country:     z.Country,
		Law:         z.Law,
		Description: z.Description,
		Effective:   z.Effective,
	}
}
