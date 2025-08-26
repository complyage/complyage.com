package models

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| ModelCallGemmaRawString: Sends a prompt and returns only the raw string response
//||------------------------------------------------------------------------------------------------||

func ModelCallGemmaRawString(prompt string) (string, error) {
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

	// Call model
	respBytes, err := ModelCall("/gemma/generate", string(payload))
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
