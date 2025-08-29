package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/verify"
	"encoding/json"
	"net/http"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDProgressHandler
//||------------------------------------------------------------------------------------------------||

func TestingJSONVerification(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Request
	//||------------------------------------------------------------------------------------------------||

	identifier := r.URL.Query().Get("identifier")

	//||------------------------------------------------------------------------------------------------||
	//|| Validate UUID
	//||------------------------------------------------------------------------------------------------||

	if identifier == "" {
		responses.Error(w, http.StatusBadRequest, "Missing verification UUID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.AgentLoad(app.SQLDB["main"], app.Storages["verifications"], identifier)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	jsonBytes, err := json.Marshal(verifyRecord)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to marshal verification record -> "+err.Error())
		return
	}

	var jsonData map[string]interface{}
	err = json.Unmarshal(jsonBytes, &jsonData)
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to parse verification JSON -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   verifyRecord.UUID,
		"status": verifyRecord.Status,
		"json":   jsonData,
	})
}
