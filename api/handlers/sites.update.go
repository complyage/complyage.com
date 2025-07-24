package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
	"encoding/json"
	"fmt"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| UpdateSiteHandler :: Updates an existing site
//||------------------------------------------------------------------------------------------------||

func SitesUpdateHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Check Method
	//||------------------------------------------------------------------------------------------------||

	if r.Method != http.MethodPost {
		responses.Error(w, http.StatusMethodNotAllowed, "Only POST allowed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	session, err := helpers.FetchSession(cookie.Value)
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

	if input.IDSite == 0 {
		responses.Error(w, http.StatusBadRequest, "Missing site ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load Original Site
	//||------------------------------------------------------------------------------------------------||

	var original models.Site
	if err := db.DB.
		Where("id_site = ? AND fid_account = ?", input.IDSite, session.ID).
		Where("site_status NOT IN ('RMVD', 'BNND')").
		First(&original).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Site not found or access denied")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Determine Status Update Rules
	//||------------------------------------------------------------------------------------------------||

	newStatus := original.SiteStatus

	if original.SiteStatus == "PNEW" {
		newStatus = "PEND"
	} else if (original.SiteStatus == "APPR" || original.SiteStatus == "ACTV") &&
		(original.SiteName != input.SiteName || original.SiteURL != input.SiteURL) {
		newStatus = "PEND"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Perform Update
	//||------------------------------------------------------------------------------------------------||

	//||------------------------------------------------------------------------------------------------||
	//|| Perform Update
	//||------------------------------------------------------------------------------------------------||

	err = db.DB.Model(&models.Site{}).
		Where("id_site = ? AND fid_account = ?", input.IDSite, session.ID).
		Updates(map[string]interface{}{
			"site_name":         input.SiteName,
			"site_description":  input.SiteDescription,
			"site_url":          input.SiteURL,
			"site_status":       newStatus,
			"site_zones":        input.SiteZones,
			"site_enforcement":  input.SiteEnforcement,
			"site_domains":      input.SiteDomains,
			"site_private":      input.SitePrivate,
			"site_public":       input.SitePublic,
			"site_redirect":     input.SiteRedirect,
			"site_permissions":  input.SitePermissions,
			"site_testmode":     input.SiteTestMode,
			"site_gate_signup":  input.SiteGateSignup,
			"site_gate_confirm": input.SiteGateConfirm,
			"site_gate_exit":    input.SiteGateExit,
		}).Error

	if err != nil {
		responses.Error(w, http.StatusInternalServerError, fmt.Sprintf("Database error while updating site %d: %v", input.IDSite, err))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	input.SiteStatus = newStatus

	responses.Success(w, http.StatusOK, map[string]any{
		"site": input,
	})
}
