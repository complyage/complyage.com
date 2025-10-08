package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/db/abstract"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/base/encrypt"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Envelope
//||------------------------------------------------------------------------------------------------||

type userPrivateResponse struct {
	LoggedIn bool `json:"loggedIn"`
	Level    int  `json:"level"`
	Private  bool `json:"private"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerifications - Location
//||------------------------------------------------------------------------------------------------||

func UserPrivateFetch(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Failed Response
	//||------------------------------------------------------------------------------------------------||

	response := userPrivateResponse{
		LoggedIn: false,
		Level:    0,
		Private:  false,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	_, account, _, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Success(w, http.StatusOK, response)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Keys
	//||------------------------------------------------------------------------------------------------||

	response.LoggedIn = true
	keys, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		responses.Success(w, http.StatusOK, response)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Not Level One
	//||------------------------------------------------------------------------------------------------||

	response.Level = keys.Level
	if keys.Level > 1 {
		privateKey, err := abstract.StoredPrivateKey(r)
		if err != nil {
			responses.Error(w, http.StatusFailedDependency, app.Err("API").Code("BAD_CACHE"))
			return
		}
		keys.Private = privateKey
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Request
	//||------------------------------------------------------------------------------------------------||

	err = encrypt.CheckPrivateKey(keys.Private, keys.CheckKey)
	if err != nil {
		responses.Success(w, http.StatusOK, response)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	response.Private = true
	responses.Success(w, http.StatusOK, response)

}
