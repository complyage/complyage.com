package helpers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/interfaces"
	"crypto/rand"
	"errors"
	"math/big"
	"regexp"
	"strconv"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Public API
//||------------------------------------------------------------------------------------------------||

func ValidateVerificationCard(vc interfaces.VerificationCardRequest) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Required pointer fields check
	//||------------------------------------------------------------------------------------------------||
	if vc.CardNumber == nil || vc.ExpMonth == nil || vc.ExpYear == nil || vc.CVC == nil || vc.BillingZip == nil {
		return errors.New("Missing required card details")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Card number (Luhn + length)
	//||------------------------------------------------------------------------------------------------||
	digits := OnlyDigits(*vc.CardNumber)
	if len(digits) < 12 || !luhn(digits) {
		return errors.New("Invalid card number")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Expiry month
	//||------------------------------------------------------------------------------------------------||
	mm := ToInt(*vc.ExpMonth)
	if mm < 1 || mm > 12 {
		return errors.New("Invalid expiry month")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Expiry year
	//||------------------------------------------------------------------------------------------------||
	yy := *vc.ExpYear
	if len(yy) == 2 {
		yy = "20" + yy
	}
	yyyy := ToInt(yy)
	if yyyy < 2000 || yyyy > 2100 {
		return errors.New("Invalid expiry year")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Expiration check (valid through end of month)
	//||------------------------------------------------------------------------------------------------||
	now := time.Now()
	exp := time.Date(yyyy, time.Month(mm)+1, 0, 23, 59, 59, 0, now.Location())
	if !exp.After(time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())) {
		return errors.New("Card expired")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| CVC check
	//||------------------------------------------------------------------------------------------------||
	cvcDigits := OnlyDigits(*vc.CVC)
	if len(cvcDigits) < 3 || len(cvcDigits) > 4 || cvcDigits != *vc.CVC {
		return errors.New("Invalid CVC")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Billing ZIP check
	//||------------------------------------------------------------------------------------------------||
	if strings.TrimSpace(*vc.BillingZip) == "" || len(*vc.BillingZip) < 3 {
		return errors.New("Invalid billing ZIP/postal code")
	}

	return nil
}

// BuildDescriptor sanitizes and injects the verification code into a descriptor string
func BuildDescriptor(template, code string) string {
	base := strings.TrimSpace(template)
	if base == "" {
		base = "Payment Verification [XXXX]"
	}
	// Placeholders
	base = strings.ReplaceAll(base, "[XXXX]", code)
	base = strings.ReplaceAll(base, "{CODE}", code)

	// Sanitize
	var b strings.Builder
	for _, r := range base {
		ch := string(r)
		if descAllowed.MatchString(ch) {
			b.WriteString(ch)
		}
	}
	out := strings.TrimSpace(b.String())

	// Trim to max 22 chars
	if len(out) > 22 {
		out = strings.TrimSpace(out[:22])
	}

	// Ensure code is present
	if !strings.Contains(out, code) {
		suffix := "-" + code
		if len(out)+len(suffix) > 22 {
			out = strings.TrimSpace(out[:22-len(suffix)])
		}
		out = strings.TrimSpace(out + suffix)
	}
	return out
}

// OnlyDigits strips all non-digit characters
func OnlyDigits(s string) string {
	b := strings.Builder{}
	for i := 0; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			b.WriteByte(s[i])
		}
	}
	return b.String()
}

// Last4 returns the last 4 digits of a number string
func Last4(digits string) string {
	if len(digits) <= 4 {
		return digits
	}
	return digits[len(digits)-4:]
}

// ToInt safely converts string to int (0 on error)
func ToInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}

// RandRange returns a random integer between min and max (inclusive)
func RandRange(min, max int) (int, error) {
	if min == max {
		return min, nil
	}
	span := max - min + 1
	n, err := rand.Int(rand.Reader, big.NewInt(int64(span)))
	if err != nil {
		return 0, err
	}
	return int(n.Int64()) + min, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Internal
//||------------------------------------------------------------------------------------------------||

func luhn(digits string) bool {
	sum := 0
	alt := false
	for i := len(digits) - 1; i >= 0; i-- {
		n := int(digits[i] - '0')
		if alt {
			n *= 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
		alt = !alt
	}
	return sum%10 == 0
}

var descAllowed = regexp.MustCompile(`[A-Za-z0-9 \.\,\-\_\*]`)
