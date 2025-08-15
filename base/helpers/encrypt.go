package helpers

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"io"
	"strings"
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

	//||------------------------------------------------------------------------------------------------||
	//|| Generate AES-256 Key
	//||------------------------------------------------------------------------------------------------||

	aesKey := make([]byte, 32)
	if _, err := rand.Read(aesKey); err != nil {
		return nil, err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt Data with AES-GCM
	//||------------------------------------------------------------------------------------------------||

	block, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, aesGCM.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	cipherData := aesGCM.Seal(nil, nonce, data, nil)

	//||------------------------------------------------------------------------------------------------||
	//|| Encrypt AES Key with RSA Public Key
	//||------------------------------------------------------------------------------------------------||

	pemBlock, _ := pem.Decode([]byte(publicKeyPEM))
	if pemBlock == nil || pemBlock.Type != "RSA PUBLIC KEY" {
		return nil, errors.New("invalid public key PEM format")
	}
	pub, err := x509.ParsePKIXPublicKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	publicKey, ok := pub.(*rsa.PublicKey)
	if !ok {
		return nil, errors.New("not an RSA public key")
	}
	encryptedAESKey, err := rsa.EncryptPKCS1v15(rand.Reader, publicKey, aesKey)
	if err != nil {
		return nil, err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Join as base64 strings: encryptedAESKey.nonce.cipherData
	//||------------------------------------------------------------------------------------------------||

	final := base64.StdEncoding.EncodeToString(encryptedAESKey) + "." +
		base64.StdEncoding.EncodeToString(nonce) + "." +
		base64.StdEncoding.EncodeToString(cipherData)

	return []byte(final), nil
}

//||------------------------------------------------------------------------------------------------||
//|| HybridDecrypt : Decrypts data encrypted with HybridEncrypt above using RSA private key (PEM)
//|| Expects base64(encryptedAESKey) + "." + base64(nonce) + "." + base64(encryptedData)
//||------------------------------------------------------------------------------------------------||

func DecryptWithPrivateKey(ciphertext []byte, privateKeyPEM string) ([]byte, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Split base64 parts
	//||------------------------------------------------------------------------------------------------||

	parts := strings.SplitN(string(ciphertext), ".", 3)
	if len(parts) != 3 {
		return nil, errors.New("invalid ciphertext format for hybrid decrypt")
	}

	encryptedAESKey, err := base64.StdEncoding.DecodeString(parts[0])
	if err != nil {
		return nil, err
	}
	nonce, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return nil, err
	}
	cipherData, err := base64.StdEncoding.DecodeString(parts[2])
	if err != nil {
		return nil, err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decrypt AES Key with RSA Private Key
	//||------------------------------------------------------------------------------------------------||

	pemBlock, _ := pem.Decode([]byte(privateKeyPEM))
	if pemBlock == nil || pemBlock.Type != "RSA PRIVATE KEY" {
		return nil, errors.New("invalid private key PEM format")
	}
	privateKey, err := x509.ParsePKCS1PrivateKey(pemBlock.Bytes)
	if err != nil {
		return nil, err
	}
	aesKey, err := rsa.DecryptPKCS1v15(rand.Reader, privateKey, encryptedAESKey)
	if err != nil {
		return nil, err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decrypt Data with AES-GCM
	//||------------------------------------------------------------------------------------------------||

	blockAES, err := aes.NewCipher(aesKey)
	if err != nil {
		return nil, err
	}
	aesGCM, err := cipher.NewGCM(blockAES)
	if err != nil {
		return nil, err
	}
	plainData, err := aesGCM.Open(nil, nonce, cipherData, nil)
	if err != nil {
		return nil, err
	}

	return plainData, nil
}
