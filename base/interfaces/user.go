package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Base
//||------------------------------------------------------------------------------------------------||

type Name struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle,omitempty"`
}

type DOB struct {
	Month int `json:"month"`
	Day   int `json:"day"`
	Year  int `json:"year,omitempty"`
}

type EmailAddress struct {
	Email string `json:"email"`
}

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

type Address struct {
	Line1   string `json:"line1,omitempty"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Postal  string `json:"postal,omitempty"`
	Country string `json:"country,omitempty"`
}

type Media struct {
	Hash     string `json:"hash"`
	Size     int64  `json:"size,omitempty"`
	Blob     string `json:"blob,omitempty"`
	MimeType string `json:"mimeType,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Complex - Username
//||------------------------------------------------------------------------------------------------||

type Username struct {
	Username  string `json:"username"`
	FidSite   string `json:"fidSite"`
	Reference Media  `json:"reference"`
	Signed    Media  `json:"signed"`
}

//||------------------------------------------------------------------------------------------------||
//|| ID Card
//||------------------------------------------------------------------------------------------------||

type Identification struct {
	IDType  string  `json:"idType,omitempty"`
	Number  string  `json:"number,omitempty"`
	Front   Media   `json:"front,omitempty"`
	Back    Media   `json:"back,omitempty"`
	Selfie  Media   `json:"selfie,omitempty"`
	Address Address `json:"address,omitempty"`
	DOB     *DOB    `json:"dob,omitempty"`
	Name    *Name   `json:"name,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Complex - Credit Card
//||------------------------------------------------------------------------------------------------||

type CreditCard struct {
	Number        string  `json:"number,omitempty"`
	LastFour      string  `json:"lastFour,omitempty"`
	CardType      string  `json:"cardType,omitempty"`
	ExpMonth      string  `json:"expMonth,omitempty"`
	ExpYear       string  `json:"expYear,omitempty"`
	CVC           string  `json:"cvc,omitempty"`
	Address       Address `json:"address"`
	TransactionId string  `json:"transactionId,omitempty"`
}
