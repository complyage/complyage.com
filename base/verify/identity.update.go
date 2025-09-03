package verify

import (
	"encoding/json"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Add Verification
//||------------------------------------------------------------------------------------------------||

func (i *Identity) AddVerification(section DataType) {
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

func (i *Identity) UpdateVerification(section DataType, display, verification string) {
	switch section {
	case DataTypeADDR:
		i.Address = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeCRCD:
		i.CreditCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeFACE:
		i.Face = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeIDEN:
		i.IDCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeMAIL:
		i.Email = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypePHNE:
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

func (i *Identity) UpdateAge(dataType DataType, dob DOB) {
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if i.VerifiedType == DataTypeIDEN && dataType != DataTypeIDEN {
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier Update
	//||------------------------------------------------------------------------------------------------||
	if i.VerifiedType == DataTypeIDEN && dataType != DataTypeIDEN {
		i.VerifiedDOB = dob
		i.VerifiedType = DataTypeIDEN
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Facial is second Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if dataType == DataTypeFACE && dataType != DataTypeIDEN {
		i.VerifiedType = DataTypeFACE
		i.VerifiedDOB = dob
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Credit Card is third Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	if dataType == DataTypeCRCD && i.VerifiedType != DataTypeIDEN && i.VerifiedType != DataTypeFACE {
		i.VerifiedType = DataTypeCRCD
		day := time.Now().Day()
		month := int(time.Now().Month())
		year := time.Now().Year() - 18
		i.VerifiedDOB = DOB{
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

func (i *Identity) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
