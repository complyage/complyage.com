package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/base/random"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesNewHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Record
	//||------------------------------------------------------------------------------------------------||

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Public/Private Keys
	//||------------------------------------------------------------------------------------------------||

	privKey := random.RandomString(32)
	pubKey := random.RandomString(32)

	//||------------------------------------------------------------------------------------------------||
	//|| Create New Site Model
	//||------------------------------------------------------------------------------------------------||

	site := models.Site{
		FidAccount:      session.ID,
		SiteStatus:      "PNEW",
		SitePrivate:     privKey,
		SitePublic:      pubKey,
		SiteEnforcement: "ALLZ",
	}

	if err := app.SQLDB["main"].DB.Create(&site).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create site")
		fmt.Println("Error creating site:", err)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Created Site
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id": site.IDSite,
	})
}
