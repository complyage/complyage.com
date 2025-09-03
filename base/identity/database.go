package identity

//||------------------------------------------------------------------------------------------------||
//|| Load Identity - Removes ID
//||------------------------------------------------------------------------------------------------||

func (iden *Identity) Save() Identity {
	save := Identity{
		Address:      iden.Address,
		CreditCard:   iden.CreditCard,
		Email:        iden.Email,
		Face:         iden.Face,
		IDCard:       iden.IDCard,
		Phone:        iden.Phone,
		Usernames:    iden.Usernames,
		Approved:     iden.Approved,
		VerifiedDOB:  iden.VerifiedDOB,
		VerifiedType: iden.VerifiedType,
	}
	return save
}
