package identity

import "strings"

//||------------------------------------------------------------------------------------------------||
//|| Add Verification
//||------------------------------------------------------------------------------------------------||

func (i *Identity) SetVerification(section string, verified bool, display, verification string) {
	// normalize key
	key := strings.ToUpper(strings.TrimSpace(section))

	// update the appropriate IdentityRecord
	switch key {
	case "ADDR":
		i.Address.Verified = verified
		i.Address.Display = display
		i.Address.Verification = verification
	case "CRCD":
		i.CreditCard.Verified = verified
		i.CreditCard.Display = display
		i.CreditCard.Verification = verification
	case "MAIL":
		i.Email.Verified = verified
		i.Email.Display = display
		i.Email.Verification = verification
	case "FACE":
		i.Face.Verified = verified
		i.Face.Display = display
		i.Face.Verification = verification
	case "IDEN":
		i.IDCard.Verified = verified
		i.IDCard.Display = display
		i.IDCard.Verification = verification
	case "PHNE":
		i.Phone.Verified = verified
		i.Phone.Display = display
		i.Phone.Verification = verification
	}
	i.ApprovedAdd(key)
	i.Save()
}
