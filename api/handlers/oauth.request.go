package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"api/oauth"
	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
	"net/http"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func OAuthResponseHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the general response
	//||------------------------------------------------------------------------------------------------||

	oauthResponse := oauth.CreateResponse()

	//||------------------------------------------------------------------------------------------------||
	//|| Parse site_public Token
	//||------------------------------------------------------------------------------------------------||

	token := r.URL.Query().Get("token")
	if token == "" {
		responses.Error(w, http.StatusBadRequest, "Missing token parameter")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Site by public token
	//||------------------------------------------------------------------------------------------------||

	site, err := helpers.OAuthSite(token)
	if err == nil {
		oauthResponse.Site = site
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	myCookie := ""
	if err == nil {
		myCookie = cookie.Value
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	oauthResponse.User = oauth.CreateUserResponse(myCookie)

	//||------------------------------------------------------------------------------------------------||
	//|| Determine Client IP and Lookup IP record
	//||------------------------------------------------------------------------------------------------||

	clientIP := helpers.GetClientIP(r)
	location, err := helpers.GetLocationByIP(clientIP)

	if err == nil && location.Country != "" && location.State != "" {
		var zone models.Zone
		if err := db.DB.
			Where("zone_country = ? AND zone_state = ?", location.Country, location.State).
			Order("zone_effective DESC").
			Limit(1).
			First(&zone).Error; err != nil {
		}

		reqs := []string{}
		for _, t := range strings.Split(*zone.ZoneRequirements, ",") {
			if s := strings.TrimSpace(t); s != "" {
				reqs = append(reqs, s)
			}
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, oauthResponse)
}
