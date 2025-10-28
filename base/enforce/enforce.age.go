package enforce

import (
	"net/http"
	"strings"

	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

func LoadEnforcementAge(r *http.Request) (Enforcement, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Querystring
	//||------------------------------------------------------------------------------------------------||
	clientID := r.URL.Query().Get("client_id")
	hostname := strings.Split(r.Host, ":")[0]
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	site, stErr := LoadSite(clientID, hostname)
	if stErr != nil {
		app.Log.Error(stErr.Error())
		return Enforcement{}, stErr
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Made Zone
	//||------------------------------------------------------------------------------------------------||
	zone := LoadZone(r, site)
	//||------------------------------------------------------------------------------------------------||
	//|| Make Scope
	//||------------------------------------------------------------------------------------------------||
	log.PrettyPrint(zone)
	scope := []string{}
	for i := range zone.Requirements {
		scope = append(scope, zone.Requirements[i].String())
	}
	log.PrettyPrint(scope)
	//||------------------------------------------------------------------------------------------------||
	//|| User
	//||------------------------------------------------------------------------------------------------||
	user := LoadUser(r)
	scopes, isAuto, _ := LoadScopesAge(strings.Join(scope, ","), site, user)
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	hasPrivateKey := false
	if user.Private != "" && user.Public != "" && user.CheckKey != "" {
		err := encrypt.CheckPrivateKey(user.Private, user.CheckKey)
		if err == nil {
			hasPrivateKey = true
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	age, ageErr := LoadAge(scopes, zone)
	if ageErr != nil {
		age = Age{
			Verified:     false,
			ZoneVerified: false,
			Methods:      []types.DataType{},
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||
	return Enforcement{
		Age:        age,
		User:       user,
		Site:       site,
		AutoScope:  isAuto,
		Scopes:     scopes,
		Zone:       zone,
		HasPrivate: hasPrivateKey,
		Missing:    scopes.MissingAge(),
	}, nil
}
