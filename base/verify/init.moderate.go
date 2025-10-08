package verify

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Errors
//||------------------------------------------------------------------------------------------------||

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| Moderate Status
	//||------------------------------------------------------------------------------------------------||
	app.Constants("ModerateStatus").AddCode("Pending", "PEND", "Pending Moderation")
	app.Constants("ModerateStatus").AddCode("Approved", "APRV", "Approved by Moderator")
	app.Constants("ModerateStatus").AddCode("PendingVerification", "PEVF", "Pending Verification")
	app.Constants("ModerateStatus").AddCode("Rejected", "RJCT", "Denied by Moderator")
	//||------------------------------------------------------------------------------------------------||
	//|| Moderate Type
	//||------------------------------------------------------------------------------------------------||
	app.Constants("ModerateType").AddCode("TwoFactor", "2FCT", "Two Factor Authentication")
	app.Constants("ModerateType").AddCode("Level1", "LVL1", "AI Agent Level 1")
	app.Constants("ModerateType").AddCode("Level2", "LVL2", "AI Agent Level 2")
	app.Constants("ModerateType").AddCode("Level3", "LVL3", "Human Authentication")
}
