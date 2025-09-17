package abstract

import (
	"net/http"
	"time"

	"github.com/complyage/base/identity"
	"github.com/ralphferrara/aria/auth/types"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Compile the Full Auth Me
//||------------------------------------------------------------------------------------------------||

type authMeResponse struct {
	ID         int64
	Identifier string
	Username   string
	Status     string
	Type       string
	Level      int
	Created    time.Time
	LastLogin  time.Time
	Identity   identity.Identity `json:"identity"`
}

//||------------------------------------------------------------------------------------------------||
//|| Auth Me Function
//||------------------------------------------------------------------------------------------------||

func AuthMe(w http.ResponseWriter, r *http.Request, authMe types.AuthMeRecord) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Identity
	//||------------------------------------------------------------------------------------------------||

	identity, err := identity.Load(authMe.ID)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Auth Me
	//||------------------------------------------------------------------------------------------------||

	me := authMeResponse{
		ID:         authMe.ID,
		Identifier: authMe.Identifier,
		Username:   authMe.Username,
		Status:     authMe.Status,
		Type:       authMe.Type,
		Level:      authMe.Level,
		Identity:   identity,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Handle the Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, me)
	return nil
}
