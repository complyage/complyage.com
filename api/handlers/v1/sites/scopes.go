package sites

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/scopes"
	"github.com/ralphferrara/aria/responses"
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

func ScopesListHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Types
	//||------------------------------------------------------------------------------------------------||

	types := []VerificationTypeResponse{}
	var i int
	for s := range scopes.ScopesList {
		i++
		types = append(types, VerificationTypeResponse{
			ID:          uint(i),
			Code:        scopes.ScopesList[s].Code,
			Description: scopes.ScopesList[s].Description,
			Level:       uint8(scopes.ScopesList[s].Level),
		})
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, types)
}
