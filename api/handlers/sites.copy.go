package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/models"
	"base/responses"
	"encoding/hex"
	"net/http"
	"strconv"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesCopyHandler(w http.ResponseWriter, r *http.Request) {

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
	//|| Load Source Site
	//||------------------------------------------------------------------------------------------------||

	var source models.Site
	if err := db.DB.
		Where("id_site = ? AND fid_account = ?", id, session.ID).
		First(&source).Error; err != nil {
		responses.Error(w, http.StatusNotFound, "Site not found or access denied")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate New Keys
	//||------------------------------------------------------------------------------------------------||

	privKeyBytes, err := helpers.GenerateRandom()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate private key")
		return
	}

	pubKeyBytes, err := helpers.GenerateRandom()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to generate public key")
		return
	}

	privKey := hex.EncodeToString(privKeyBytes)
	pubKey := hex.EncodeToString(pubKeyBytes)

	//||------------------------------------------------------------------------------------------------||
	//|| Create New Copied Site
	//||------------------------------------------------------------------------------------------------||

	newSite := source
	newSite.IDSite = 0
	newSite.SiteStatus = "PNEW"
	newSite.SitePrivate = privKey
	newSite.SitePublic = pubKey

	if !strings.HasSuffix(newSite.SiteName, " - Copy") {
		newSite.SiteName += " - Copy"
	}

	if err := db.DB.Create(&newSite).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to create copied site")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return New Site ID
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"id": newSite.IDSite,
	})
}
