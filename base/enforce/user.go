package enforce

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/identity"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/auth/db"
)

//||------------------------------------------------------------------------------------------------||
//|| User
//||------------------------------------------------------------------------------------------------||

type User struct {
	ID       int64             `json:"id"`
	Username string            `json:"username"`
	Level    int               `json:"level"`
	Status   string            `json:"status"`
	Identity identity.Identity `json:"identity"`
	KeyLevel int               `json:"keyLevel"`
	Public   string            `json:"public"`
	Private  string            `json:"private"`
	CheckKey string            `json:"checkKey"`
}

//||------------------------------------------------------------------------------------------------||
//|| Load User
//||------------------------------------------------------------------------------------------------||

func LoadUser(r *http.Request) User {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		return User{}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Session
	//||------------------------------------------------------------------------------------------------||

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		return User{}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Database Account
	//||------------------------------------------------------------------------------------------------||

	account, err := db.GetAccountByID(fmt.Sprintf("%d", session.ID))
	if err != nil {
		return User{}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Identity
	//||------------------------------------------------------------------------------------------------||

	identity, err := identity.Load(account.ID)
	if err != nil {
		return User{}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Response Record
	//||------------------------------------------------------------------------------------------------||

	user := User{}
	user.ID = account.ID
	user.Username = account.Username
	user.Level = account.Level
	user.Status = account.Status
	user.Identity = identity
	user.KeyLevel = 0
	user.Public = ""
	user.Private = ""
	user.CheckKey = ""

	//||------------------------------------------------------------------------------------------------||
	//|| Get Keys
	//||------------------------------------------------------------------------------------------------||

	keys, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		return user
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Keys
	//||------------------------------------------------------------------------------------------------||

	user.KeyLevel = keys.Level
	user.Public = keys.Public
	user.Private = keys.Private
	user.CheckKey = keys.CheckKey

	//||------------------------------------------------------------------------------------------------||
	//|| Private Key Stored, not stored in database
	//||------------------------------------------------------------------------------------------------||

	if user.KeyLevel > 1 {
		user.Private = ""
		pookie := GetStoredPrivateKey(r)
		if pookie == "" {
			return user
		}
		user.Private = pookie
	}

	//||------------------------------------------------------------------------------------------------||
	//|| User
	//||------------------------------------------------------------------------------------------------||
	return user
}
