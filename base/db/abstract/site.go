package abstract

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/oauth"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Helper: OAuthSites
//||------------------------------------------------------------------------------------------------||

func GetSiteByPublic(publicKey string) (models.Site, error) {
	var s models.Site
	err := app.SQLDB["main"].DB.
		Where("site_public = ?", publicKey).
		Where("site_status NOT IN ('RMVD','BNND')").
		First(&s).Error
	if err == nil {
		return s, nil
	}
	return models.Site{}, err
}

//||------------------------------------------------------------------------------------------------||
//|| Helper: OAuthSites
//||------------------------------------------------------------------------------------------------||

func OAuthSite(publicKey string) (oauth.OAuthSite, error) {
	var s models.Site
	if err := app.SQLDB["main"].DB.
		Where("site_public = ?", publicKey).
		Where("site_status NOT IN ('RMVD','BNND')").
		First(&s).Error; err != nil {
		return oauth.OAuthSite{}, err
	}
	return oauth.OAuthSite{
		Name:        s.Name,
		URL:         s.URL,
		Logo:        s.Logo,
		Description: s.Description,
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Load All Sites
//||------------------------------------------------------------------------------------------------||

func ReturnAllSites() ([]models.Site, error) {
	var results []models.Site
	err := app.SQLDB["main"].DB.Table("sites").Where("site_status NOT IN ?", []string{"RMVD", "BNND"}).Find(&results).Error
	if err != nil {
		return []models.Site{}, fmt.Errorf("failed to load sites: %w", err)
	}
	return results, nil
}
