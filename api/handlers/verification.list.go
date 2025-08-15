//||------------------------------------------------------------------------------------------------||
//|| Handler: GetVerifications
//|| Endpoint: /api/verification/list
//|| Description: Returns all existing verifications for current user (uuid, type, meta, status, created, updated)
//||------------------------------------------------------------------------------------------------||

package handlers

import (
	"base/db"
	"base/helpers"
	"base/responses"
	"net/http"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| VerificationRow
//||------------------------------------------------------------------------------------------------||

type VerificationRow struct {
	UUID      string    `json:"uuid"      gorm:"column:verification_uuid;type:varchar(64);uniqueIndex"`
	Display   string    `json:"display"   gorm:"column:verification_display;type:varchar(64)"`
	Type      string    `json:"type"      gorm:"column:verification_type;type:varchar(4)"`
	Status    string    `json:"status"    gorm:"column:verification_status;type:varchar(4)"`
	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}

//||------------------------------------------------------------------------------------------------||
//|| Get Verifications List
//||------------------------------------------------------------------------------------------------||

func GetVerificationsList(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Session
	//||------------------------------------------------------------------------------------------------||

	session, err := helpers.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Query
	//||------------------------------------------------------------------------------------------------||

	q := `
		SELECT
			verification_display,
			verification_uuid,
			verification_type,
			verification_meta,
			verification_status,
			created_at,
			updated_at
		FROM verifications
		WHERE fid_account = ?
		`
	args := []any{session.ID}

	// Optional: filter by type
	if vtype := strings.ToUpper(r.URL.Query().Get("type")); vtype != "" {
		q += " AND verification_type = ?"
		args = append(args, vtype)
	}
	q += ` AND verification_status IN('APPR', 'PEVF', 'VERF')
		ORDER BY created_at DESC
		LIMIT 100`

	//||------------------------------------------------------------------------------------------------||
	//|| Query DB
	//||------------------------------------------------------------------------------------------------||

	var results []VerificationRow
	if err := db.DB.Raw(q, args...).Scan(&results).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Database error")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return as JSON
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, results)
}
