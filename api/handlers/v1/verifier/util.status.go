package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

type VerifyStatusResponse struct {
	UUID   string        `json:"uuid"`
	Status string        `json:"status"`
	Type   string        `json:"type"`
	Step   int           `json:"step"`
	Steps  []verify.Step `json:"steps"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: VerifyIDStatusHandler
//||------------------------------------------------------------------------------------------------||

func VerifyStatusHandler(w http.ResponseWriter, r *http.Request) {

	app.Log.Info("Handler: Veridy Status")

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Query Params
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
	//|| Load Verification Record
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.CheckLoad(identifier, account.ID)
	if err != nil {
		app.Log.Error("Failed to load verification record: ", err.Error())
		responses.Error(w, http.StatusBadRequest, app.Err("Verify").Code("VERIFY_LOAD_UUID"))
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Minimal Status Object
	//||------------------------------------------------------------------------------------------------||

	resp := VerifyStatusResponse{
		UUID:   verifyRecord.UUID,
		Status: verifyRecord.Status.String(),
		Type:   verifyRecord.Type.String(),
		Step:   verifyRecord.Step,
		Steps:  verifyRecord.Steps,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, resp)
}
