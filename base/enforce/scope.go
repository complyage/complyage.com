package enforce

import (
	"fmt"
	"strings"

	"github.com/complyage/base/db/models"
	"github.com/complyage/base/identity"
	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| EnforcedScope
//||------------------------------------------------------------------------------------------------||

type EnforcedScope struct {
	Code         types.DataType `json:"code"`
	Verified     bool           `json:"verified"`
	Verification string         `json:"verification"`
	DOB          types.DOB      `json:"dob,omitempty"`
}

type EnforcedScopes []EnforcedScope

//||------------------------------------------------------------------------------------------------||
//|| Convert Scope
//||------------------------------------------------------------------------------------------------||

func LoadScopes(scope string, site Site, user User) (EnforcedScopes, bool, error) {

	var finalScope []models.SiteScope

	if site.ScopeAuto {
		return EnforcedScopes{}, true, nil
	}

	if scope != "" {
		cleanScope := strings.Trim(scope, "[] ")
		scopePerm := strings.NewReplacer(" ", "|", ",", "|").Replace(cleanScope)
		scopeParts := strings.Split(scopePerm, "|")
		if len(scopeParts) == 0 {
			return EnforcedScopes{}, false, app.Err("OAuth").Error("INVALID_SCOPE")
		}

		for _, s := range scopeParts {
			part := strings.TrimSpace(s)
			if part == "" {
				continue
			}

			// normalize to DataType
			dt, ok := types.StringToDataType(part)
			if !ok {
				return EnforcedScopes{}, false, app.Err("OAuth").Error("INVALID_SCOPE")
			}

			found := false
			for _, ss := range site.Scopes {
				// site scope code → DataType
				siteDt, ok := types.StringToDataType(ss.Code)
				if !ok {
					continue
				}
				if dt == siteDt {
					found = true
					if !ss.Enabled {
						return EnforcedScopes{}, false, app.Err("OAuth").Error("DISABLED_SCOPE")
					}
					finalScope = append(finalScope, ss)
					break
				}
			}

			if !found {
				return EnforcedScopes{}, false, app.Err("OAuth").Error("INVALID_SCOPE")
			}
		}
	}

	log.PrettyPrint(finalScope)
	enforcedScope := SiteScopesToEnforcedScopes(finalScope, user)
	return enforcedScope, false, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Scopes List
//||------------------------------------------------------------------------------------------------||

func (es *EnforcedScopes) ToList() []string {
	var list []string
	for _, s := range *es {
		list = append(list, s.Code.String())
	}
	return list
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
		dt, ok := types.StringToDataType(s.Code)
		if !ok {
			fmt.Println("Error converting scope code to DataType:", s.Code)
			continue
		}
		isVerified, verification, enforcedDOB := GetScopeVerified(dt, user.Identity)
		enforced := EnforcedScope{
			Code:         dt,
			Verified:     isVerified,
			Verification: verification,
			DOB:          enforcedDOB,
		}
		enforcedScopes = append(enforcedScopes, enforced)
	}
	return enforcedScopes
}

//||------------------------------------------------------------------------------------------------||
//|| Create Site Scopes
//||------------------------------------------------------------------------------------------------||

func GetScopeVerified(scope types.DataType, identity identity.Identity) (bool, string, types.DOB) {
	switch scope {
	case types.DATATYPES.MAIL:
		return identity.Email.Verified, identity.Email.Verification, types.DOB{}
	case types.DATATYPES.PHNE:
		return identity.Phone.Verified, identity.Phone.Verification, types.DOB{}
	case types.DATATYPES.ADDR:
		return identity.Address.Verified, identity.Address.Verification, types.DOB{}
	case types.DATATYPES.IDEN:
		return identity.IDCard.Verified, identity.IDCard.Verification, identity.IDCard.DOB
	case types.DATATYPES.CRCD:
		return identity.CreditCard.Verified, identity.CreditCard.Verification, identity.CreditCard.DOB
	case types.DATATYPES.FACE:
		return identity.Face.Verified, identity.Face.Verification, identity.Face.DOB
	default:
		return false, "", types.DOB{}
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Missing Count Generate
//||------------------------------------------------------------------------------------------------||

func (es EnforcedScopes) Missing() int {
	missing := 0
	for _, s := range es {
		if !s.Verified {
			missing++
		}
	}
	return missing
}
