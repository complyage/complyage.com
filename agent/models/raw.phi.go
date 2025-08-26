package models

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| ModelCallPhiRawString: Sends a prompt (and model) and returns only the raw string response from Phi3
//||------------------------------------------------------------------------------------------------||

func ModelCallPhiRawString(prompt, model string) (string, error) {
	type req struct {
		Prompt string `json:"prompt,omitempty"`
		Model  string `json:"model,omitempty"`
	}
	type resp struct {
		Success bool   `json:"success"`
		Error   string `json:"error"`
		Data    struct {
			Raw string `json:"raw"`
		} `json:"data"`
	}

	// Build and marshal request - FIX: include both prompt and model
	payload, err := json.Marshal(req{Prompt: prompt, Model: model})
	if err != nil {
		return "", fmt.Errorf("failed to marshal request: %w", err)
	}

	// Call phi3 model endpoint
	respBytes, err := ModelCall("/ollama/generate", string(payload))
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
