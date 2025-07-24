package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/responses"
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

type SiteListResponse struct {
	IDSite   string `json:"id"        gorm:"column:id_site"`
	SiteName string `json:"name"      gorm:"column:site_name"`
	SiteURL  string `json:"url"       gorm:"column:site_url"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesListHandler(w http.ResponseWriter, r *http.Request) {

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
	//|| Fetch Sites
	//||------------------------------------------------------------------------------------------------||

	var sites []SiteListResponse
	if err := db.DB.
		Table("sites").
		Select("id_site", "site_name", "site_url").
		Where("fid_account = ?", session.ID).
		Where("site_status NOT IN ('RMVD', 'BNND')").
		Order("site_name ASC").
		Scan(&sites).Error; err != nil {

		responses.Error(w, http.StatusInternalServerError, "Error loading sites")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"sites": sites,
	})
}
