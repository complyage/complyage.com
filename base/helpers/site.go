package helpers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/interfaces"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| Helper: OAuthSites
//||------------------------------------------------------------------------------------------------||

func GetSiteByPublic(publicKey string) (models.Site, error) {

	// load the site record (expect exactly one)
	var s models.Site
	if err := db.DB.
		Where("site_public = ?", publicKey).
		Where("site_status NOT IN ('RMVD','BNND')").
		First(&s).Error; err != nil {
		return s, err
	}

	// map to OAuthSite
	return models.Site{}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Helper: OAuthSites
//||------------------------------------------------------------------------------------------------||

func OAuthSite(publicKey string) (interfaces.OAuthSite, error) {

	// load the site record (expect exactly one)
	var s models.Site
	if err := db.DB.
		Where("site_public = ?", publicKey).
		Where("site_status NOT IN ('RMVD','BNND')").
		First(&s).Error; err != nil {
		return interfaces.OAuthSite{}, err
	}

	// map to OAuthSite
	return interfaces.OAuthSite{
		Name:        s.SiteName,
		URL:         s.SiteURL,
		Logo:        s.SiteLogo,
		Description: s.SiteDescription,
	}, nil
}
