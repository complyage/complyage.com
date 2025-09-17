package verifier

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/verify"

	"github.com/complyage/complyagent.com/publish"

	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
)

//||------------------------------------------------------------------------------------------------||
//|| VerifyIDProgressHandler
//||------------------------------------------------------------------------------------------------||

func TestingResetVerification(w http.ResponseWriter, r *http.Request) {

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
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch a Session
	//||------------------------------------------------------------------------------------------------||

	session, err := actions.FetchSession(cookie.Value)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification Record (with account public for decrypt)
	//||------------------------------------------------------------------------------------------------||

	account, err := abstract.GetAccountByVerificationUUID(identifier)
	if err != nil || account == nil {
		responses.Error(w, http.StatusBadRequest, "Account not found for verification")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Session Account
	//||------------------------------------------------------------------------------------------------||

	if session.ID != account.ID {
		responses.Error(w, http.StatusBadRequest, "Session does not match account")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt
	//||------------------------------------------------------------------------------------------------||

	encrypt, err := abstract.GetKeyByAccount(uint(account.ID))
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to get encryption keys")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.Load(app.SQLDB["main"], app.Storages["verifications"], identifier, encrypt.Private, encrypt.Public)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Verification record not found -> "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Steps
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Steps = []verify.Step{}
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("AGENT_L1"), "")

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Status to Pending Verification
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.UpdateStatusPendingVerification()

	//||------------------------------------------------------------------------------------------------||
	//|| Update the Verification Record
	//||------------------------------------------------------------------------------------------------||

	if verifyRecord.Type == verify.DataTypeIDEN {
		fmt.Println("Starting ID Verification Reset for: " + verifyRecord.UUID)
		insert := publish.AgentVerifyIDStart(verifyRecord)
		if insert != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
			return
		}
	} else {
		fmt.Println("Starting Facial Verification Reset for: " + verifyRecord.UUID)
		insert := publish.AgentVerifyFaceStart(verifyRecord)
		if insert != nil {
			responses.Error(w, http.StatusInternalServerError, "Failed to start ID verification: "+insert.Error())
			return
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Return Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"uuid":   verifyRecord.UUID,
		"status": verifyRecord.Status,
	})
}
