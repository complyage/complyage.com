package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| UpdateSiteHandler :: Updates an existing site
//||------------------------------------------------------------------------------------------------||

func SitesUpdateHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decode JSON Body
	//||------------------------------------------------------------------------------------------------||

	var input models.Site
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON: "+err.Error())
		return
	}

	if input.ID == 0 {
		responses.Error(w, http.StatusBadRequest, "Missing site ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Original Site
	//||------------------------------------------------------------------------------------------------||

	var original models.Site
	if err := app.SQLDB["main"].DB.
		Where("id_site = ? AND fid_account = ?", input.ID, session.ID).
		Where("site_status NOT IN ('RMVD', 'BNND')").
		First(&original).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Site not found or access denied")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Determine Status Update Rules
	//||------------------------------------------------------------------------------------------------||

	newStatus := original.Status

	if original.Status == "PNEW" {
		newStatus = "PEND"
	} else if (original.Status == "APPR" || original.Status == "ACTV") &&
		(original.Name != input.Name || original.URL != input.URL) {
		newStatus = "PEND"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Perform Update
	//||------------------------------------------------------------------------------------------------||

	err = app.SQLDB["main"].DB.Model(&models.Site{}).
		Where("id_site = ? AND fid_account = ?", input.ID, session.ID).
		Updates(map[string]interface{}{
			"site_name":         input.Name,
			"site_description":  input.Description,
			"site_url":          input.URL,
			"site_status":       newStatus,
			"site_zones":        input.Zones,
			"site_enforcement":  input.Enforcement,
			"site_domains":      input.Domains,
			"site_private":      input.Private,
			"site_public":       input.Public,
			"site_redirect":     input.Redirect,
			"site_permissions":  input.Permissions,
			"site_testmode":     input.TestMode,
			"site_gate_signup":  input.GateSignup,
			"site_gate_confirm": input.GateConfirm,
			"site_gate_exit":    input.GateExit,
		}).Error

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, fmt.Sprintf("Database error while updating site %d: %v", input.ID, err))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	input.Status = newStatus

	responses.Success(w, http.StatusOK, map[string]any{
		"site": input,
	})
}
