package enforce

import (
	"strings"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/identity"
	"github.com/complyage/base/scopes"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| EnforcedScope
//||------------------------------------------------------------------------------------------------||

type EnforcedScope struct {
	Code         string
	Description  string
	Icon         string
	Level        int
	Verified     bool
	Verification string
}

type EnforcedScopes []EnforcedScope

//||------------------------------------------------------------------------------------------------||
//|| Convert Scope
//||------------------------------------------------------------------------------------------------||

func LoadScopes(scope string, siteScopes models.SiteScopes, user User) (EnforcedScopes, bool, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Handle Scope
	//||------------------------------------------------------------------------------------------------||

	if SiteScopesIsAuto(siteScopes) {
		return EnforcedScopes{}, true, nil
	}

	//||------------------------------------------------------------------------------------------------||
	//||
	//||------------------------------------------------------------------------------------------------||

	requestedScopes := siteScopes

	//||------------------------------------------------------------------------------------------------||
	//|| Handle Scope
	//||------------------------------------------------------------------------------------------------||

	if scope != "" {

		//||------------------------------------------------------------------------------------------------||
		//|| Clean and Split Requested Scopes
		//||------------------------------------------------------------------------------------------------||

		cleanScope := strings.Trim(scope, "[] ") // remove brackets
		scopePerm := strings.NewReplacer(" ", "|", ",", "|").Replace(cleanScope)
		scopeParts := strings.Split(scopePerm, "|")

		//||------------------------------------------------------------------------------------------------||
		//|| Clean and Split Requested Scopes
		//||------------------------------------------------------------------------------------------------||

		for _, s := range scopeParts {

			//||------------------------------------------------------------------------------------------------||
			//|| Skip Blanks
			//||------------------------------------------------------------------------------------------------||

			s = strings.TrimSpace(strings.ToUpper(s))
			if s == "" {
				continue
			}

			//||------------------------------------------------------------------------------------------------||
			//|| Check if this scope is in allowed site permissions
			//||------------------------------------------------------------------------------------------------||

			found := false
			for _, ss := range siteScopes {
				if strings.EqualFold(s, ss.Code) {
					found = true
					//||------------------------------------------------------------------------------------------------||
					//|| An unapproved scope was requested
					//||------------------------------------------------------------------------------------------------||
					if !ss.Enabled {
						return EnforcedScopes{}, false, app.Err("OAuth").Error("DISABLED_SCOPE")
					}
					//||------------------------------------------------------------------------------------------------||
					//|| TODO Next Version :: Check if if Approval is needed
					//||------------------------------------------------------------------------------------------------||
					requestedScopes = append(requestedScopes, ss)
					break
				}
			}

			//||------------------------------------------------------------------------------------------------||
			//|| An unapproved scope was requested
			//||------------------------------------------------------------------------------------------------||

			if !found {
				return EnforcedScopes{}, false, app.Err("OAuth").Error("INVALID_SCOPE")
			}

		}
	}

	enforcedScope := SiteScopesToEnforcedScopes(requestedScopes, user)
	return enforcedScope, false, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Check for Auto
//||------------------------------------------------------------------------------------------------||

func SiteScopesIsAuto(scopes models.SiteScopes) bool {
	for _, s := range scopes {
		if strings.EqualFold(s.Code, "AUTO") {
			return true
		}
	}
	return false
}

//||------------------------------------------------------------------------------------------------||
//|| Create Enforced Scopes
//||------------------------------------------------------------------------------------------------||

func SiteScopesToEnforcedScopes(siteScopes models.SiteScopes, user User) EnforcedScopes {
	enforcedScopes := EnforcedScopes{}
	for _, s := range siteScopes {
		isVerified, verification := GetScopeVerified(s.Code, user.Identity)
		enforced := EnforcedScope{
			Code:         s.Code,
			Description:  scopes.Description(s.Code),
			Icon:         scopes.Icon(s.Code),
			Level:        scopes.Level(s.Code),
			Verified:     isVerified,
			Verification: verification,
		}
		enforcedScopes = append(enforcedScopes, enforced)
	}
	return enforcedScopes
}

//||------------------------------------------------------------------------------------------------||
//|| Create Site Scopes
//||------------------------------------------------------------------------------------------------||

func GetScopeVerified(scope string, identity identity.Identity) (bool, string) {
	scope = strings.TrimSpace(strings.ToUpper(scope))
	switch scope {
	case "MAIL":
		return identity.Email.Verified, identity.Email.Verification
	case "PHNE":
		return identity.Phone.Verified, identity.Phone.Verification
	case "ADDR":
		return identity.Address.Verified, identity.Address.Verification
	case "IDEN":
		return identity.IDCard.Verified, identity.IDCard.Verification
	case "CRCD":
		return identity.CreditCard.Verified, identity.CreditCard.Verification
	case "FACE":
		return identity.Face.Verified, identity.Face.Verification
	}
	return false, ""
}
