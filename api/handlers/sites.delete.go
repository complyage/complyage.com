package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
	"net/http"
	"strconv"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesDeleteHandler(w http.ResponseWriter, r *http.Request) {

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

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
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
	//|| Delete Site if Owned by User
	//||------------------------------------------------------------------------------------------------||

	if err := db.DB.
		Model(&models.Site{}).
		Where("id_site = ? AND fid_account = ?", id, session.ID).
		Update("site_status", "RMVD").Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to mark site as removed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id":      id,
		"success": true,
	})
}
