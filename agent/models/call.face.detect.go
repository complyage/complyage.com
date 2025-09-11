package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/publish"
	"encoding/json"
	"math"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Request/Response Structs
//||------------------------------------------------------------------------------------------------||

// ModelFaceDetectRequest is the request struct for face detection.
type ModelFaceDetectRequest struct {
	Media  string `json:"image"`
	Prompt string `json:"prompt"`
}

// ModelFaceDetectResponse matches the /face/detect endpoint response.
type ModelFaceResponse struct {
	Success bool                  `json:"success"`
	Error   string                `json:"error"`
	Faces   []ModelFaceDetectFace `json:"faces"`
}

type ModelFaceDetectFace struct {
	BoundingBox ModelFaceBoundingBox `json:"bounding_box"`
	Age         int                  `json:"age"`
	Gender      string               `json:"gender"`
	Confidence  *float64             `json:"confidence"`
}

type ModelFaceBoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

type ModelFaceDetectResponse struct {
	Success     bool                 `json:"success"`
	BoundingBox ModelFaceBoundingBox `json:"bounding_box,omitempty"`
	Age         int                  `json:"age,omitempty"`
	AgeMin      int                  `json:"age_min,omitempty"`
	AgeMax      int                  `json:"age_max,omitempty"`
	Gender      string               `json:"gender,omitempty"`
	Confidence  *float64             `json:"confidence,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Min/Max Age
//||------------------------------------------------------------------------------------------------||

func AgeRange(age int) (min, max int) {
	if age < 0 {
		return 0, 0
	}
	if age < 18 {
		min, max = age-1, age+1
	} else {
		// Growth ([larger] -1.5-2 [smaller])
		const growth = 1.5
		window := int(1 + math.Sqrt(float64(age-18))/growth)
		min, max = age-window, age+window
	}
	if min < 0 {
		min = 0
	}
	return min, max
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallFaceDetect: Calls the face detect model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallFaceDetect(media publish.AgentMedia, prompt string) (*ModelFaceDetectResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelFaceDetectRequest{
		Media:  media.Base64,
		Prompt: prompt,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, app.Err("agent").Error("MODEL_FACE_DETECT_MARSHAL_FAIL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/face/detect", string(payload))
	if err != nil {
		return nil, app.Err("agent").Error("MODEL_FACE_DETECT_MODEL_FAIL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var modelResp ModelFaceResponse
	if err := json.Unmarshal(respBytes, &modelResp); err != nil {
		return nil, app.Err("agent").Error("MODEL_FACE_DETECT_RESPONSE_FAIL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	if (!modelResp.Success || len(modelResp.Faces) == 0) && modelResp.Error == "" {
		return nil, app.Err("agent").Error("MODEL_FACE_DETECT_NO_FACE")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Min Max
	//||------------------------------------------------------------------------------------------------||

	min, max := AgeRange(modelResp.Faces[0].Age)

	//||------------------------------------------------------------------------------------------------||
	//|| Return Response
	//||------------------------------------------------------------------------------------------------||

	resp := ModelFaceDetectResponse{
		BoundingBox: modelResp.Faces[0].BoundingBox,
		Age:         modelResp.Faces[0].Age,
		AgeMin:      min,
		AgeMax:      max,
		Gender:      modelResp.Faces[0].Gender,
		Confidence:  modelResp.Faces[0].Confidence,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	return &resp, nil
}
