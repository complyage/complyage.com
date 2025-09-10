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

func HandleFacialAge(av publish.AgentVerification) error {
	//||------------------------------------------------------------------------------------------------||
	//|| Handler Start
	//||------------------------------------------------------------------------------------------------||

	moderator := os.Getenv("AGENT_ID")

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.AgentLoad(app.SQLDB["main"], app.Storages["verifications"], av.Identifier, verify.DataTypeFACE)
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
	//|| Media/DOB/Age
	//||------------------------------------------------------------------------------------------------||

	selfie := verifyRecord.Encrypted.Data.FACE.Selfie
	selfieMedia := publish.AgentMedia{
		Mime:   selfie.Mime,
		Base64: selfie.Base64,
	}
	dob := verifyRecord.Encrypted.Data.FACE.DOB
	age := AgeFromMDY(dob.Month, dob.Day, dob.Year)
	StepInfo(1, fmt.Sprintf("Provided [DOB: %s, AGE: %d]", dob.String(), age), "")

	//||------------------------------------------------------------------------------------------------||
	//|| Step 1. Call Facial Age Model
	//||------------------------------------------------------------------------------------------------||

	StepInfo(2, "Calling Facial Age model for Selfie", "")
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("FACE_AGE"), "")
	verifyRecord.IncrementStep()
	faceResponse, err := models.ModelCallFaceDetect(selfieMedia, "")
	if err != nil {
		StepError(2, "Failed Calling Age Model", err.Error())
		switch err.Error() {
		case "MODEL_FACE_DETECT_MARSHAL_FAIL":
			verifyRecord.CanReset = false
		case "MODEL_FACE_DETECT_MODEL_FAIL":
			verifyRecord.CanReset = true
			verifyRecord.ResetAttempts++
		case "MODEL_FACE_DETECT_RESPONSE_FAIL":
			verifyRecord.CanReset = true
			verifyRecord.ResetAttempts++
		case "MODEL_FACE_DETECT_NO_FACES":
			verifyRecord.CanReset = false
		default:
			verifyRecord.CanReset = true
			verifyRecord.ResetAttempts++
		}
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get(err.Error()))
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 2. Fill Facial Struct, Verify Confidence & Age Range
	//||------------------------------------------------------------------------------------------------||

	conf := 0.0
	if faceResponse.Confidence != nil {
		conf = *faceResponse.Confidence
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 2. Create Facial Object
	//||------------------------------------------------------------------------------------------------||

	facial := verify.Facial{
		DOB:    dob,
		Selfie: selfie,
		Age:    faceResponse.Age,
		Min:    faceResponse.AgeMin,
		Max:    faceResponse.AgeMax,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 3. Verify Confidence & Age Range
	//||------------------------------------------------------------------------------------------------||

	StepInfo(2, fmt.Sprintf("Facial Age: %d, Range: %d-%d, Confidence: %.2f", faceResponse.Age, faceResponse.AgeMin, faceResponse.AgeMax, conf), "")
	verifyRecord.IncrementStep()
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("FACE_AGE_EST"), fmt.Sprintf("%d", faceResponse.Age))

	//||------------------------------------------------------------------------------------------------||
	//|| Step 3. Verify Confidence & Age Range
	//||------------------------------------------------------------------------------------------------||

	if conf >= 0.66 {
		facial.DOBMatch = age >= faceResponse.AgeMin && age <= faceResponse.AgeMax
	} else {
		facial.DOBMatch = false
	}
	verifyRecord.IncrementStep()

	//||------------------------------------------------------------------------------------------------||
	//|| Step 3. Fill Facial Struct, Verify Confidence & Age Range
	//||------------------------------------------------------------------------------------------------||

	if conf < 0.66 {
		StepError(3, "Confidence too low", "")
		verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("FACE_AGE_EST"), fmt.Sprintf("%d", faceResponse.Age))
		return verifyRecord.UpdateStatusReject(moderator, app.Err("agent").Get("RJT_FACE_CONFIDENCE_LOW"))
	}

	if !facial.DOBMatch {
		StepError(3, "DOB not within predicted range", "")
		verifyRecord.IncrementStep()
		verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("DOB_MISMATCH"), fmt.Sprintf("%d <> [%d-%d]", age, faceResponse.AgeMin, faceResponse.AgeMax))
		verifyRecord.UpdateAge(verify.DataTypeFACE, faceResponse.Age, verifyRecord.UUID)
		return verifyRecord.UpdateStatusVerified(moderator)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verified by Agent
	//||------------------------------------------------------------------------------------------------||

	StepInfo(4, "VERIFICATION COMPLETED!", "")
	verifyRecord.Display = facial.Mask()
	verifyRecord.AddStep(app.Constants("VERIFY_STEP_TYPES").Get("DOB_MATCH"), "")
	verifyRecord.Encrypted.Data.FACE = facial
	verifyRecord.UpdateVerification(verify.DataTypeFACE, dob.Mask(), verifyRecord.UUID)
	verifyRecord.UpdateDOB(verify.DataTypeFACE, dob, verifyRecord.UUID)
	verifyRecord.IncrementStep()
	return verifyRecord.UpdateStatusVerified(os.Getenv("AGENT_ID"))
}
