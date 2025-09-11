package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"net/http"
	"strconv"
	"strings"

	"github.com/ralphferrara/aria/base/random"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesCopyHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	_, _, session, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session. Please re-login")
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
	//|| Load Source Site
	//||------------------------------------------------------------------------------------------------||

	var source models.Site
	if err := app.SQLDB["main"].DB.
		Where("id_site = ? AND fid_account = ?", id, session.ID).
		First(&source).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Site not found or access denied")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update Fields
	//||------------------------------------------------------------------------------------------------||

	newSite := source
	newSite.Status = "PNEW"
	newSite.Private = random.RandomString(32)
	newSite.Public = random.RandomString(32)
	newSite.AgentPrivate = random.RandomString(32)
	if !strings.HasSuffix(newSite.Name, " - Copy") {
		newSite.Name += " - Copy"
	} else {
		newSite.Name += " (" + random.RandomString(4) + ")"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create New Copied Site
	//||------------------------------------------------------------------------------------------------||

	if err := app.SQLDB["main"].DB.Create(&newSite).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create copied site")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return New Site ID
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id": newSite.ID,
	})
}
