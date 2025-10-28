package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
	"strconv"

	"github.com/complyage/base/db/models"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

type siteLoadResponse struct {
	Site models.Site `json:"site"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesLoadHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	_, _, session, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse ID Param
	//||------------------------------------------------------------------------------------------------||

	idStr := r.URL.Query().Get("id")
	if idStr == "" {
		responses.Error(w, http.StatusBadRequest, "Missing id parameter")
		return
	}

	id, err := strconv.Atoi(idStr)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid id parameter")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Site
	//||------------------------------------------------------------------------------------------------||

	var site models.Site
	err = app.SQLDB["main"].DB.
		Where("id_site = ? AND fid_account = ?", id, session.ID).
		Where("site_status NOT IN ('RMVD', 'BNND')").
		First(&site).Error

	if err != nil {
		responses.Error(w, http.StatusNotFound, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Logo Hash
	//||------------------------------------------------------------------------------------------------||

	checkLogo, logoErr := app.Storages["sites"].Get(site.Logo)
	if logoErr != nil || checkLogo == nil || len(checkLogo) == 0 {
		site.Logo = ""
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Site
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, siteLoadResponse{
		Site: site,
	})
}
