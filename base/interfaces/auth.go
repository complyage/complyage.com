package interfaces

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Verification Record for TwoFactor
//||------------------------------------------------------------------------------------------------||

type BIPList struct {
	Word1 string
	Word2 string
	Word3 string
	Word4 string
	Word5 string
	Word6 string
}

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
	ID          int64
	Email       string
	Username    string
	Status      string
	Type        string
	Level       int8
	Security    int
	Private     string
	PrivateHash string
	Public      string
	Created     int64
	Expires     int64
	Identity    Identity
}
