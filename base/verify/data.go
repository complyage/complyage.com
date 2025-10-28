package verify

import (
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
//|| Set Data
//||------------------------------------------------------------------------------------------------||

func (v *Verification) SetDataPhone(phone types.PhoneNumber) {
	v.Data.PHNE = phone
	v.Display = phone.Mask()
}

func (v *Verification) SetDataEmail(email types.EmailAddress) {
	v.Data.MAIL = email
	v.Display = email.Mask()
}

func (v *Verification) SetDataFACE(face types.Facial) {
	v.Data.FACE = face
	v.Display = face.Mask()
}

func (v *Verification) SetDataIDEN(iden types.Identification) {
	v.Data.IDEN = iden
	v.Display = iden.Mask()
}

func (v *Verification) SetDataADDR(addr types.Address) {
	v.Data.ADDR = addr
	v.Display = addr.Mask()
}

func (v *Verification) SetDataCRCD(crcd types.CreditCard) {
	v.Data.CRCD = crcd
	v.Display = crcd.Mask()
}

//||------------------------------------------------------------------------------------------------||
//|| Get Data but masked for display
//||------------------------------------------------------------------------------------------------||

func (v *Verification) GetDataMask() string {
	switch v.Type {
	case types.DataTypeFACE:
		return v.Data.FACE.Mask()
	case types.DataTypeMAIL:
		return v.Data.MAIL.Mask()
	case types.DataTypePHNE:
		return v.Data.PHNE.Mask()
	case types.DataTypeADDR:
		return v.Data.ADDR.Mask()
	case types.DataTypeCRCD:
		return v.Data.CRCD.Mask()
	case types.DataTypeIDEN:
		return v.Data.IDEN.Mask()
	default:
		return "UNKNOWN"
	}
}
