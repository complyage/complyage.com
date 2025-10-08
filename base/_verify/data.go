package verify

import (
	"fmt"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| Data Types
//||------------------------------------------------------------------------------------------------||

type Data struct {
	FACE types.Facial         `json:"FACE,omitempty"`
	MAIL types.EmailAddress   `json:"MAIL,omitempty"`
	PHNE types.PhoneNumber    `json:"PHNE,omitempty"`
	ADDR types.Address        `json:"ADDR,omitempty"`
	CRCD types.CreditCard     `json:"CRCD,omitempty"`
	IDEN types.Identification `json:"IDEN,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Init
//||------------------------------------------------------------------------------------------------||

func (e *Encrypted) DataInit() error {
	switch e.Type {
	case DataTypeFACE:
		e.Data = Data{FACE: types.Facial{}}
	case DataTypeMAIL:
		e.Data = Data{MAIL: types.EmailAddress{}}
	case DataTypePHNE:
		e.Data = Data{PHNE: types.PhoneNumber{}}
	case DataTypeADDR:
		e.Data = Data{ADDR: types.Address{}}
	case DataTypeCRCD:
		e.Data = Data{CRCD: types.CreditCard{}}
	case DataTypeIDEN:
		e.Data = Data{IDEN: types.Identification{}}
	default:
		return fmt.Errorf("unknown data type: %s", e.Type)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Set Data
//||------------------------------------------------------------------------------------------------||

func (v *Verification) SetDataPhone(phone types.PhoneNumber) {
	v.Encrypted.Data.PHNE = phone
	v.Display = phone.Mask()
}

func (v *Verification) SetDataEmail(email types.EmailAddress) {
	v.Encrypted.Data.MAIL = email
	v.Display = email.Mask()
}

func (v *Verification) SetDataFACE(face types.Facial) {
	v.Encrypted.Data.FACE = face
	v.Display = face.Mask()
}

func (v *Verification) SetDataIDEN(iden types.Identification) {
	v.Encrypted.Data.IDEN = iden
	v.Display = iden.Mask()
}

func (v *Verification) SetDataADDR(addr types.Address) {
	v.Encrypted.Data.ADDR = addr
	v.Display = addr.Mask()
}

func (v *Verification) SetDataCRCD(crcd types.CreditCard) {
	v.Encrypted.Data.CRCD = crcd
	v.Display = crcd.Mask()
}

//||------------------------------------------------------------------------------------------------||
//|| Get Data but masked for display
//||------------------------------------------------------------------------------------------------||

func (e *Encrypted) GetDataMasked() string {
	switch e.Type {
	case DataTypeFACE:
		return e.Data.FACE.Mask()
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
