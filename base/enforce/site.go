package enforce

import (
	"fmt"
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
	ID       uint   `json:"id"`
	ClientId string `json:"clientId"`
	Verifier string `json:"verifier"`
	// Public      string            `json:"-"`
	// Private     string            `json:"-"`
	Name        string            `json:"name"`
	Logo        string            `json:"logo"`
	Description string            `json:"description"`
	URL         string            `json:"url"`
	Redirect    string            `json:"redirect"`
	Webhook     string            `json:"webhook"`
	TestMode    bool              `json:"testMode"`
	Enforcement string            `json:"enforcement"`
	Zones       models.SiteZones  `json:"zones"`
	ScopeAuto   bool              `json:"scopeAuto"`
	Scopes      models.SiteScopes `json:"scopes"`
	State       string            `json:"state"`
}

//||------------------------------------------------------------------------------------------------||
//|| Enforce Site
//||------------------------------------------------------------------------------------------------||

func LoadSite(clientID, hostName string) (Site, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| APIKey
	//||------------------------------------------------------------------------------------------------||

	if clientID == "" {
		return Site{}, app.Err("Enforce").Error("MISSING_CLIENT_ID")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Site By API Key
	//||------------------------------------------------------------------------------------------------||

	site, err := sites.FetchSiteByClientId(clientID)
	if err != nil {
		return Site{}, app.Err("Enforce").Error("INVALID_SITE_KEY")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make sure site is active
	//||------------------------------------------------------------------------------------------------||

	if site.Status != "ACTV" {
		return Site{}, app.Err("Enforce").Error("INVALID_SITE_STATUS")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Hostname
	//||------------------------------------------------------------------------------------------------||

	domainMatch := false
	cleanHost := strings.Replace(hostName, "http://", "", 1)
	cleanHost = strings.Replace(cleanHost, "https://", "", 1)
	hostOnly := strings.Split(cleanHost, ":")[0]
	for _, domain := range strings.Split(site.Domains, ",") {
		d := strings.TrimSpace(domain)
		d = strings.Split(d, ":")[0]
		fmt.Println("Comparing domain:", d, "== CurrentHost:", hostOnly)
		if d == "" {
			continue
		}
		if strings.EqualFold(d, hostOnly) {
			domainMatch = true
			break
		}
	}
	if !domainMatch && site.Domains != "*" {
		fmt.Println("Domain does not match:", hostOnly, site.Domains)
		return Site{}, app.Err("Enforce").Error("INVALID_DOMAIN")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Site
	//||------------------------------------------------------------------------------------------------||

	eSite := Site{
		ID:          site.ID,
		ClientId:    site.ClientID,
		Verifier:    site.CheckKey,
		Name:        site.Name,
		Logo:        site.Logo,
		Description: site.Description,
		URL:         site.URL,
		Webhook:     site.Webhook,
		Enforcement: site.Enforcement,
		Redirect:    site.Redirect,
		TestMode:    site.TestMode,
		Zones:       site.Zones,
		ScopeAuto:   site.ScopeAuto,
		Scopes:      site.Scopes,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Hostname
	//||------------------------------------------------------------------------------------------------||

	return eSite, nil

}
