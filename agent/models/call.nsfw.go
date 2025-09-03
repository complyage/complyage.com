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

// ModelNSFWRequest is the request struct for NSFW model API.
type ModelNSFWRequest struct {
	Media  publish.AgentMedia `json:"media"`
	Prompt string             `json:"prompt"`
}

type ModelNSFWResponse struct {
	Success bool                   `json:"success"`
	Error   string                 `json:"error"`
	Data    *ModelNSFWResponseData `json:"data"`
}

type ModelNSFWResponseData struct {
	NSFW       bool                       `json:"nsfw"`
	Reason     string                     `json:"reason"`
	Score      float64                    `json:"score"`
	Detections []ModelNSFWDetectionResult `json:"detections"`
}

type ModelNSFWDetectionResult struct {
	Class string  `json:"class"`
	Score float64 `json:"score"`
	Box   [4]int  `json:"box"`
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallNSFW: Calls the NSFW OCR model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallNSFW(media publish.AgentMedia, prompt string) (*ModelNSFWResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelNSFWRequest{
		Media:  media,
		Prompt: prompt,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal NSFW OCR request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/v1/nsfw", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelNSFWResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
