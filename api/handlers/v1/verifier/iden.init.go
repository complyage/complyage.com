package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type idenInitResponse struct {
	Identifier string `json:"identifier"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func IdentifierVerifyInitHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: IDEN Init")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.AccountCheckLogin(r, true, 1)
	if err != nil {
		app.Log.Info(err.Error())
		responses.Error(w, http.StatusUnauthorized, app.Err("API").Code("NO_SESSION"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord := verify.Create(types.DataTypeIDEN, account)
	verifyRecord.Save()
	verifyRecord.DatabaseInsert()

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, idenInitResponse{
		Identifier: verifyRecord.UUID,
	})

}
