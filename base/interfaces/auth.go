package interfaces

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification Record for TwoFactor
//||------------------------------------------------------------------------------------------------||

type VerificationRecord struct {
	Code     string    // The secret 2FA code
	Key      []byte    // Encryption key for token
	Type     string    // Account type (USER, VNDR)
	Email    string    // Email address for verification
	Attempts int       // Track brute force attempts
	Created  time.Time // Email address for verification
	Expires  time.Time // Expiration time for the verification
}

//||------------------------------------------------------------------------------------------------||
//|| Session Record
//||------------------------------------------------------------------------------------------------||

type SessionRecord struct {
	ID            string
	Email         string
	Username      string
	Status        string
	Type          string
	Level         int8
	Advanced      bool
	Private       string
	PrivateCheck  string
	Public        string
	Created       int64
	Expires       int64
	Verifications []UserVerification
}
