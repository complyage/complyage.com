package enforce

import (
	"strings"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/sites"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Zone
//||------------------------------------------------------------------------------------------------||

type Zone struct {
	Zone    string
	Enforce string
}

//||------------------------------------------------------------------------------------------------||
//|| Site
//||------------------------------------------------------------------------------------------------||

type Site struct {
	ClientID    string
	Name        string
	Logo        string
	Description string
	URL         string
	Redirect    string
	TestMode    bool
	Enforcement string
	Zones       models.SiteZones
	Scopes      models.SiteScopes
}

//||------------------------------------------------------------------------------------------------||
//|| Enforce Site
//||------------------------------------------------------------------------------------------------||

func LoadSite(clientID, hostName string) (Site, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| APIKey
	//||------------------------------------------------------------------------------------------------||

	if clientID == "" {
		return Site{}, app.Err("OAuth").Error("MISSING_CLIENT_ID")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Site By API Key
	//||------------------------------------------------------------------------------------------------||

	site, err := sites.FetchSiteByPublic(clientID)
	if err != nil {
		return Site{}, app.Err("OAuth").Error("INVALID_SITE_KEY")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make sure site is active
	//||------------------------------------------------------------------------------------------------||

	if site.Status != "ACTV" {
		return Site{}, app.Err("OAuth").Error("INVALID_SITE_STATUS")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Hostname
	//||------------------------------------------------------------------------------------------------||

	domainMatch := false
	for _, domain := range strings.Split(site.Domains, ",") {
		d := strings.TrimSpace(domain)
		if d == "" {
			continue
		}
		if strings.EqualFold(d, hostName) {
			domainMatch = true
			break
		}
	}

	if !domainMatch && site.Domains != "*" {
		return Site{}, app.Err("OAuth").Error("INVALID_DOMAIN")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Site
	//||------------------------------------------------------------------------------------------------||

	eSite := Site{
		ClientID:    site.Public,
		Name:        site.Name,
		Logo:        site.Logo,
		Description: site.Description,
		URL:         site.URL,
		Enforcement: site.Enforcement,
		Redirect:    site.Redirect,
		TestMode:    site.TestMode,
		Zones:       site.Zones,
		Scopes:      site.Scopes,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Hostname
	//||------------------------------------------------------------------------------------------------||

	return eSite, nil

}
