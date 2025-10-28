package tests

import (
	"testing"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| Test AccountSessionChecked Struct with Real KeyPair
//||------------------------------------------------------------------------------------------------||

func GenerateAccountSessionChecked(t *testing.T) types.AccountSessionChecked {
	// Generate a real key pair for testing
	priv, pub, check := GenerateKeyPair(t)

	// Create a test session using the generated keys
	session := types.AccountSessionChecked{
		ID:         1,
		Salt:       "test_salt",
		Username:   "testuser",
		Identifier: "testuser@example.com",
		Status:     "ACTIVE",
		Level:      5,
		KeysLoaded: true,
		Private:    priv,
		Public:     pub,
		CheckKey:   check,
	}

	return session
}
