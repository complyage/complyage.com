package verify

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Data Types
//||------------------------------------------------------------------------------------------------||

type Data struct {
	UAGE DOB            `json:"UAGE,omitempty"`
	MAIL EmailAddress   `json:"MAIL,omitempty"`
	PHNE PhoneNumber    `json:"PHNE,omitempty"`
	ADDR Address        `json:"ADDR,omitempty"`
	CRCD CreditCard     `json:"CRCD,omitempty"`
	IDEN Identification `json:"IDEN,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Init
//||------------------------------------------------------------------------------------------------||

func (e *Encrypted) DataInit() error {
	switch e.Type {
	case DataTypeUAGE:
		e.Data = Data{UAGE: DOB{}}
	case DataTypeMAIL:
		e.Data = Data{MAIL: EmailAddress{}}
	case DataTypePHNE:
		e.Data = Data{PHNE: PhoneNumber{}}
	case DataTypeADDR:
		e.Data = Data{ADDR: Address{}}
	case DataTypeCRCD:
		e.Data = Data{CRCD: CreditCard{}}
	case DataTypeIDEN:
		e.Data = Data{IDEN: Identification{}}
	default:
		return fmt.Errorf("unknown data type: %s", e.Type)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Set Data
//||------------------------------------------------------------------------------------------------||

func (v *Verification) SetDataPhone(phone PhoneNumber) {
	v.Encrypted.Data.PHNE = phone
	v.Display = phone.Mask()
}

func (v *Verification) SetDataEmail(email EmailAddress) {
	v.Encrypted.Data.MAIL = email
	v.Display = email.Mask()
}

func (v *Verification) SetDataDOB(dob DOB) {
	v.Encrypted.Data.UAGE = dob
	v.Display = dob.Mask()
}

func (v *Verification) SetDataIDEN(iden Identification) {
	v.Encrypted.Data.IDEN = iden
	v.Display = iden.Mask()
}

func (v *Verification) SetDataADDR(addr Address) {
	v.Encrypted.Data.ADDR = addr
	v.Display = addr.Mask()
}

func (v *Verification) SetDataCRCD(crcd CreditCard) {
	v.Encrypted.Data.CRCD = crcd
	v.Display = crcd.Mask()
}

//||------------------------------------------------------------------------------------------------||
//|| Get Data but masked for display
//||------------------------------------------------------------------------------------------------||

func (e *Encrypted) GetDataMasked() string {
	switch e.Type {
	case DataTypeUAGE:
		return e.Data.UAGE.Mask()
	case DataTypeMAIL:
		return e.Data.MAIL.Mask()
	case DataTypePHNE:
		return e.Data.PHNE.Mask()
	case DataTypeADDR:
		return e.Data.ADDR.Mask()
	case DataTypeCRCD:
		return e.Data.CRCD.Mask()
	case DataTypeIDEN:
		return e.Data.IDEN.Mask()
	default:
		return "UNKNOWN"
	}
}
