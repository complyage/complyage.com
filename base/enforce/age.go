package enforce

import (
	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/log"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct Age
//||------------------------------------------------------------------------------------------------||

type Age struct {
	Verified     bool             `json:"verified"`
	ZoneVerified bool             `json:"zone_verified"`
	DOB          types.DOB        `json:"dob,omitempty"`
	Methods      []types.DataType `json:"methods,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Load Age
//||------------------------------------------------------------------------------------------------||

func LoadAge(scopes EnforcedScopes, zone EnforcedZone) (Age, error) {
	age := Age{
		Verified:     false,
		ZoneVerified: false,
	}
	level := 0
	for _, s := range scopes {
		app.Log.Info("Checking Age Scope:", s.Code.String(), s.Verified)
		log.PrettyPrint(s)
		//||------------------------------------------------------------------------------------------------||
		//|| Lowest Form
		//||------------------------------------------------------------------------------------------------||
		if s.Code == types.DataTypeCRCD && s.Verified {
			age.Methods = append(age.Methods, s.Code)
			if zone.AllowsMethod(s.Code) {
				age.ZoneVerified = true
			}
			if level < 1 {
				level = 1
				age.Verified = true
				age.DOB = s.DOB
			}
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Medium
		//||------------------------------------------------------------------------------------------------||
		if s.Code == types.DataTypeFACE && s.Verified {
			age.Methods = append(age.Methods, s.Code)
			if zone.AllowsMethod(s.Code) {
				age.ZoneVerified = true
			}
			if level < 2 {
				level = 2
				age.Verified = true
				age.DOB = s.DOB
			}
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Medium
		//||------------------------------------------------------------------------------------------------||
		if s.Code == types.DataTypeIDEN && s.Verified {
			age.Methods = append(age.Methods, s.Code)
			if zone.AllowsMethod(s.Code) {
				age.ZoneVerified = true
			}
			if level < 3 {
				level = 3
				age.Verified = true
				age.DOB = s.DOB
			}
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Medium
	//||------------------------------------------------------------------------------------------------||

	return age, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Has Method
//||------------------------------------------------------------------------------------------------||

func (a Age) HasMethod(method types.DataType) bool {
	for _, m := range a.Methods {
		if m == method {
			return true
		}
	}
	return false
}

//||------------------------------------------------------------------------------------------------||
//|| Has Method String
//||------------------------------------------------------------------------------------------------||

func (a Age) HasMethodString(method string) bool {
	for _, m := range a.Methods {
		if m.String() == method {
			return true
		}
	}
	return false
}
