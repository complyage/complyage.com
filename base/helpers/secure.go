package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
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
