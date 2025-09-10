package verify

import (
	"encoding/json"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Add Verification
//||------------------------------------------------------------------------------------------------||

func (i *Identity) AddVerification(section DataType) {
	key := section.String()
	// Convert to set for fast duplicate check
	set := make(map[string]struct{}, len(i.Approved))
	for _, v := range i.Approved {
		set[v] = struct{}{}
	}
	// Add and rebuild the slice
	set[key] = struct{}{}
	i.Approved = make([]string, 0, len(set))
	for v := range set {
		i.Approved = append(i.Approved, v)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Update Verification
//||------------------------------------------------------------------------------------------------||

func (v *Verification) UpdateVerification(section DataType, display, verification string) {
	switch section {
	case DataTypeADDR:
		v.Identity.Address = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeCRCD:
		v.Identity.CreditCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeFACE:
		v.Identity.Face = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeIDEN:
		v.Identity.IDCard = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypeMAIL:
		v.Identity.Email = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	case DataTypePHNE:
		v.Identity.Phone = IdentityRecord{
			Display:      display,
			Verification: verification,
		}
	}
	v.Identity.AddVerification(section)
}

//||------------------------------------------------------------------------------------------------||
//|| VerifyDOB
//||------------------------------------------------------------------------------------------------||

func (v *Verification) UpdateAge(dataType DataType, age int, verification string) {
	//||------------------------------------------------------------------------------------------------||
	//|| Generate DOB
	//||------------------------------------------------------------------------------------------------||
	dobTime := time.Now().AddDate(-age, 0, 0)
	dob := DOB{
		Day:   dobTime.Day(),
		Month: int(dobTime.Month()),
		Year:  dobTime.Year(),
	}
	displayStr := fmt.Sprintf("%d", age)
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	idRecord := IdentityRecord{
		Verified:     true,
		Age:          age,
		DOB:          dob,
		Display:      displayStr,
		Verification: verification,
	}
	fmt.Println("Updating Age to", age, "for", dataType.String(), "->", displayStr)
	fmt.Println("DOB:", dob.String())
	fmt.Println("Verification:", verification)
	fmt.Println(displayStr)
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	switch dataType {
	case DataTypeIDEN:
		v.Identity.IDCard = idRecord
	case DataTypeFACE:
		v.Identity.Face = idRecord
	case DataTypeCRCD:
		v.Identity.CreditCard = idRecord
	}
	v.DatabaseSaveIdentity()
}

//||------------------------------------------------------------------------------------------------||
//|| dob
//||------------------------------------------------------------------------------------------------||

func (v *Verification) UpdateDOB(dataType DataType, dob DOB, verification string) {
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	idRecord := IdentityRecord{
		Verified:     true,
		Age:          dob.Age(),
		DOB:          dob,
		Display:      dob.Mask(),
		Verification: verification,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Iden is top Tier but we got something else
	//||------------------------------------------------------------------------------------------------||
	switch dataType {
	case DataTypeIDEN:
		v.Identity.IDCard = idRecord
	case DataTypeFACE:
		v.Identity.Face = idRecord
	case DataTypeCRCD:
		v.Identity.CreditCard = idRecord
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||
	v.DatabaseSaveIdentity()
}

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (i *Identity) String() string {
	b, _ := json.Marshal(i)
	return string(b)
}
