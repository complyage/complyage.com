package verify

import "github.com/ralphferrara/aria/app"

func init() {
	//||------------------------------------------------------------------------------------------------||
	//|| Initialize Steps
	//||------------------------------------------------------------------------------------------------||
	lib := app.Constants("VerifyStatus")
	lib.AddCode("Pending", "PEND", "Pending")
	lib.AddCode("PendingVerification", "PEVF", "Pending Verification")
	lib.AddCode("InProgress", "INPR", "In Progress")
	lib.AddCode("Escalaed", "ESCL", "Escalated")
	lib.AddCode("Verified", "VERF", "Verified")
	lib.AddCode("Rejected", "RJCT", "Rejected")
	lib.AddCode("Expired", "EXPD", "Expired")
	//||------------------------------------------------------------------------------------------------||
	//|| Verification Types
	//||------------------------------------------------------------------------------------------------||
	lib = app.Constants("VerifyType")
	lib.AddCode("Email", "MAIL", "Email Verification")
	lib.AddCode("Phone", "PHNE", "Phone Verification")
	lib.AddCode("Address", "ADDR", "Address Verification")
	lib.AddCode("CreditCard", "CRCD", "Credit Card Verification")
	lib.AddCode("ID Card", "IDEN", "IDCARD Verification")
	lib.AddCode("Face", "FACE", "Facial Age Analyis")
	lib.AddCode("Username", "USER", "Username Verification")
	//||------------------------------------------------------------------------------------------------||
	//|| Transaction Types
	//||------------------------------------------------------------------------------------------------||
	lib = app.Constants("TransactionStatus")
	lib.AddCode("Pending", "PENDING", "Transaction is pending")
	lib.AddCode("Processing", "PROCESSING", "Transaction is processing")
	lib.AddCode("Approved", "APPROVED", "Transaction is approved")
	lib.AddCode("Refunded", "REFUNDED", "Transaction is refunded")
	lib.AddCode("Chargeback", "CHARGEBACK", "Transaction is chargeback")
	//||------------------------------------------------------------------------------------------------||
	//|| Moderation Types
	//||------------------------------------------------------------------------------------------------||
	lib = app.Constants("ModerateType")
	lib.AddCode("TwoFactor", "TWO_FACTOR", "Two Factor Authentication")
	lib.AddCode("AILevel1", "AI_LEVEL_1", "AI Level 1 Moderation")
	lib.AddCode("AILevel2", "AI_LEVEL_2", "AI Level 2 Moderation")
	lib.AddCode("Human", "HUMAN", "Human Moderation")
	//||------------------------------------------------------------------------------------------------||
	//|| Verify Steps
	//||------------------------------------------------------------------------------------------------||
	lib = app.Constants("VerifyStep")
	lib.AddCode("Initial", "INITIAL", "Initial Step")
	lib.AddCode("StatusChange", "STATUS_CHANGE", "Status Change")
	lib.AddCode("AgentLevel1", "AGENT_LEVEL_1", "Agent Level 1 Processing")
	lib.AddCode("Moderate", "MODERATE", "Moderation Step")
	lib.AddCode("SentEmail", "SENT_EMAIL", "Sent Email")
	lib.AddCode("SentSMS", "SENT_SMS", "Sent SMS")
	lib.AddCode("PayProcess", "PAY_PROCESS", "Payment Processed")
	lib.AddCode("CodeEntry", "CODE_ENTRY", "Code Entry")
	lib.AddCode("FaceMatch", "FACE_MATCH", "Face Match")
	lib.AddCode("FaceLive", "FACE_LIVE", "Face Live")
	lib.AddCode("MediaUpload", "MEDIA_UPLOAD", "Media Upload")
	lib.AddCode("OCR", "OCR", "Optical Character Recognition")
	lib.AddCode("DOB", "DOB", "Date of Birth Verified")
	lib.AddCode("AgeVerified", "AGE_VERIFIED", "Age Verified")
	lib.AddCode("Complete", "COMPLETE", "Verification Complete")
}
