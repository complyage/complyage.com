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
	"fmt"
	"net/http"
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

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate Public/Private Keys
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
	//|| Create New Site Model
	//||------------------------------------------------------------------------------------------------||

	site := models.Site{
		FidAccount:      session.ID,
		SiteStatus:      "PNEW",
		SitePrivate:     privKey,
		SitePublic:      pubKey,
		SiteEnforcement: "ALLZ",
	}

	if err := db.DB.Create(&site).Error; err != nil {
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
