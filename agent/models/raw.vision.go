package models

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| ModelCallVisionRawString: Sends a prompt and returns only the raw string response from Vision
//||------------------------------------------------------------------------------------------------||

func ModelCallVisionRawString(prompt string) (string, error) {
	type req struct {
		Prompt string `json:"prompt,omitempty"`
	}
	type resp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Data    struct {
			Raw string `json:"raw"`
		} `json:"data"`
	}

	// Build and marshal request
	payload, err := json.Marshal(req{Prompt: prompt})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Call vision model endpoint
	respBytes, err := ModelCall("/vision/content", string(payload))
	if err != nil {
		return "", fmt.Errorf("model call failed: %w", err)
	}

	// Parse response
	var r resp
	if err := json.Unmarshal(respBytes, &r); err != nil {
		return "", fmt.Errorf("failed to parse response: %w", err)
	}
	if !r.Success {
		return "", fmt.Errorf("model error: %s", r.Error)
	}

	return r.Data.Raw, nil
}
