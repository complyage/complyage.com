package helpers

import (
	"base/interfaces"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"strings"
)

func GenerateKeyPair() (privateKeyPEM string, publicKeyPEM string, err error) {
	key, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		return "", "", err
	}

	// Encode private key
	privDER := x509.MarshalPKCS1PrivateKey(key)
	privBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}
	privateKeyPEM = string(pem.EncodeToMemory(&privBlock))

	// Encode public key
	pubDER, err := x509.MarshalPKIXPublicKey(&key.PublicKey)
	if err != nil {
		return "", "", err
	}
	pubBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDER,
	}
	publicKeyPEM = string(pem.EncodeToMemory(&pubBlock))

	return privateKeyPEM, publicKeyPEM, nil
}

func GenerateBIP39Keys(bipList interfaces.BIPList) (privateKeyPEM string, publicKeyPEM string, err error) {
	const keySize = 2048 // fixed RSA key size for consistency

	// Build mnemonic from BIPList
	words := []string{
		strings.ToLower(strings.TrimSpace(bipList.Word1)),
		strings.ToLower(strings.TrimSpace(bipList.Word2)),
		strings.ToLower(strings.TrimSpace(bipList.Word3)),
		strings.ToLower(strings.TrimSpace(bipList.Word4)),
		strings.ToLower(strings.TrimSpace(bipList.Word5)),
		strings.ToLower(strings.TrimSpace(bipList.Word6)),
	}

	// Filter out any empty words
	var clean []string
	for _, w := range words {
		if w != "" {
			clean = append(clean, w)
		}
	}
	if len(clean) == 0 {
		return "", "", errors.New("no valid BIP39 words provided")
	}

	// Create deterministic seed from mnemonic
	mnemonic := strings.Join(clean, " ")
	seed := sha256.Sum256([]byte(mnemonic))

	// Use deterministic reader seeded with hash
	reader := NewDeterministicReader(seed[:])
	privKey, err := rsa.GenerateKey(reader, keySize)
	if err != nil {
		return "", "", err
	}

	// Encode Private Key (PKCS#1)
	privDER := x509.MarshalPKCS1PrivateKey(privKey)
	privBlock := pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	}
	privateKeyPEM = string(pem.EncodeToMemory(&privBlock))

	// Encode Public Key
	pubDER, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return "", "", err
	}
	pubBlock := pem.Block{
		Type:  "RSA PUBLIC KEY",
		Bytes: pubDER,
	}
	publicKeyPEM = string(pem.EncodeToMemory(&pubBlock))

	return privateKeyPEM, publicKeyPEM, nil
}
