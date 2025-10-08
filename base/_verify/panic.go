package verify

import (
	"time"

	"github.com/ralphferrara/aria/app"
)

func panicReject(uuid string, verifyType DataType, err app.ErrorsEntry) Verification {
	//||------------------------------------------------------------------------------------------------||
	//|| Create a Verification
	//||------------------------------------------------------------------------------------------------||
	verification := Verification{
		CanReset:      false,
		ResetAttempts: 9999,
		Status:        STATUSES.Rejected,
		UUID:          uuid,
		Type:          verifyType,
		Timestamp:     time.Now().UTC(),
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Steps
	//||------------------------------------------------------------------------------------------------||
	verification.StepsInit()
	verification.Step = 9999
	verification.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("PANIC_REJECT"), err.Message)
	//||------------------------------------------------------------------------------------------------||
	//|| Create a Verification
	//||------------------------------------------------------------------------------------------------||
	verification.UpdateStatusReject("SYSTEM", app.Err("verify").Get("VERIFY_PANIC"))
	verification.Save()
	//||------------------------------------------------------------------------------------------------||
	//|| Create a Verification
	//||------------------------------------------------------------------------------------------------||
	return verification
}
