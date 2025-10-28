package tests

import (
	"testing"

	"github.com/ralphferrara/aria/base/encrypt"
)

//||------------------------------------------------------------------------------------------------||
//|| Generate the Key Pair
//||------------------------------------------------------------------------------------------------||

func GenerateKeyPair(t *testing.T) (privateKey, publicKey, checkKey string) {
	privateKey, publicKey, err := encrypt.GenerateKeyPair()
	if err != nil {
		t.Fatalf("failed to generate Key Pair: %v", err)
	}
	checkKey, ckErr := encrypt.GenerateCheckKey(privateKey)
	if ckErr != nil {
		t.Fatalf("failed to generate Check Key: %v", err)
	}
	return privateKey, publicKey, checkKey
}
