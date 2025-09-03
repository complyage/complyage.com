package verifier

import (
	"base/db/models"
	"base/verify"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

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

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Type from Query
	//||------------------------------------------------------------------------------------------------||

	listType := r.URL.Query().Get("type")
	if !verify.IsValidDataType(listType) {
		responses.Error(w, http.StatusBadRequest, "Invalid type")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Query
	//||------------------------------------------------------------------------------------------------||

	var results []models.Verify
	db := app.SQLDB["main"].DB
	db = db.Where("fid_account = ?", session.ID)
	db = db.Where("verify_status NOT IN ?", []string{"PEND", "EXPD"})
	db = db.Where("verify_type = ?", listType)
	db = db.Order("verify_created DESC")
	db = db.Limit(100)
	db = db.Find(&results)
	if db.Error != nil {
		responses.Error(w, http.StatusInternalServerError, "Database error : "+db.Error.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return as JSON
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, results)
}
