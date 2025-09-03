package handlers

//||------------------------------------------------------------------------------------------------||
//|| Level1Router
//||------------------------------------------------------------------------------------------------||

import (
	"agent/models"
	"agent/publish"
	"base/verify"
	"encoding/json"
	"fmt"
	"os"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Facial Struct
//||------------------------------------------------------------------------------------------------||

type Facial struct {
	DOB      verify.DOB   `json:"dob,omitempty"`
	DOBMatch bool         `json:"dob_match,omitempty"`
	Selfie   verify.Media `json:"selfie,omitempty"`
	Age      int          `json:"age,omitempty"`
	Min      int          `json:"min,omitempty"`
	Max      int          `json:"max,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handlers
//||------------------------------------------------------------------------------------------------||

func HandleFacialAge(av publish.AgentVerification) error {
	//||------------------------------------------------------------------------------------------------||
	//|| Handler Start
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("Handling VERIFY_FACE for %s (level %d)\n", av.Identifier, av.Level)
	moderator := os.Getenv("AGENT_ID")

	//||------------------------------------------------------------------------------------------------||
	//|| Verification Record matches Account
	//||------------------------------------------------------------------------------------------------||

	verifyRecord, err := verify.AgentLoad(app.SQLDB["main"], app.Storages["verifications"], av.Identifier)
	if err != nil {
		StepError("Verification record not found", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, "Verification record not found")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Mark as In Progress
	//||------------------------------------------------------------------------------------------------||

	verifyRecord.Step = 1
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

	//||------------------------------------------------------------------------------------------------||
	//|| Step 1. Call Facial Age Model
	//||------------------------------------------------------------------------------------------------||

	StepInfo("STEP 1. Calling Facial Age model for Selfie", "")
	verifyRecord.IncrementStep()
	faceResponse, err := models.ModelCallFaceDetect(selfieMedia, "Explain these faces age in years.")
	if err != nil {
		StepError("STEP 1: ", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, "RJT_FACE_SERVICE_FAIL")
	}

	// Pretty-print faceResponse JSON to console
	if jsonBytes, err := json.MarshalIndent(faceResponse, "", "   "); err == nil {
		fmt.Printf("FaceResponse:\n%s\n", string(jsonBytes))
	} else {
		fmt.Println("Failed to pretty print FaceResponse:", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Response
	//||------------------------------------------------------------------------------------------------||

	if !faceResponse.Success || len(faceResponse.Faces) == 0 {
		StepError("STEP 1: Face not detected", "")
		return verifyRecord.UpdateStatusReject(moderator, "RJT_FACE_NOT_FOUND")
	}
	if len(faceResponse.Faces) != 1 {
		StepError("STEP 2: Multiple Faces Detected", "")
		return verifyRecord.UpdateStatusReject(moderator, "RJT_TOO_MANY_FACES")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 3. Fill Facial Struct, Verify Confidence & Age Range
	//||------------------------------------------------------------------------------------------------||

	face := faceResponse.Faces[0]
	conf := 0.0
	if face.Confidence != nil {
		conf = *face.Confidence
	}
	min, max := 0, 0
	if face.AgeMin != nil {
		min = *face.AgeMin
	}
	if face.AgeMax != nil {
		max = *face.AgeMax
	}

	facial := Facial{
		DOB:    dob,
		Selfie: selfie,
		Age:    age,
		Min:    min,
		Max:    max,
	}
	if conf >= 0.66 {
		facial.DOBMatch = age >= min && age <= max
	} else {
		facial.DOBMatch = false
	}

	// Pretty-print the Facial struct to console
	if jsonBytes, err := json.MarshalIndent(facial, "", "   "); err == nil {
		fmt.Printf("FacialStruct:\n%s\n", string(jsonBytes))
	}

	if conf < 0.66 {
		StepError("STEP 3: Confidence too low", "")
		return verifyRecord.UpdateStatusReject(moderator, "RJT_FACE_CONFIDENCE_LOW")
	}
	if !facial.DOBMatch {
		StepError("STEP 3: Age not within predicted range", "")
		return verifyRecord.UpdateStatusVerified(moderator)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verified by Agent
	//||------------------------------------------------------------------------------------------------||

	StepInfo("VERIFICATION COMPLETED!", "")
	verifyRecord.IncrementStep()
	return verifyRecord.UpdateStatusVerified(os.Getenv("AGENT_ID"))
}
