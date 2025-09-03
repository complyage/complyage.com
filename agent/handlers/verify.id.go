package handlers

//||------------------------------------------------------------------------------------------------||
//|| Level1Router
//||------------------------------------------------------------------------------------------------||

import (
	"agent/models"
	"agent/publish"
	"agent/steps"
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
	//|| Handlers
	//||------------------------------------------------------------------------------------------------||

	fmt.Printf("Handling VERIFY_ID for %s (level %d)\n", av.Identifier, av.Level)
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

	//||------------------------------------------------------------------------------------------------||
	//|| Step 1. Call OCR model to extract text from ID document
	//||------------------------------------------------------------------------------------------------||

	StepInfo("STEP 1. Calling OCR model for ID document", "")
	verifyRecord.IncrementStep()
	ocrResponse, err := models.ModelCallOCR(frontMedia, "Extract all text from the ID document.")
	if err != nil {
		StepError("STEP 1: ", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, "RJT_OCR_SERVICE_FAIL")
	}

	if ocrResponse.Success && ocrResponse.Data != nil {
		StepInfo("STEP 1: ", ocrResponse.Data.Text)
	} else {
		StepError("STEP 1: Text was not received", "")
		return verifyRecord.UpdateStatusReject(moderator, "RJT_OCR_SERVICE_FAIL")
	}

	dob, err := steps.ExtractDOBFromOCR(ocrResponse.Data.Text)
	if err != nil {
		fmt.Println("DOB extraction failed:", err)
		return verifyRecord.UpdateStatusReject(moderator, "RJT_DOB_EXTRACTION_FAIL")
	}

	fmt.Println("-----------------------------")
	fmt.Printf("DOB extracted: %02d/%02d/%04d (certainty %.2f)\n", dob.Month, dob.Day, dob.Year, dob.Certainty)
	fmt.Println("-----------------------------")

	//||------------------------------------------------------------------------------------------------||
	//|| Step 2. Call Gemma to extract structured data from OCR text
	//||------------------------------------------------------------------------------------------------||

	// StepInfo("STEP 2. Extracting the OCR data to JSON", "")
	// verifyRecord.IncrementStep()
	// jsonResponse, err := models.ModelCallGemmaJSON(ocrResponse.Data.Text, jsonStructure)
	// if err != nil {
	// 	StepError("STEP 2: Gemma JSON extraction failed:", err.Error())
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_CONVERT_JSON_FAIL")
	// }

	// //||------------------------------------------------------------------------------------------------||
	// //|| Step 3. Convert the JSON response to a map
	// //||------------------------------------------------------------------------------------------------||

	// StepInfo("STEP 3. Parsing JSON response", "")
	// verifyRecord.IncrementStep()
	// parsedJSON, err := ParseJSON(jsonResponse.Data.Raw)
	// if err != nil {
	// 	StepError("STEP 3: Failed to Parse Gemma Response:", err.Error())
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_JSON_FAIL")
	// }
	// verifyRecord.AddStep(verify.StepOCR, verify.StepOCR.Description(""))

	//||------------------------------------------------------------------------------------------------||
	//|| Step 4. Check that DOB exists and is over 18
	//||------------------------------------------------------------------------------------------------||

	// StepInfo("STEP 4. Parse DOB", "")
	// dobMap, ok := parsedJSON["dob"].(map[string]interface{})
	// if !ok {
	// 	StepError("STEP 4: DOB field missing or invalid in JSON response", "")
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_DOB_INVALID")
	// }

	// dayF, okDay := dobMap["day"].(float64)
	// monthF, okMonth := dobMap["month"].(float64)
	// yearF, okYear := dobMap["year"].(float64)
	// if !okDay || !okMonth || !okYear {
	// 	StepError("STEP 4b: DOB fields missing or not numbers", "")
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_DOB_INVALID")
	// }

	// pretty, _ := json.MarshalIndent(parsedJSON, "", "  ")
	// verifyRecord.AddStep(verify.StepOCR, string(pretty))

	//||------------------------------------------------------------------------------------------------||
	//|| Step 5. Check DOB is over 18
	//||------------------------------------------------------------------------------------------------||

	// StepInfo("STEP 5. Checking if DOB > 18", "")
	// verifyRecord.IncrementStep()
	// day := int(dayF)
	// month := int(monthF)
	// year := int(yearF)
	// verifyRecord.AddStep(verify.StepAgeVerified, verify.StepAgeVerified.Description(fmt.Sprintf("%d-%d-%d-", year, month, day)))
	// age := AgeFromMDY(month, day, year)
	// if age < 18 {
	// 	StepError(fmt.Sprintf("STEP 5: User is underage (age = %d)", age), "")
	// 	return verifyRecord.UpdateStatusReject(moderator, "RJT_AGE_UNDER_18")
	// }
	// verifyRecord.AddStep(verify.StepAgeVerified, verify.StepAgeVerified.Description(fmt.Sprintf("%d", age)))

	//||------------------------------------------------------------------------------------------------||
	//|| Step 6. Call face detect model
	//||------------------------------------------------------------------------------------------------||

	StepInfo("STEP 6. Check for ID Face", "")
	frontFaceDetectResponse, err := models.ModelCallFaceDetect(frontMedia, "Detect the face on the ID document.")
	if err != nil {
		StepError("STEP 6: Facial Detection Failed", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, "RJT_FRONT_FACE_DETECT_FAIL")
	}

	if !frontFaceDetectResponse.Success || len(frontFaceDetectResponse.Faces) == 0 {
		StepError("STEP 6: Face detect model failed: ", frontFaceDetectResponse.Error)
		return verifyRecord.UpdateStatusReject(moderator, "RJT_FRONT_FACE_DETECT_MISSING")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 7. Call face detect model
	//||------------------------------------------------------------------------------------------------||

	StepInfo("STEP 7. Check for Selfie Face ", "")
	selfieFaceDetectResponse, err := models.ModelCallFaceDetect(selfieMedia, "Detect the face on the selfie.")
	if err != nil {
		StepError("STEP 7: Facial Detection Failed", err.Error())
		return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_FAIL")
	}

	if !selfieFaceDetectResponse.Success || len(selfieFaceDetectResponse.Faces) == 0 {
		StepError("STEP : Face detect model failed: ", selfieFaceDetectResponse.Error)
		return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_MISSING")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Step 8. Call face compare model
	//||------------------------------------------------------------------------------------------------||

	StepInfo("STEP 8. Checking if faces match.", "")
	faceMatchResponse, err := models.ModelCallFaceCompare(frontMedia.Base64, selfieMedia.Base64, map[string]interface{}{
		"metric":    "cosine",
		"threshold": 0.4,
	})

	if err != nil {
		StepError("STEP 8: Face compare model failed: ", faceMatchResponse.Error)
		return verifyRecord.UpdateStatusReject(moderator, "RJT_SELFIE_FACE_DETECT_MISSING")
	}

	if !faceMatchResponse.Success || !faceMatchResponse.Data.Match {
		StepError("STEP 8: Face compare mismatch: ", faceMatchResponse.Error)
		return verifyRecord.UpdateStatusReject(os.Getenv("AGENT_ID"), "Face compare mismatch: "+faceMatchResponse.Error)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Verified by Agent
	//||------------------------------------------------------------------------------------------------||

	StepInfo("VERIFICATION COMPLETED!", "")
	verifyRecord.IncrementStep()
	return verifyRecord.UpdateStatusVerified(os.Getenv("AGENT_ID"))
}
