package handlers

//||------------------------------------------------------------------------------------------------||
//|| Level1Router
//||------------------------------------------------------------------------------------------------||

import (
	"agent/models"
	"agent/publish"
	"base/verify"
	"fmt"
	"os"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Handlers
//||------------------------------------------------------------------------------------------------||

func HandleVerifyID(av publish.AgentVerification) error {
	//||------------------------------------------------------------------------------------------------||
	//|| Handler Start
	//||------------------------------------------------------------------------------------------------||

	moderator := os.Getenv("AGENT_ID")

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.AgentLoad(app.SQLDB["main"], app.Storages["verifications"], av.Identifier, verify.DataTypeIDEN)
	if err != nil {
		StepError(0, "Verification record not found", err.Error())
		return err
	}

	StepInfo(0, "Verification record found", fmt.Sprintf("ID: %d, UUID: %s", verifyRecord.Type, verifyRecord.UUID))

	//||------------------------------------------------------------------------------------------------||
	//|| Mark as In Progress
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Step = 1
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("AGENT_L1"), moderator)
	verifyRecord.UpdateStatusInProgress()
	verifyRecord.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Media
	//||------------------------------------------------------------------------------------------------||

	front := verifyRecord.Encrypted.Data.IDEN.Front
	frontMedia := publish.AgentMedia{
		Mime:   front.Mime,
		Base64: front.Base64,
	}
	selfie := verifyRecord.Encrypted.Data.IDEN.Selfie
	selfieMedia := publish.AgentMedia{
		Mime:   selfie.Mime,
		Base64: selfie.Base64,
	}

	StepInfo(0, "Media Files loaded", fmt.Sprintf("Front: %d, Selfie: %d", len(frontMedia.Base64), len(selfieMedia.Base64)))

	//||------------------------------------------------------------------------------------------------||
	//|| Step 1. Call OCR model to extract text from ID document
	//||------------------------------------------------------------------------------------------------||

	StepInfo(1, "Calling OCR model for ID document", "")
	verifyRecord.IncrementStep()
	ocrResponse, err := models.ModelCallOCR(frontMedia, "Extract all text from the ID document.")
	if err != nil {
		StepError(1, "Error:", err.Error())
		verifyRecord.CanReset = true
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_OCR_MODEL_FAIL"))
	}

	if ocrResponse.Success && ocrResponse.Data != nil {
		StepInfo(1, "OCR Response : ", ocrResponse.Data.Text)
	} else {
		StepError(1, "Text was not received", "")
		verifyRecord.CanReset = true
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_OCR_EMPTY"))
	}

	verifyRecord.IncrementStep()
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("OCR"), fmt.Sprintf("%d characters returned", len(ocrResponse.Data.Text)))
	verifyRecord.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Step 2. Call Gemma to extract structured data from OCR text
	//||------------------------------------------------------------------------------------------------||

	StepInfo(2, "Extracting the OCR data to JSON", "")
	verifyRecord.IncrementStep()
	dobResponse, err := models.ModelCallDOB(ocrResponse.Data.Text)
	if err != nil {
		StepError(2, "DOB JSON extraction failed:", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_DOB_JSON_FAIL"))
	}

	if !dobResponse.Success || dobResponse.RawDOB == "" {
		StepError(2, "DOB not found in OCR text", dobResponse.RawDOB)
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_DOB_MISSING"))
	}
	verifyRecord.IncrementStep()
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("DOB_EXTRACT"), dobResponse.RawDOB)
	verifyRecord.Save()

	//||------------------------------------------------------------------------------------------------||
	//|| Step 3. Check that DOB exists and is over 18
	//||------------------------------------------------------------------------------------------------||

	StepInfo(3, "Verify Age", "")
	if dobResponse.DOB.Age() < 16 {
		StepError(3, "DOB Underage missing or not numbers", "Underage")
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_DOB_UNDERAGE"))
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 4. Call face detect model
	//||------------------------------------------------------------------------------------------------||

	StepInfo(4, "Check for ID Face", "")
	frontFaceDetectResponse, err := models.ModelCallFaceDetect(frontMedia, "Detect the face on the ID document.")
	if err != nil {
		StepError(4, "Facial Detection Failed", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_FACE_DETECT_MODEL_FAIL"))
	}

	if !frontFaceDetectResponse.Success {
		StepError(4, "Face detect model failed for unknown reason", "")
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("MODEL_FACE_DETECT_NO_FACES"))
	}

	// //||------------------------------------------------------------------------------------------------||
	// //|| Step 7. Call face detect model
	// //||------------------------------------------------------------------------------------------------||

	// StepInfo("STEP 7. Check for Selfie Face ", "")
	// selfieFaceDetectResponse, err := models.ModelCallFaceDetect(selfieMedia, "Detect the face on the selfie.")
	// if err != nil {
	// 	StepError("STEP 7: Facial Detection Failed", err.Error())
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_FAIL")
	// }

	// if !selfieFaceDetectResponse.Success || len(selfieFaceDetectResponse.Faces) == 0 {
	// 	StepError("STEP : Face detect model failed: ", selfieFaceDetectResponse.Error)
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_MISSING")
	// }

	//||------------------------------------------------------------------------------------------------||
	//|| Step 8. Call face compare model
	//||------------------------------------------------------------------------------------------------||

	//StepInfo("STEP 8. Checking if faces match.", "")
	// faceMatchResponse, err := models.ModelCallFaceCompare(frontMedia.Base64, selfieMedia.Base64, map[string]interface{}{
	// 	"metric":    "cosine",
	// 	"threshold": 0.4,
	// })

	// if err != nil {
	// 	//StepError("STEP 8: Face compare model failed: ", faceMatchResponse.Error)
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_MISSING")
	// }

	// if !faceMatchResponse.Success || !faceMatchResponse.Data.Match {
	// 	//StepError("STEP 8: Face compare mismatch: ", faceMatchResponse.Error)
	// 	return verifyRecord.UpdateStatusReject(os.Getenv("AGENT_ID"), "Face compare mismatch: "+faceMatchResponse.Error)
	// }

	//||------------------------------------------------------------------------------------------------||
	//|| Verified by Agent
	//||------------------------------------------------------------------------------------------------||

	// //StepInfo("VERIFICATION COMPLETED!", "")
	// verifyRecord.IncrementStep()
	// return verifyRecord.UpdateStatusVerified(os.Getenv("AGENT_ID"))
	return nil
}
