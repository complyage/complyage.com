package interfaces

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification Record for TwoFactor
//||------------------------------------------------------------------------------------------------||

type AuthVerification struct {
	Code     string    // The secret 2FA code
	Key      []byte    // Encryption key for token
	Type     string    // Account type (USER, VNDR)
	Email    string    // Email address for verification
	Attempts int       // Track brute force attempts
	Created  time.Time // Email address for verification
	Expires  time.Time // Expiration time for the verification
}
