package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/responses"
	"net/http"
)

// ||------------------------------------------------------------------------------------------------||
// || Response DTO
// ||------------------------------------------------------------------------------------------------||

type VerificationTypeResponse struct {
	ID          uint   `json:"id"          gorm:"column:id_verification_type"`
	Code        string `json:"code"        gorm:"column:verification_code"`
	Description string `json:"description" gorm:"column:verification_description"`
	Level       uint8  `json:"level"       gorm:"column:verification_level"`
}

// ||------------------------------------------------------------------------------------------------||
// || Handler
// ||------------------------------------------------------------------------------------------------||

func VerificationTypesListHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Types
	//||------------------------------------------------------------------------------------------------||
	var types []VerificationTypeResponse
	if err := db.DB.
		Table("verification_types").
		Select("id_verification_type", "verification_code", "verification_description", "verification_level").
		Order("verification_level ASC, verification_code ASC").
		Scan(&types).Error; err != nil {

		responses.Error(w, http.StatusInternalServerError, "Error loading verification types")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, map[string]any{
		"verification_types": types,
	})
}
