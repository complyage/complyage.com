package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct for Response
//||------------------------------------------------------------------------------------------------||

type VerificationStatusResponse struct {
	TypeCode     string `json:"vType" gorm:"column:type_code"`
	LatestStatus string `json:"vStatus" gorm:"column:latest_status"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerifications - Returns latest status of each verification type for logged-in user
//||------------------------------------------------------------------------------------------------||

func UserDashboard(w http.ResponseWriter, r *http.Request) {

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
	//|| Query Latest Verification Statuses
	//||------------------------------------------------------------------------------------------------||

	var results []VerificationStatusResponse
	query := `
		SELECT 
			vt.verification_code AS type_code,
			COALESCE(v.verification_status, 'MISS') AS latest_status
		FROM verification_types vt
			LEFT JOIN (
				SELECT verification_type, MAX(id_verification) AS max_id
				FROM verifications
				WHERE fid_account = ?
				GROUP BY verification_type
			) latest ON latest.verification_type = vt.verification_code
		LEFT JOIN verifications v ON v.id_verification = latest.max_id
		ORDER BY vt.id_verification_type;
    `

	if err := app.SQLDB["main"].DB.Raw(query, accountID).Scan(&results).Error; err != nil {
		fmt.Println("Failed to query verification statuses:", err)
		responses.Error(w, http.StatusInternalServerError, "Database query failed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"accountID":     accountID,
		"verifications": results,
	})
}
