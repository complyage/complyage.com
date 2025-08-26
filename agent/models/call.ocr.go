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

// ModelOCRRequest is the request struct for OCR extraction.
type ModelOCRRequest struct {
	Media  string `json:"image"`
	Prompt string `json:"prompt"`
}

type ModelOCRResponse struct {
	Success bool                  `json:"success"`
	Error   string                `json:"error"`
	Data    *ModelOCRResponseData `json:"data"`
	Elapsed float64               `json:"elapsed"`
}

type ModelOCRResponseData struct {
	Text string `json:"text"`
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallOCR: Calls the OCR extraction model API endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallOCR(media publish.AgentMedia, prompt string) (*ModelOCRResponse, error) {

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

	req := ModelOCRRequest{
		Media:  resizedBase64,
		Prompt: prompt,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal OCR request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Calling OCR model...")
	respBytes, err := ModelCall("/ocr/extract", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelOCRResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
