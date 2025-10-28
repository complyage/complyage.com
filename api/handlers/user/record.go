package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/identity"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct for Response
//||------------------------------------------------------------------------------------------------||

type userRecordData struct {
	ID         int64
	Username   string
	Identifier string
	Status     string
	Level      int
	Identity   identity.Identity
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerifications - Returns latest status of each verification type for logged-in user
//||------------------------------------------------------------------------------------------------||

func UserRecord(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	var response userRecordData

	//||------------------------------------------------------------------------------------------------||
	//|| Check Account
	//||------------------------------------------------------------------------------------------------||

	_, account, session, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	identity, err := identity.Load(account.ID)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get account identity: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	response.ID = account.ID
	response.Username = account.Username
	response.Identifier = session.Identifier
	response.Status = account.Status
	response.Level = session.Level
	response.Identity = identity

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, response)
}
