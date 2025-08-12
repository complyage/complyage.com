package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/responses"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct for Full Verification Response
//||------------------------------------------------------------------------------------------------||

type UserVerificationFull struct {
	ID        int64     `json:"id" gorm:"column:id_verification"`
	Type      string    `json:"type" gorm:"column:verification_type"`
	Status    string    `json:"status" gorm:"column:verification_status"`
	Meta      string    `json:"meta" gorm:"column:verification_meta"`
	CreatedAt time.Time `json:"created" gorm:"column:verification_created"`
	UpdatedAt time.Time `json:"updated" gorm:"column:verification_updated"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerificationsFull - Lists all verifications with paging support
//||------------------------------------------------------------------------------------------------||

func UserVerifications(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
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

	accountID := session.ID

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Paging Parameters
	//||------------------------------------------------------------------------------------------------||

	pageParam := r.URL.Query().Get("page")
	limitParam := r.URL.Query().Get("limit")

	page, err := strconv.Atoi(pageParam)
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(limitParam)
	if err != nil || limit < 1 || limit > 100 {
		limit = 20 // default limit
	}

	offset := (page - 1) * limit

	//||------------------------------------------------------------------------------------------------||
	//|| Query with Paging
	//||------------------------------------------------------------------------------------------------||

	var verifications []UserVerificationFull
	query := `
		SELECT 
			id_verification,
			verification_type,
			verification_status,
			verification_meta,
			created_at,
			updated_at
		FROM verifications
		WHERE fid_account = ?
		ORDER BY created_at DESC
		LIMIT ? OFFSET ?;
	`

	if err := db.DB.Raw(query, accountID, limit, offset).Scan(&verifications).Error; err != nil {
		fmt.Println("Failed to fetch user verifications:", err)
		responses.Error(w, http.StatusInternalServerError, "Database query failed")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Total Count for Pagination
	//||------------------------------------------------------------------------------------------------||

	var total int64
	if err := db.DB.Raw(`SELECT COUNT(*) FROM verifications WHERE fid_account = ?`, accountID).Scan(&total).Error; err != nil {
		fmt.Println("Failed to count verifications:", err)
		total = int64(len(verifications)) // fallback
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"accountID":     accountID,
		"page":          page,
		"limit":         limit,
		"total":         total,
		"verifications": verifications,
	})
}
