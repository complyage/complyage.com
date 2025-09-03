package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/images"
	"agent/publish"
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

// ModelVisionAnalyzeRequest is the request struct for vision analysis.
type ModelVisionAnalyzeRequest struct {
	Media  string `json:"image"`
	Prompt string `json:"prompt"`
}

type ModelVisionAnalyzeResponse struct {
	Success bool                    `json:"success"`
	Error   string                  `json:"error"`
	Data    *ModelVisionAnalyzeData `json:"data"`
	Elapsed float64                 `json:"elapsed"`
}

type ModelVisionAnalyzeData struct {
	Text string `json:"text"`
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallVisionAnalyze: Calls the vision analyze model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallVisionAnalyze(media publish.AgentMedia, prompt string) (*ModelVisionAnalyzeResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Resize Image
	//||------------------------------------------------------------------------------------------------||

	const maxDim = 1024 // Adjust as needed (e.g. 800, 640, etc.)
	resizedBase64, err := images.ResizeBase64Image(media.Base64, maxDim)
	if err != nil {
		return nil, fmt.Errorf("failed to resize image: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelVisionAnalyzeRequest{
		Media:  resizedBase64,
		Prompt: prompt,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal vision analyze request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/vision/analyze", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelVisionAnalyzeResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
