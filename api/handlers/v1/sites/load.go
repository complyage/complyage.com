package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func SitesLoadHandler(w http.ResponseWriter, r *http.Request) {

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
	if err := app.SQLDB["main"].DB.
		Where("id_site = ? AND fid_account = ?", id, session.ID).
		Where("site_status NOT IN ('RMVD', 'BNND')").
		First(&site).Error; err != nil {
		responses.Success(w, http.StatusOK, map[string]any{
			"success": false,
			"site":    nil,
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create Logo Hash
	//||------------------------------------------------------------------------------------------------||

	host := os.Getenv("VITE_COMPLYAGE_MINIO_URL")
	bucket := os.Getenv("VITE_MINIO_BUCKET")
	hashSecret := os.Getenv("MINIO_HASH")
	hashInput := fmt.Sprintf("%s%d%s", hashSecret, site.ID, hashSecret)
	hashBytes := sha256.Sum256([]byte(hashInput))
	hashHex := fmt.Sprintf("%d_%s", site.ID, hex.EncodeToString(hashBytes[:]))
	logoURL := fmt.Sprintf("%s/%s/sites/logos/%s.webp", host, bucket, hashHex)

	//||------------------------------------------------------------------------------------------------||
	//|| Check if Logo exists
	//||------------------------------------------------------------------------------------------------||

	req, err := http.NewRequest("HEAD", logoURL, nil)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to construct request")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Definitely not the best way to do this, but feeling lazy
	//||------------------------------------------------------------------------------------------------||

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil || resp.StatusCode != http.StatusOK {
		responses.Success(w, http.StatusOK, map[string]any{
			"site":    site,
			"hash":    hashHex,
			"missing": true,
		})
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Site
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"site":    site,
		"hash":    hashHex,
		"missing": false,
	})
}
