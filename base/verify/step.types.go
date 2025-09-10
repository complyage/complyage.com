package verify

import (
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Load Step New Type
//||------------------------------------------------------------------------------------------------||

func initStepTypes() {

	app.Constants("VERIFY_STEP_TYPES").AddString("MODERATE", "Moderation step")
	app.Constants("VERIFY_STEP_TYPES").AddString("APPROVAL", "Verification Approved")
	app.Constants("VERIFY_STEP_TYPES").AddString("REJECTED", "Verification Rejected")
	app.Constants("VERIFY_STEP_TYPES").AddString("EXPIRED", "Verification Expired")

	app.Constants("VERIFY_STEP_TYPES").AddString("SENT_EMAIL", "Email sent to user")
	app.Constants("VERIFY_STEP_TYPES").AddString("SENT_SMS", "SMS sent to user")
	app.Constants("VERIFY_STEP_TYPES").AddString("PAY_PROCESS", "Payment processed")
	app.Constants("VERIFY_STEP_TYPES").AddString("CODE_ENTRY", "Code entry verification")

	app.Constants("VERIFY_STEP_TYPES").AddString("FACE_MATCH", "Face matching verification")
	app.Constants("VERIFY_STEP_TYPES").AddString("FACE_LIVE", "Live face verification")
	app.Constants("VERIFY_STEP_TYPES").AddString("UPLOAD", "Media uploaded")
	app.Constants("VERIFY_STEP_TYPES").AddString("OCR", "Optical Character Recognition performed")
	app.Constants("VERIFY_STEP_TYPES").AddString("DOB_EXTRACT", "DOB was extracted")
	app.Constants("VERIFY_STEP_TYPES").AddString("DOB", "DOB verified")
	app.Constants("VERIFY_STEP_TYPES").AddString("AGE_VERIFIED", "Age verified")
	app.Constants("VERIFY_STEP_TYPES").AddString("COMPLETE", "Verification process completed")

	//||------------------------------------------------------------------------------------------------||
	//|| General Actions
	//||------------------------------------------------------------------------------------------------||

	app.Constants("VERIFY_STEP_TYPES").AddString("INITIAL", "Verification Process Initialized...")
	app.Constants("VERIFY_STEP_TYPES").AddString("PANIC_REJECT", "Verification Suffered an unrecoverable error and was Rejected")

	//||------------------------------------------------------------------------------------------------||
	//|| Status
	//||------------------------------------------------------------------------------------------------||

	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_PEVF", "Status changed to Pending Verification")
	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_INPR", "Status changed to In Progress")
	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_ESCL", "Status changed to Escalated")
	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_VERF", "Status changed to Verified")
	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_RJCT", "Status changed to Rejected")
	app.Constants("VERIFY_STEP_TYPES").AddString("STATUS_EXPD", "Status changed to Expired")

	//||------------------------------------------------------------------------------------------------||
	//|| General Queue
	//||------------------------------------------------------------------------------------------------||

	app.Constants("VERIFY_STEP_TYPES").AddString("QUEUED_L1", "Entered into Level 1 verification queue")
	app.Constants("VERIFY_STEP_TYPES").AddString("AGENT_L1", "Agent Level 1 processing")

	//||------------------------------------------------------------------------------------------------||
	//|| Facial
	//||------------------------------------------------------------------------------------------------||

	app.Constants("VERIFY_STEP_TYPES").AddString("FACE_AGE", "Estimating facial age data")
	app.Constants("VERIFY_STEP_TYPES").AddString("FACE_AGE_EST", "Facial age estimated")
	app.Constants("VERIFY_STEP_TYPES").AddString("DOB_MATCH", "Date of Birth matched")
	app.Constants("VERIFY_STEP_TYPES").AddString("DOB_MISMATCH", "Date of Birth did not match age estimate")

}
