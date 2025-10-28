package enforce

import (
	"net/http"
	"strings"

	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/base/encrypt"
)

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

type Enforcement struct {
	Age        Age            `json:"age"`
	User       User           `json:"user"`
	Site       Site           `json:"site"`
	AutoScope  bool           `json:"autoScope"`
	Scopes     EnforcedScopes `json:"scopes"`
	Zone       EnforcedZone   `json:"zone"`
	HasPrivate bool           `json:"hasPrivate"`
	Missing    int            `json:"missing"` // how many required scopes are missing verification
	State      string         `json:"state"`
	RetunURL   string         `json:"returnURL"`
}

//||------------------------------------------------------------------------------------------------||
//|| Serves the HTML file with dynamic replacements
//||------------------------------------------------------------------------------------------------||

func LoadEnforcementScopes(r *http.Request) (Enforcement, error) {
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
	scopes, isAuto, _ := LoadScopes(scope, site, user)
	zone := LoadZone(r, site)
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
		Missing:    scopes.Missing(),
		State:      state,
	}, nil
}
