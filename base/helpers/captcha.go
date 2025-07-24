package helpers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| Veridy Turnstile Token
//||------------------------------------------------------------------------------------------------||

func VerifyTurnstileToken(token string, remoteIP string) (bool, error) {
	secret := os.Getenv("TURNSTILE_SECRET")
	if secret == "" {
		return false, errors.New("missing TURNSTILE_SECRET")
	}

	body := fmt.Sprintf("secret=%s&response=%s&remoteip=%s", secret, token, remoteIP)
	resp, err := http.Post(
		"https://challenges.cloudflare.com/turnstile/v0/siteverify",
		"application/x-www-form-urlencoded",
		bytes.NewBufferString(body),
	)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result struct {
		Success bool `json:"success"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return false, err
	}

	return result.Success, nil
}
