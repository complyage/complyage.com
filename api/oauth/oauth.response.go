package oauth

import (
	"base/interfaces"
)

//||------------------------------------------------------------------------------------------------||
//|| Create Response
//||------------------------------------------------------------------------------------------------||

func CreateResponse() interfaces.OAuthResponse {

	//||------------------------------------------------------------------------------------------------||
	//|| OAuth User
	//||------------------------------------------------------------------------------------------------||

	user := interfaces.OAuthUser{
		Status:        "NONE",
		Username:      "Anonymous",
		Verifications: []interfaces.OAuthVerification{},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Site
	//||------------------------------------------------------------------------------------------------||

	site := interfaces.OAuthSite{
		Name:        "Unknown Site",
		URL:         "https://unknown.site",
		Logo:        "/img/public/logo.missing.png",
		Description: "Site not found. No description available",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Zone
	//||------------------------------------------------------------------------------------------------||

	zone := interfaces.OAuthZone{
		State:         "Unknown",
		Country:       "Unknown",
		IP:            "Unknown",
		Requirements:  []interfaces.OAuthRequirements{},
		Description:   "No zone information available",
		Law:           "No law information available",
		EffectiveDate: "Unknown",
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	response := interfaces.OAuthResponse{
		Site:         site,
		User:         user,
		Zone:         zone,
		Status:       "FAIL",
		Requirements: []interfaces.OAuthRequirements{},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	return response
}
