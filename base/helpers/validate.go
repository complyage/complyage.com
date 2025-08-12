package helpers

import (
	"base/interfaces"
	"crypto/rsa"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"errors"
	"fmt"
	"net/mail"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Validate Email
//||------------------------------------------------------------------------------------------------||

func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

//||------------------------------------------------------------------------------------------------||
//|| ValidateBIP39
//||------------------------------------------------------------------------------------------------||

func ValidateBIP39(wordListJSON string) (interfaces.BIPList, error) {
	var bipList interfaces.BIPList

	// Check if empty
	if strings.TrimSpace(wordListJSON) == "" {
		return bipList, errors.New("missing BIP39 word list")
	}

	// Unmarshal JSON into struct
	if err := json.Unmarshal([]byte(wordListJSON), &bipList); err != nil {
		return bipList, fmt.Errorf("invalid BIP39 word list format: %v", err)
	}

	// Convert to slice for validation
	words := []string{
		strings.ToLower(strings.TrimSpace(bipList.Word1)),
		strings.ToLower(strings.TrimSpace(bipList.Word2)),
		strings.ToLower(strings.TrimSpace(bipList.Word3)),
		strings.ToLower(strings.TrimSpace(bipList.Word4)),
		strings.ToLower(strings.TrimSpace(bipList.Word5)),
		strings.ToLower(strings.TrimSpace(bipList.Word6)),
	}

	// Build a lookup map from the official BIP39 list
	validMap := make(map[string]struct{})
	for _, w := range bip39() {
		validMap[w] = struct{}{}
	}

	// Validate all words
	for _, w := range words {
		if _, ok := validMap[w]; !ok {
			return bipList, fmt.Errorf("invalid BIP39 word: %s", w)
		}
	}

	// If valid, return the parsed BIPList
	return bipList, nil
}

//||------------------------------------------------------------------------------------------------||
//|| ValidateKeyPair
//||------------------------------------------------------------------------------------------------||

func ValidateKeyPair(privateKeyPEM, publicKeyPEM string) error {
	// Decode and parse the private key
	privBlock, _ := pem.Decode([]byte(privateKeyPEM))
	if privBlock == nil || privBlock.Type != "RSA PRIVATE KEY" {
		return errors.New("invalid private key PEM format")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(privBlock.Bytes)
	if err != nil {
		return errors.New("failed to parse RSA private key: " + err.Error())
	}

	// Decode and parse the public key
	pubBlock, _ := pem.Decode([]byte(publicKeyPEM))
	if pubBlock == nil || pubBlock.Type != "RSA PUBLIC KEY" {
		return errors.New("invalid public key PEM format")
	}
	pubKeyIface, err := x509.ParsePKIXPublicKey(pubBlock.Bytes)
	if err != nil {
		return errors.New("failed to parse RSA public key: " + err.Error())
	}
	pubKey, ok := pubKeyIface.(*rsa.PublicKey)
	if !ok {
		return errors.New("provided public key is not an RSA key")
	}

	// Ensure public key matches the private key
	if pubKey.N.Cmp(privateKey.PublicKey.N) != 0 || pubKey.E != privateKey.PublicKey.E {
		return errors.New("public key does not match the private key")
	}

	return nil
}
