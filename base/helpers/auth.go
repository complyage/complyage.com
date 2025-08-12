package helpers

import (
	"base/db"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"os"

	"golang.org/x/crypto/argon2"
)

//||------------------------------------------------------------------------------------------------||
//|| Generate a Random Key
//||------------------------------------------------------------------------------------------------||

func GenerateRandom() ([]byte, error) {
	key := make([]byte, 32)
	_, err := rand.Read(key)
	return key, err
}

//||------------------------------------------------------------------------------------------------||
//|| Generate Username
//||------------------------------------------------------------------------------------------------||

func GenerateUsername() (string, error) {
	for i := 0; i < 5; i++ {
		// generate 6 random bytes -> 12 hex characters
		randomBytes := make([]byte, 6)
		if _, err := rand.Read(randomBytes); err != nil {
			return "", fmt.Errorf("failed to generate random bytes: %v", err)
		}

		username := "user_" + hex.EncodeToString(randomBytes)

		// check existence in DB
		var count int64
		err := db.DB.Table("accounts").
			Where("account_username = ?", username).
			Count(&count).Error
		if err != nil {
			return "", fmt.Errorf("db query failed: %v", err)
		}

		if count == 0 {
			return username, nil // ✅ found unique username
		}
	}

	return "", fmt.Errorf("failed to generate a unique username after 5 attempts")
}

//||------------------------------------------------------------------------------------------------||
//|| Generate an Email Hash
//||------------------------------------------------------------------------------------------------||

func GenerateEmailHash(email string) string {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Pepper
	//||------------------------------------------------------------------------------------------------||

	accountPepper := os.Getenv("ACCOUNT_PEPPER")
	if accountPepper == "" {
		fmt.Println("[GenerateEmailHash] ACCOUNT_PEPPER is not set") // Corrected log message
		return ""
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Hash
	//||------------------------------------------------------------------------------------------------||

	emailHash := sha256.Sum256([]byte(accountPepper + email))

	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||

	return base64.StdEncoding.EncodeToString(emailHash[:])
}

//||------------------------------------------------------------------------------------------------||
//|| Verify the Password
//||------------------------------------------------------------------------------------------------||

func VerifyPassword(password string, accountSalt string, accountPassword string) bool {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Pepper
	//||------------------------------------------------------------------------------------------------||

	pepper := os.Getenv("ACCOUNT_PEPPER")
	if pepper == "" {
		fmt.Println("[VerifyPassword] ACCOUNT_PEPPER is not set") // Corrected log message
		return false
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Decode the stored salt
	//||------------------------------------------------------------------------------------------------||

	decodedSalt, err := base64.StdEncoding.DecodeString(accountSalt)
	if err != nil {
		fmt.Println("[VerifyPassword] Failed to base64 decode accountSalt:", err)
		return false // Cannot verify if salt cannot be decoded
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make the Hash using the DECODED salt
	//||------------------------------------------------------------------------------------------------||

	hashedPassword := argon2.IDKey([]byte(pepper+password+pepper), decodedSalt, 1, 64*1024, 4, 32)

	fmt.Println("[VerifyPassword] Computed Hashed Password (Base64):", base64.StdEncoding.EncodeToString(hashedPassword))
	fmt.Println("[VerifyPassword] Stored Account Password (Base64):", accountPassword)

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	return (base64.StdEncoding.EncodeToString(hashedPassword) == accountPassword)
}

//||------------------------------------------------------------------------------------------------||
//|| Generate Password (for account creation)
//||------------------------------------------------------------------------------------------------||

func GeneratePassword(password string, passwordSalt string) (string, string) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get the Pepper
	//||------------------------------------------------------------------------------------------------||

	pepper := os.Getenv("ACCOUNT_PEPPER")
	if pepper == "" {
		fmt.Println("[GeneratePassword] ACCOUNT_PEPPER is not set")
		return "", ""
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Salt
	//||------------------------------------------------------------------------------------------------||

	salt := passwordSalt
	if salt == "" {
		saltBytes, pwErr := GenerateRandom()
		if pwErr != nil {
			fmt.Println("[GeneratePassword] Failed to generate salt:", pwErr)
			return "", ""
		}
		salt = string(saltBytes)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Make the Hash
	//||------------------------------------------------------------------------------------------------||

	hash := argon2.IDKey([]byte(pepper+password+pepper), []byte(salt), 1, 64*1024, 4, 32)

	//||------------------------------------------------------------------------------------------------||
	//|| Response - Encode both salt and hash to Base64 strings for storage
	//||------------------------------------------------------------------------------------------------||

	return base64.StdEncoding.EncodeToString(hash), base64.StdEncoding.EncodeToString([]byte(salt))
}
