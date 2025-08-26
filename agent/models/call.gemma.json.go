package models

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/template"
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

// ModelGemmaJSONRequest is the request struct for the Gemma JSON extraction API.
type ModelGemmaJSONRequest struct {
	Prompt string `json:"prompt,omitempty"`
}

// ModelGemmaJSONResponse matches the /gemma/generate endpoint response.
type ModelGemmaJSONResponse struct {
	Success bool                  `json:"success"`
	Error   string                `json:"error"`
	Data    *ModelGemmaJSONResult `json:"data"`
	Elapsed float64               `json:"elapsed,omitempty"`
}

type ModelGemmaJSONResult struct {
	Raw string `json:"raw"`
}

//||------------------------------------------------------------------------------------------------||
//|| Prompt
//||------------------------------------------------------------------------------------------------||

// Don't use "json" as a variable or parameter name!
func prompt(text string, jsonStruct any) string {
	tpl := template.Create("promptIDGemma")
	tpl.Add("OCRTEXT", text)

	promptStr, err := tpl.Compile()
	if err != nil {
		fmt.Println("Failed to build prompt:", err)
		return "Failed to build prompt"
	}
	return promptStr
}

//||------------------------------------------------------------------------------------------------||
//|| ModelCallGemmaJSON: Calls the Gemma JSON extraction endpoint and parses the response
//||------------------------------------------------------------------------------------------------||

func ModelCallGemmaJSON(text string, jsonStruct any) (*ModelGemmaJSONResponse, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Prompt
	//||------------------------------------------------------------------------------------------------||

	promptStr := prompt(text, jsonStruct)

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	req := ModelGemmaJSONRequest{
		Prompt: promptStr,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	payload, err := json.Marshal(req)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal Gemma JSON request: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Call Model
	//||------------------------------------------------------------------------------------------------||

	respBytes, err := ModelCall("/gemma/generate", string(payload))
	if err != nil {
		return nil, fmt.Errorf("model call failed: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Response
	//||------------------------------------------------------------------------------------------------||

	var resp ModelGemmaJSONResponse
	if err := json.Unmarshal(respBytes, &resp); err != nil {
		return nil, fmt.Errorf("failed to parse model response: %w", err)
	}

	return &resp, nil
}
