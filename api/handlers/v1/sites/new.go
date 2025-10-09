package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/db/models"

	"github.com/ralphferrara/aria/base/random"
	"github.com/ralphferrara/aria/base/sign"
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

	privPEM, pubPEM, err := sign.GenerateKeyPair()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate keys")
		fmt.Println("Error generating keys:", err)
		return
	}
	clientId := "CLID-" + random.RandomString(27)
	agentKey := "AGNT-" + random.RandomString(8) + "-" + random.RandomString(8) + "-" + random.RandomString(8)
	validateKey := random.RandomString(32)

	//||------------------------------------------------------------------------------------------------||
	//|| Create New Site Model
	//||------------------------------------------------------------------------------------------------||

	site := models.Site{
		FidAccount:   session.ID,
		Status:       "PNEW",
		ClientID:     clientId,
		Private:      privPEM,
		Public:       pubPEM,
		TestMode:     true,
		Domains:      "*",
		Enforcement:  "ALLZ",
		CheckKey:     validateKey,
		AgentPrivate: agentKey,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Record
	//||------------------------------------------------------------------------------------------------||

	if err := app.SQLDB["main"].DB.Create(&site).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create site")
		fmt.Println("Error creating site:", err)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Created Site
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id": site.ID,
	})
}
