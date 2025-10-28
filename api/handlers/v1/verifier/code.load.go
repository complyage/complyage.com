package verifier

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request Response
//||------------------------------------------------------------------------------------------------||

type codeLoadResponse struct {
	Identifier string            `json:"identifier"`
	Status     verify.StatusType `json:"status"`
	Type       types.DataType    `json:"type"`
	Details    string            `json:"details"`
	Message    string            `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Card Code Verification Attempt
//||------------------------------------------------------------------------------------------------||

func VerificationCodeLoad(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Code Load")

	//||------------------------------------------------------------------------------------------------||
	//|| Get Params
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

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
	//|| Check
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load
	//||------------------------------------------------------------------------------------------------||

	app.Log.Data(fmt.Sprintf("2FA Code Check : %s - %s [%s]", verifyRecord.Type.String(), verifyRecord.UUID, verifyRecord.TwoFactor.Code))

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, codeLoadResponse{
		Identifier: identifier,
		Status:     verifyRecord.Status,
		Type:       verifyRecord.Type,
		Details:    "Verification code is valid",
	})
}
