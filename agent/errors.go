package main

import "github.com/ralphferrara/aria/app"

//||------------------------------------------------------------------------------------------------||
//|| Define this project's errors
//||------------------------------------------------------------------------------------------------||

func initErr() {
	app.Err("agent").Add("VERIFY_PANIC", "An unknown error occured", true)
	app.Err("agent").Add("2FA_TOO_MANY", "Too many attempts at Two factor verification", true)
	//||------------------------------------------------------------------------------------------------||
	//|| Face
	//||------------------------------------------------------------------------------------------------||
	app.Err("agent").Add("MODEL_FACE_DETECT_MARSHAL_FAIL", "Failed to marshal face detect request", true)
	app.Err("agent").Add("MODEL_FACE_DETECT_MODEL_FAIL", "Failed to load Face Model", true)
	app.Err("agent").Add("MODEL_FACE_DETECT_RESPONSE_FAIL", "Failed to parse model response", true)
	app.Err("agent").Add("MODEL_FACE_DETECT_NO_FACES", "Failure to detect face", true)
	app.Err("agent").Add("RJT_FACE_CONFIDENCE_LOW", "Failure to detect face", true)
	//||------------------------------------------------------------------------------------------------||
	//|| OCR
	//||------------------------------------------------------------------------------------------------||
	app.Err("agent").Add("MODEL_OCR_MODEL_FAIL", "Failed to load OCR Model", true)
	app.Err("agent").Add("MODEL_OCR_EMPTY", "OCR did not retrieve any text", true)
	//||------------------------------------------------------------------------------------------------||
	//|| DOB
	//||------------------------------------------------------------------------------------------------||
	app.Err("agent").Add("MODEL_DOB_MARSHAL_REQUEST", "DOB could not marshal request", false)
	app.Err("agent").Add("MODEL_DOB_MARSHAL_RESPONSE", "DOB could not marshal request", false)
	app.Err("agent").Add("MODEL_DOB_MODEL_FAIL", "DOB could not contact model", false)
	app.Err("agent").Add("MODEL_DOB_JSON_FAIL", "DOB did not retrieve any date of birth", false)
	app.Err("agent").Add("MODEL_DOB_INVALID_FORMAT", "DOB was not in a valid format", false)
	app.Err("agent").Add("MODEL_DOB_MISSING", "DOB was not retrieved", false)
	app.Err("agent").Add("MODEL_DOB_MISSING", "DOB was not retrieved", false)
	app.Err("agent").Add("MODEL_DOB_UNDERAGE", "DOB was underage", true)

}
