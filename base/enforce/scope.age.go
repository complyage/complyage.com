package enforce

import (
	"strings"

	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Convert Scope
//||------------------------------------------------------------------------------------------------||

func LoadScopesAge(scope string, site Site, user User) (EnforcedScopes, bool, error) {

	var finalScope []models.SiteScope

	if scope != "" {
		cleanScope := strings.Trim(scope, "[] ")
		scopePerm := strings.NewReplacer(" ", "|", ",", "|").Replace(cleanScope)
		scopeParts := strings.Split(scopePerm, "|")
		for _, s := range scopeParts {
			if s == "IDEN" || s == "FACE" || s == "CRCD" {
				scopeItem := models.SiteScope{
					Code:    s,
					Status:  "ACTV",
					Enabled: true,
				}
				finalScope = append(finalScope, scopeItem)
			}
		}
	}

	log.PrettyPrint(finalScope)
	enforcedScope := SiteScopesToEnforcedScopes(finalScope, user)
	return enforcedScope, false, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Missing Count Generate
//||------------------------------------------------------------------------------------------------||

func (es EnforcedScopes) MissingAge() int {
	missing := len(es)
	for _, s := range es {
		if s.Verified {
			missing = 0
		}
	}
	return missing
}
