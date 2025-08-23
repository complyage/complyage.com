package verify

import (
	"fmt"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Helper
//||------------------------------------------------------------------------------------------------||

func maskString(s string) string {
	if len(s) <= 2 {
		return s
	}
	return s[:2] + strings.Repeat("*", len(s)-2)
}

//||------------------------------------------------------------------------------------------------||
//|| Name
//||------------------------------------------------------------------------------------------------||

type Name struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle,omitempty"`
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s %s", n.First, n.Middle, n.Last)
}

func (n *Name) Mask() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", maskString(n.First), maskString(n.Middle), maskString(n.Last))
}

//||------------------------------------------------------------------------------------------------||
//|| DOB
//||------------------------------------------------------------------------------------------------||

type DOB struct {
	Month int `json:"month"`
	Day   int `json:"day"`
	Year  int `json:"year,omitempty"`
}

func (d *DOB) String() string {
	return fmt.Sprintf("%02d/%02d/%04d", d.Month, d.Day, d.Year)
}

func (d *DOB) Mask() string {
	if d == nil {
		return ""
	}
	// mask year fully, keep month/day
	return fmt.Sprintf("%02d/%02d/****", d.Month, d.Day)
}

//||------------------------------------------------------------------------------------------------||
//|| EmailAddress
//||------------------------------------------------------------------------------------------------||

type EmailAddress struct {
	Email string `json:"email"`
}

func (e *EmailAddress) String() string {
	return e.Email
}

func (e *EmailAddress) Mask() string {
	if e == nil || e.Email == "" {
		return ""
	}
	parts := strings.SplitN(e.Email, "@", 2)
	if len(parts) != 2 {
		return maskString(e.Email)
	}
	return maskString(parts[0]) + "@****"
}

//||------------------------------------------------------------------------------------------------||
//|| PhoneNumber
//||------------------------------------------------------------------------------------------------||

type PhoneNumber struct {
	CountryCode string `json:"countryCode"`
	Number      string `json:"number"`
}

func (p *PhoneNumber) String() string {
	return fmt.Sprintf("+%s %s", p.CountryCode, p.Number)
}

func (p *PhoneNumber) Mask() string {
	if p == nil {
		return ""
	}
	return fmt.Sprintf("+%s %s", p.CountryCode, maskString(p.Number))
}

//||------------------------------------------------------------------------------------------------||
//|| Address
//||------------------------------------------------------------------------------------------------||

type Address struct {
	Line1   string `json:"line1,omitempty"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Postal  string `json:"postal,omitempty"`
	Country string `json:"country,omitempty"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s %s, %s, %s %s, %s",
		a.Line1, a.Line2, a.City, a.State, a.Postal, a.Country)
}

func (a *Address) Mask() string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%s, %s, %s",
		maskString(a.City), maskString(a.State), maskString(a.Country))
}

//||------------------------------------------------------------------------------------------------||
//|| Media
//||------------------------------------------------------------------------------------------------||

type Media struct {
	Hash   string `json:"hash"`
	Size   int64  `json:"size,omitempty"`
	Base64 string `json:"blob,omitempty"`
	Mime   string `json:"mime,omitempty"`
}

func (m *Media) String() string {
	return fmt.Sprintf("Media(hash=%s, mime=%s, size=%d)", m.Hash, m.Mime, m.Size)
}

func (m *Media) Mask() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("Media(hash=%s, mime=%s, size=%d)", maskString(m.Hash), m.Mime, m.Size)
}

//||------------------------------------------------------------------------------------------------||
//|| Username
//||------------------------------------------------------------------------------------------------||

type Username struct {
	Username  string `json:"username"`
	FidSite   string `json:"fidSite"`
	Reference Media  `json:"reference"`
	Signed    Media  `json:"signed"`
}

func (u *Username) String() string {
	return fmt.Sprintf("%s@%s", u.Username, u.FidSite)
}

func (u *Username) Mask() string {
	if u == nil {
		return ""
	}
	return fmt.Sprintf("%s@%s", maskString(u.Username), maskString(u.FidSite))
}

//||------------------------------------------------------------------------------------------------||
//|| Identification
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

func (i *Identification) String() string {
	return fmt.Sprintf("%s %s (%s)", i.IDType, i.Number, i.Name.String())
}

func (i *Identification) Mask() string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%s %s (%s)", i.IDType, maskString(i.Number), i.Name.Mask())
}

//||------------------------------------------------------------------------------------------------||
//|| CreditCard
//||------------------------------------------------------------------------------------------------||

type CreditCard struct {
	LastFour      string  `json:"lastFour,omitempty"`
	CardType      string  `json:"cardType,omitempty"`
	ExpYear       string  `json:"expYear,omitempty"`
	Address       Address `json:"address"`
	TransactionId string  `json:"transactionId,omitempty"`
}

func (c *CreditCard) String() string {
	return fmt.Sprintf("%s ending in %s exp %s", c.CardType, c.LastFour, c.ExpYear)
}

func (c CreditCard) Mask() string {
	return fmt.Sprintf("%s ending in %s exp ****", c.CardType, c.LastFour)
}
