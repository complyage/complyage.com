package helpers

import (
	"crypto/rand"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification Record for TwoFactor
//||------------------------------------------------------------------------------------------------||

func Generate6DigitCode() (string, error) {
	n := make([]byte, 3)
	_, err := rand.Read(n)
	if err != nil {
		return "", err
	}
	code := (uint(n[0])<<16 | uint(n[1])<<8 | uint(n[2])) % 1000000
	return fmt.Sprintf("%06d", code), nil
}
