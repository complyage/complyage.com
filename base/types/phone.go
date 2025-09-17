package types

import (
	"fmt"
	"regexp"
	"strings"
)

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

func PhoneFromString(raw string) (PhoneNumber, error) {
	// Remove spaces, dashes, parentheses
	cleaned := strings.ReplaceAll(raw, " ", "")
	cleaned = strings.ReplaceAll(cleaned, "-", "")
	cleaned = strings.ReplaceAll(cleaned, "(", "")
	cleaned = strings.ReplaceAll(cleaned, ")", "")

	// Regex: +<countryCode><number>
	re := regexp.MustCompile(`^\+?(\d{1,4})(\d{6,15})$`)
	matches := re.FindStringSubmatch(cleaned)
	if matches == nil {
		return PhoneNumber{}, fmt.Errorf("invalid phone number: %s", raw)
	}

	return PhoneNumber{
		CountryCode: matches[1],
		Number:      matches[2],
	}, nil
}
