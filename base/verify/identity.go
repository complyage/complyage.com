package verify

//||------------------------------------------------------------------------------------------------||
//|| Identity
//||------------------------------------------------------------------------------------------------||

type Identity struct {
	MAIL IdentityOption `json:"email,omitempty"`
	FACE IdentityOption `json:"face,omitempty"`
	PHNE IdentityOption `json:"phone,omitempty"`
	CRCD IdentityOption `json:"credit_card,omitempty"`
	ADDR IdentityOption `json:"address,omitempty"`
	IDEN IdentityOption `json:"id_card,omitempty"`
	LIST []DataType     `json:"approved,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Identity
//||------------------------------------------------------------------------------------------------||

type IdentityOption struct {
	Display          string `json:"display,omitempty"`
	VerificationUUID string `json:"verification,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Update Identity
//||------------------------------------------------------------------------------------------------||

func (v *Verification) UpdateIdentity() {
	//||------------------------------------------------------------------------------------------------||
	//|| Update Identity Record
	//||------------------------------------------------------------------------------------------------||
	switch v.Type {
	case DataTypeFACE:
		v.Identity.FACE = IdentityOption{
			Display:          v.Encrypted.Data.FACE.Mask(),
			VerificationUUID: v.UUID,
		}
	case DataTypeMAIL:
		v.Identity.MAIL = IdentityOption{
			Display:          v.Encrypted.Data.MAIL.Mask(),
			VerificationUUID: v.UUID,
		}
	case DataTypePHNE:
		v.Identity.PHNE = IdentityOption{
			Display:          v.Encrypted.Data.PHNE.Mask(),
			VerificationUUID: v.UUID,
		}
	case DataTypeADDR:
		v.Identity.ADDR = IdentityOption{
			Display:          v.Encrypted.Data.ADDR.Mask(),
			VerificationUUID: v.UUID,
		}
	case DataTypeCRCD:
		v.Identity.CRCD = IdentityOption{
			Display:          v.Encrypted.Data.CRCD.Mask(),
			VerificationUUID: v.UUID,
		}
	case DataTypeIDEN:
		v.Identity.IDEN = IdentityOption{
			Display:          v.Encrypted.Data.IDEN.Mask(),
			VerificationUUID: v.UUID,
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Add to List if not already there
	//||------------------------------------------------------------------------------------------------||
	v.Identity.LIST = append(v.Identity.LIST, v.Type)
}

//||------------------------------------------------------------------------------------------------||
//|| Remove Identity
//||------------------------------------------------------------------------------------------------||

func (v *Verification) RemoveIdentity(dataType DataType) {
	//||------------------------------------------------------------------------------------------------||
	//|| Update Identity Record
	//||------------------------------------------------------------------------------------------------||
	switch dataType {
	case DataTypeFACE:
		v.Identity.FACE = IdentityOption{}
	case DataTypeMAIL:
		v.Identity.MAIL = IdentityOption{}
	case DataTypePHNE:
		v.Identity.PHNE = IdentityOption{}
	case DataTypeADDR:
		v.Identity.ADDR = IdentityOption{}
	case DataTypeCRCD:
		v.Identity.CRCD = IdentityOption{}
	case DataTypeIDEN:
		v.Identity.IDEN = IdentityOption{}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Remove from List if there
	//||------------------------------------------------------------------------------------------------||
	newList := []DataType{}
	for _, dt := range v.Identity.LIST {
		if dt != dataType {
			newList = append(newList, dt)
		}
	}
	v.Identity.LIST = newList
}
