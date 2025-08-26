package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/publish"
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

// ModelFaceDetectRequest is the request struct for face detection.
type ModelFaceDetectRequest struct {
	Media  string `json:"image"`
	Prompt string `json:"prompt"`
}

// ModelFaceDetectResponse matches the /face/detect endpoint response.
type ModelFaceDetectResponse struct {
	Success bool                  `json:"success"`
	Error   string                `json:"error"`
	Faces   []ModelFaceDetectFace `json:"faces"`
}

type ModelFaceDetectFace struct {
	BoundingBox ModelFaceBoundingBox `json:"bounding_box"`
	Age         int                  `json:"age"`
	AgeRange    string               `json:"age_range"`
	Gender      string               `json:"gender"`
	Confidence  *float64             `json:"confidence"` // nullable
}

type ModelFaceBoundingBox struct {
	X      float64 `json:"x"`
	Y      float64 `json:"y"`
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
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
		return nil, fmt.Errorf("failed to marshal face detect request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/face/detect", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelFaceDetectResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
