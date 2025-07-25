package helpers

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"math/big"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Static BIP39 Wordlist (embed the actual list here)
//||------------------------------------------------------------------------------------------------||

const bip39Words = `abandon
ability
able
about
above
absent
absorb
abstract
absurd
abuse
access
accident
account
accuse
...` // ← Replace with full 2048 BIP39 word list (newline-separated)

//||------------------------------------------------------------------------------------------------||
//|| Return a list of `count` random BIP39 words from the embedded list
//||------------------------------------------------------------------------------------------------||

func GetBip39List(count int) ([]string, error) {
	words := strings.Split(strings.TrimSpace(bip39Words), ",")
	if len(words) != 2048 {
		return nil, fmt.Errorf("invalid BIP39 word list: got %d words", len(words))
	}

	result := make([]string, 0, count)
	for i := 0; i < count; i++ {
		indexBig, err := rand.Int(rand.Reader, big.NewInt(2048))
		if err != nil {
			return nil, fmt.Errorf("failed to generate secure random index: %v", err)
		}
		result = append(result, words[indexBig.Int64()])
	}

	return result, nil
}

//||------------------------------------------------------------------------------------------------||
//|| GenerateDeterministicRSAKeysFromWords
//|| Returns: *rsa.PrivateKey, privatePEM string, publicPEM string
//||------------------------------------------------------------------------------------------------||

func GeneratePrivatePublicBIP39(input string, bits int) (*rsa.PrivateKey, string, string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Normalize & clean word list
	//||------------------------------------------------------------------------------------------------||

	parts := strings.Split(input, ",")
	var clean []string
	for _, word := range parts {
		word = strings.ToLower(strings.TrimSpace(word))
		if word != "" {
			clean = append(clean, word)
		}
	}

	if len(clean) == 0 {
		return nil, "", "", errors.New("no valid BIP39 words provided")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create seed from SHA256 of mnemonic
	//||------------------------------------------------------------------------------------------------||

	mnemonic := strings.Join(clean, " ")
	seed := sha256.Sum256([]byte(mnemonic))

	//||------------------------------------------------------------------------------------------------||
	//|| Deterministic reader → seeded with hash
	//||------------------------------------------------------------------------------------------------||

	reader := NewDeterministicReader(seed[:])
	privKey, err := rsa.GenerateKey(reader, bits)
	if err != nil {
		return nil, "", "", err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Encode Private Key to PEM
	//||------------------------------------------------------------------------------------------------||

	privDER := x509.MarshalPKCS1PrivateKey(privKey)
	privPEM := new(bytes.Buffer)
	_ = pem.Encode(privPEM, &pem.Block{
		Type:  "RSA PRIVATE KEY",
		Bytes: privDER,
	})

	//||------------------------------------------------------------------------------------------------||
	//|| Encode Public Key to PEM
	//||------------------------------------------------------------------------------------------------||

	pubDER, err := x509.MarshalPKIXPublicKey(&privKey.PublicKey)
	if err != nil {
		return nil, "", "", err
	}

	pubPEM := new(bytes.Buffer)
	_ = pem.Encode(pubPEM, &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: pubDER,
	})

	return privKey, privPEM.String(), pubPEM.String(), nil
}

//||------------------------------------------------------------------------------------------------||
