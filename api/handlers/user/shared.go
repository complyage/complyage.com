package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"
	"time"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct to return combined shared data
//||------------------------------------------------------------------------------------------------||

type SharedItemResponse struct {
	IDShared           uint      `json:"id_shared" gorm:"column:id_shared"`
	FIDSite            uint      `json:"fid_site" gorm:"column:fid_site"`
	FIDVerification    uint      `json:"fid_verification" gorm:"column:fid_verification"`
	SharedTimestamp    time.Time `json:"shared_timestamp" gorm:"column:shared_timestamp"`
	SiteName           string    `json:"site_name" gorm:"column:site_name"`
	SiteURL            string    `json:"site_url" gorm:"column:site_url"`
	VerificationType   string    `json:"verification_type" gorm:"column:verification_type"`
	VerificationStatus string    `json:"verification_status" gorm:"column:verification_status"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserSharedHandler - Returns user's shared verifications grouped by site
//||------------------------------------------------------------------------------------------------||

func UserSharedHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
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

	accountID := session.ID

	//||------------------------------------------------------------------------------------------------||
	//|| Query to fetch shared data joined with sites and verifications
	//||------------------------------------------------------------------------------------------------||

	var sharedItems []SharedItemResponse
	query := `
        SELECT 
            s.id_shared,
            s.fid_site,
            s.fid_verification,
            s.shared_timestamp,
            st.site_name,
            st.site_url,
            v.verification_type,
            v.verification_status
        FROM shared s
        JOIN verifications v ON v.id_verification = s.fid_verification
        JOIN sites st ON st.id_site = s.fid_site
        WHERE v.fid_account = ?
        ORDER BY st.site_name, s.shared_timestamp DESC
    `

	if err := app.SQLDB["main"].DB.Raw(query, accountID).Scan(&sharedItems).Error; err != nil {
		fmt.Println("❌ Failed to fetch shared data:", err)
		responses.Error(w, http.StatusInternalServerError, "Database query failed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"accountID": accountID,
		"shared":    sharedItems,
	})
}
