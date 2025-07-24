package helpers

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
)

//||------------------------------------------------------------------------------------------------||
//|| GenerateCheckKEy
//||------------------------------------------------------------------------------------------------||

func GenerateCheckKey(privateKeyPEM string) (string, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return "", errors.New("invalid private key PEM format")
	}

	hash := sha256.Sum256([]byte(privateKeyPEM))
	checkKey := hex.EncodeToString(hash[:])

	return checkKey, nil
}

//||------------------------------------------------------------------------------------------------||
//|| CheckPrivateKey
//||------------------------------------------------------------------------------------------------||

func CheckPrivateKey(privateKeyPEM string, checkKey string) error {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return errors.New("invalid private key PEM format")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return err
	}

	if privateKey.N.String() != checkKey {
		return errors.New("private key does not match the provided key")
	}

	hash := sha256.Sum256([]byte(privateKeyPEM))
	if hex.EncodeToString(hash[:]) != checkKey {
		return errors.New("private key does not match the provided key")
	}

	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Encrypts data using RSA public key in PEM format
//||------------------------------------------------------------------------------------------------||

func EncryptWithPublicKey(data []byte, publicKeyPEM string) ([]byte, error) {
	block, _ := pem.Decode([]byte(publicKeyPEM))
	if block == nil || block.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM format")
	}

	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}

	encryptedData, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, data)
	if err != nil {
		return nil, err
	}

	return encryptedData, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Decrypts data using RSA private key in PEM format
//||------------------------------------------------------------------------------------------------||

func DecryptWithPrivateKey(ciphertext []byte, privateKeyPEM string) ([]byte, error) {
	block, _ := pem.Decode([]byte(privateKeyPEM))
	if block == nil || block.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key PEM format")
	}

	privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	plaintext, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, ciphertext)
	if err != nil {
		return nil, err
	}

	return plaintext, nil
}
