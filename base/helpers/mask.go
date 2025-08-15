package helpers

import (
	"fmt"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Mask Email
//||------------------------------------------------------------------------------------------------||

func MaskEmail(email string) string {
	parts := strings.Split(email, "@")
	if len(parts) != 2 {
		return "***"
	}
	local := parts[0]
	if len(local) > 2 {
		local = local[:2] + "***"
	} else {
		local = local[:1] + "***"
	}
	return local + "@" + parts[1]
}

//||------------------------------------------------------------------------------------------------||
//|| Mask Phone
//||------------------------------------------------------------------------------------------------||

func MaskPhone(countryCode, number string) string {
	if len(number) > 4 {
		return "+" + countryCode + " ***" + number[len(number)-4:]
	}
	return "+" + countryCode + " ***"
}

//||------------------------------------------------------------------------------------------------||
//|| Mask Address
//||------------------------------------------------------------------------------------------------||

func MaskAddress(street1, city, country string) string {
	return city + ", " + country
}

//||------------------------------------------------------------------------------------------------||
//|| Mask CreditCard
//||------------------------------------------------------------------------------------------------||

func MaskCreditCard(cardType string, last4 string) string {
	upperType := strings.ToUpper(cardType)
	switch upperType {
	case "AMEX", "AMERICAN EXPRESS", "AMEX_CREDIT":
		return "AMEX-" + last4
	case "VISA", "VISA_CREDIT":
		return "VISA-" + last4
	case "MASTERCARD", "MASTERCARD_CREDIT":
		return "MC-" + last4
	case "DISCOVER", "DISCOVER_CREDIT":
		return "DISC-" + last4
	case "DINERS", "DINERS_CREDIT":
		return "DINERS-" + last4
	case "JCB", "JCB_CREDIT":
		return "JCB-" + last4
	case "UNIONPAY", "UNIONPAY_CREDIT":
		return "UNIONPAY-" + last4
	}
	return "**** **** **** " + last4
}

//||------------------------------------------------------------------------------------------------||
//|| Mask Username
//||------------------------------------------------------------------------------------------------||

func MaskUsername(username string, siteID int64) string {
	return username + fmt.Sprintf(" (site:%d)", siteID)
}
