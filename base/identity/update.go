package identity

import (
	"agent/verify"
	"encoding/json"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Add Verification
//||------------------------------------------------------------------------------------------------||

func (i *Identity) AddVerification(section verify.DataType) {
	for _, v := range i.Approved {
		if v == string(section) {
			return
		}
	}
	i.Approved = append(i.Approved, section.String())
}

//||------------------------------------------------------------------------------------------------||
//|| Update Verification
//||------------------------------------------------------------------------------------------------||

func (i *Identity) UpdateVerification(section verify.DataType, display, verification string) {
	switch section {
	case verify.DataTypeADDR:
		i.Address = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case verify.DataTypeCRCD:
		i.CreditCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case verify.DataTypeFACE:
		i.Face = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case verify.DataTypeIDEN:
		i.IDCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case verify.DataTypeMAIL:
		i.Email = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case verify.DataTypePHNE:
		i.Phone = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	}
	i.AddVerification(section)
	i.Save()
}

//||------------------------------------------------------------------------------------------------||
//|| VerifyAge
//||------------------------------------------------------------------------------------------------||

func (i *Identity) UpdateAge(dataType verify.DataType, dob verify.DOB) {
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if i.VerifiedType == verify.DataTypeIDEN && dataType != verify.DataTypeIDEN {
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier Update
	//||------------------------------------------------------------------------------------------------||
	if i.VerifiedType == verify.DataTypeIDEN && dataType != verify.DataTypeIDEN {
		i.VerifiedDOB = dob
		i.VerifiedType = verify.DataTypeIDEN
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Facial is second Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if dataType == verify.DataTypeFACE && dataType != verify.DataTypeIDEN {
		i.VerifiedType = verify.DataTypeFACE
		i.VerifiedDOB = dob
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Credit Card is third Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if dataType == verify.DataTypeCRCD && i.VerifiedType != verify.DataTypeIDEN && i.VerifiedType != verify.DataTypeFACE {
		i.VerifiedType = verify.DataTypeCRCD
		day := time.Now().Day()
		month := int(time.Now().Month())
		year := time.Now().Year() - 18
		i.VerifiedDOB = verify.DOB{
			Day:   day,
			Month: month,
			Year:  year,
		}
		return
	}

}

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (i Identity) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
