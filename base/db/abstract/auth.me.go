package abstract

import (
	"fmt"
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
	ID         int64             `json:"id"`
	Identifier string            `json:"identifier"`
	Username   string            `json:"username"`
	Status     string            `json:"status"`
	Type       string            `json:"type"`
	Level      int               `json:"level"`
	Created    time.Time         `json:"created"`
	LastLogin  time.Time         `json:"last_login"`
	Identity   identity.Identity `json:"identity"`
}

//||------------------------------------------------------------------------------------------------||
//|| Auth Me Function
//||------------------------------------------------------------------------------------------------||

func AuthMe(w http.ResponseWriter, r *http.Request, authMe types.AuthMeRecord) error {

	fmt.Println("AuthMe Invoked for Account ID:", authMe.ID)

	//||------------------------------------------------------------------------------------------------||
	//|| Check
	//||------------------------------------------------------------------------------------------------||
	//|| Get Identity
	//||------------------------------------------------------------------------------------------------||

	iden, err := identity.Load(authMe.ID)
	if err != nil {
		identity.Create(authMe.ID)
		iden, err = identity.Load(authMe.ID)
		if err != nil {
			fmt.Println("Could not load identity for account ID:", authMe.ID, "error:", err.Error())
			return err
		}
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
		Identity:   iden,
	}

	fmt.Println("AuthMe Compiled for Account ID:", me.ID)

	//||------------------------------------------------------------------------------------------------||
	//|| Handle the Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, me)
	return nil
}
