package interfaces

import (
	"base/constants"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Verified Media
//||------------------------------------------------------------------------------------------------||

type VerifiedMedia struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Hash string `json:"hash"`
		Size int64  `json:"size"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Email
//||------------------------------------------------------------------------------------------------||

type VerifiedEmail struct {
	Display   string    `json:"display"`
	Data      string    `json:"data"`
	Decrypted *string   `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Age
//||------------------------------------------------------------------------------------------------||

type VerifiedAge struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Month int `json:"month"`
		Day   int `json:"day"`
		Year  int `json:"year"`
	} `json:"decrypted,omitempty"`
	Media     []VerifiedMedia `json:"media"`
	Timestamp time.Time       `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Phone
//||------------------------------------------------------------------------------------------------||

type VerifiedPhone struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		CountryCode string `json:"countryCode"`
		Number      string `json:"number"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Address
//||------------------------------------------------------------------------------------------------||

type VerifiedAddress struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Street1 string `json:"street1"`
		Street2 string `json:"street2"`
		City    string `json:"city"`
		State   string `json:"state"`
		Zip     string `json:"zip"`
		Country string `json:"country"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Identification
//||------------------------------------------------------------------------------------------------||

type VerifiedIdentification struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Name struct {
			First  string `json:"first"`
			Last   string `json:"last"`
			Middle string `json:"middle"`
		} `json:"name"`
		DOB struct {
			Month int `json:"month"`
			Day   int `json:"day"`
			Year  int `json:"year"`
		} `json:"dob"`
		Address struct {
			Street1 string `json:"street1"`
			Street2 string `json:"street2"`
			City    string `json:"city"`
			State   string `json:"state"`
			Zip     string `json:"zip"`
			Country string `json:"country"`
		} `json:"address"`
		IDNumber string `json:"idNumber"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Credit Card
//||------------------------------------------------------------------------------------------------||

type VerifiedCreditCard struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Last4   string `json:"last4"`
		Type    string `json:"type"`
		Expires struct {
			Month int `json:"month"`
			Year  int `json:"year"`
		} `json:"expires"`
		TransactionID string `json:"transactionId"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Verified Username
//||------------------------------------------------------------------------------------------------||

type VerifiedUsername struct {
	Display   string `json:"display"`
	Data      string `json:"data"`
	Decrypted *struct {
		Reference VerifiedMedia   `json:"reference"`
		Media     []VerifiedMedia `json:"media"`
		SiteID    int64           `json:"siteId"`
		Username  string          `json:"username"`
	} `json:"decrypted,omitempty"`
	Timestamp time.Time `json:"timestamp"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity (Epic Verified Profile)
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	Email      *VerifiedEmail               `json:"email"`      // nil represents "MISS"
	Age        *VerifiedAge                 `json:"age"`        // nil represents "MISS"
	Phone      *VerifiedPhone               `json:"phone"`      // nil represents "MISS"
	Address    *VerifiedAddress             `json:"address"`    // nil represents "MISS"
	CreditCard *VerifiedCreditCard          `json:"creditCard"` // nil represents "MISS"
	Usernames  map[int64]VerifiedUsername   `json:"usernames"`
	Approved   []constants.VerificationType `json:"approved"` // List of approved verification types
}
