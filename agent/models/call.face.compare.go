package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

// ModelFaceCompareRequest is the request struct for face comparison.
type ModelFaceCompareRequest struct {
	Image1 string      `json:"image1"` // base64 string
	Image2 string      `json:"image2"` // base64 string
	Params interface{} `json:"params"` // can be a map or struct
}

// ModelFaceCompareResponse matches the /face/compare endpoint response.
type ModelFaceCompareResponse struct {
	Success bool                  `json:"success"`
	Error   string                `json:"error"`
	Data    *ModelFaceCompareData `json:"data"`
}

type ModelFaceCompareData struct {
	Match      bool       `json:"match"`
	Score      float64    `json:"score"`
	Threshold  float64    `json:"threshold"`
	Metric     string     `json:"metric"`
	Image1Bbox [4]float64 `json:"image1_bbox"`
	Image2Bbox [4]float64 `json:"image2_bbox"`
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallFaceCompare: Calls the face compare model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

// NOTE: image1, image2 are expected as base64-encoded strings, params can be any struct or map.
func ModelCallFaceCompare(image1 string, image2 string, params interface{}) (*ModelFaceCompareResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelFaceCompareRequest{
		Image1: image1,
		Image2: image2,
		Params: params,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal face compare request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/face/compare", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelFaceCompareResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
