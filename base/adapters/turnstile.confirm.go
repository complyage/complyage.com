package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type TurnstileVerifyResponse struct {
	Success bool     `json:"success"`
	Error   []string `json:"error-codes"`
}

func VerifyTurnstile(token, ip string) error {
	secret := os.Getenv("TURNSTILE_SECRET_KEY")

	payload := []byte(fmt.Sprintf("secret=%s&response=%s&remoteip=%s", secret, token, ip))
	req, _ := http.NewRequest("POST", "https://challenges.cloudflare.com/turnstile/v0/siteverify", bytes.NewBuffer(payload))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return fmt.Errorf("turnstile http: %w", err)
	}
	defer resp.Body.Close()

	var result TurnstileVerifyResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return fmt.Errorf("turnstile decode: %w", err)
	}

	if !result.Success {
		return fmt.Errorf("turnstile failed: %v", result.Error)
	}
	return nil
}
