package helpers

import (
	"base/interfaces"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Encrypt Helper
//||------------------------------------------------------------------------------------------------||

func encryptAndEncodeJSON(v interface{}, publicKey string) (string, error) {
	jsonBytes, err := json.Marshal(v)
	if err != nil {
		return "", err
	}
	encrypted, err := EncryptWithPublicKey(jsonBytes, publicKey)
	if err != nil {
		return "", err
	}
	return encodeBase64(encrypted), nil
}

//||------------------------------------------------------------------------------------------------||
//|| encodeBase64 is a small wrapper for base64 encoding
//||------------------------------------------------------------------------------------------------||

func encodeBase64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

//||------------------------------------------------------------------------------------------------||
//|| Create Email
//||------------------------------------------------------------------------------------------------||

func CreateEmail(email, publicKey string) (interfaces.VerifiedEmail, error) {
	if email == "" {
		return interfaces.VerifiedEmail{}, fmt.Errorf("email cannot be empty")
	}

	encrypted, err := encryptAndEncodeJSON(email, publicKey)
	if err != nil {
		return interfaces.VerifiedEmail{}, err
	}

	return interfaces.VerifiedEmail{
		Display:   MaskEmail(email),
		Data:      encrypted,
		Timestamp: time.Now(),
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create Phone
//||------------------------------------------------------------------------------------------------||

func CreatePhone(countryCode, number, publicKey string) (interfaces.VerifiedPhone, error) {
	if number == "" {
		return interfaces.VerifiedPhone{}, fmt.Errorf("phone cannot be empty")
	}

	decrypted := struct {
		CountryCode string `json:"countryCode"`
		Number      string `json:"number"`
	}{CountryCode: countryCode, Number: number}

	encrypted, err := encryptAndEncodeJSON(decrypted, publicKey)
	if err != nil {
		return interfaces.VerifiedPhone{}, err
	}

	return interfaces.VerifiedPhone{
		Display:   MaskPhone(countryCode, number),
		Data:      encrypted,
		Timestamp: time.Now(),
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create Address
//||------------------------------------------------------------------------------------------------||

func CreateAddress(street1, street2, city, state, zip, country, publicKey string) (interfaces.VerifiedAddress, error) {
	if street1 == "" && city == "" {
		return interfaces.VerifiedAddress{}, fmt.Errorf("address is invalid")
	}

	decrypted := struct {
		Street1 string `json:"street1"`
		Street2 string `json:"street2"`
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
		Country string `json:"country"`
	}{street1, street2, city, state, zip, country}

	encrypted, err := encryptAndEncodeJSON(decrypted, publicKey)
	if err != nil {
		return interfaces.VerifiedAddress{}, err
	}

	return interfaces.VerifiedAddress{
		Display:   MaskAddress(street1, city, country),
		Data:      encrypted,
		Timestamp: time.Now(),
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create Age
//||------------------------------------------------------------------------------------------------||

func CreateAge(year, month, day int, publicKey string) (interfaces.VerifiedAge, error) {
	if year == 0 {
		return interfaces.VerifiedAge{}, fmt.Errorf("birthdate is invalid")
	}

	decrypted := struct {
		Year  int `json:"year"`
		Month int `json:"month"`
		Day   int `json:"day"`
	}{year, month, day}

	encrypted, err := encryptAndEncodeJSON(decrypted, publicKey)
	if err != nil {
		return interfaces.VerifiedAge{}, err
	}

	display := fmt.Sprintf("%02d/%02d/%04d", month, day, year)

	return interfaces.VerifiedAge{
		Display:   display,
		Data:      encrypted,
		Media:     []interfaces.VerifiedMedia{},
		Timestamp: time.Now(),
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create Credit Card
//||------------------------------------------------------------------------------------------------||

func CreateCreditCard(last4, cardType string, expiresMonth, expiresYear int, transactionID, publicKey string) (interfaces.VerifiedCreditCard, error) {
	if last4 == "" {
		return interfaces.VerifiedCreditCard{}, fmt.Errorf("credit card is invalid")
	}

	decrypted := struct {
		Last4   string `json:"last4"`
		Type    string `json:"type"`
		Expires struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"expires"`
		TransactionID string `json:"transactionId"`
	}{Last4: last4, Type: cardType, TransactionID: transactionID}
	decrypted.Expires.Month = expiresMonth
	decrypted.Expires.Year = expiresYear

	encrypted, err := encryptAndEncodeJSON(decrypted, publicKey)
	if err != nil {
		return interfaces.VerifiedCreditCard{}, err
	}

	return interfaces.VerifiedCreditCard{
		Display:   MaskCreditCard(last4),
		Data:      encrypted,
		Timestamp: time.Now(),
	}, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create Username
//||------------------------------------------------------------------------------------------------||

func CreateUsername(username string, siteID int64, publicKey string) (interfaces.VerifiedUsername, error) {
	if username == "" {
		return interfaces.VerifiedUsername{}, fmt.Errorf("username is empty")
	}

	decrypted := struct {
		SiteID   int64  `json:"siteId"`
		Username string `json:"username"`
	}{SiteID: siteID, Username: username}

	encrypted, err := encryptAndEncodeJSON(decrypted, publicKey)
	if err != nil {
		return interfaces.VerifiedUsername{}, err
	}

	return interfaces.VerifiedUsername{
		Display:   MaskUsername(username, siteID),
		Data:      encrypted,
		Timestamp: time.Now(),
	}, nil
}
