package enforce

import (
	"net/http"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

type Enforcement struct {
	User      User
	Site      Site
	AutoScope bool
	Scopes    EnforcedScopes
	Zone      EnforcedZone
	State     string
}

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

func LoadEnforcement(r *http.Request) (Enforcement, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Querystring
	//||------------------------------------------------------------------------------------------------||
	clientID := r.URL.Query().Get("client_id")
	scope := r.URL.Query().Get("scope")
	state := r.URL.Query().Get("state")
	hostname := strings.Split(r.Host, ":")[0]
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	site, stErr := LoadSite(clientID, hostname)
	if stErr != nil {
		return Enforcement{}, stErr
	}
	//||------------------------------------------------------------------------------------------------||
	//|| User
	//||------------------------------------------------------------------------------------------------||
	user := LoadUser(r)
	scopes, isAuto, _ := LoadScopes(scope, site.Scopes, user)
	zone := LoadZone(r, site)
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	return Enforcement{
		User:      user,
		Site:      site,
		AutoScope: isAuto,
		Scopes:    scopes,
		Zone:      zone,
		State:     state,
	}, nil
}
