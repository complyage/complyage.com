package verify

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Load Identity - Removes ID
//||------------------------------------------------------------------------------------------------||

func (iden *Identity) Save() Identity {
	fmt.Println("Saving Identity Record")
	save := Identity{
		Address:    iden.Address,
		CreditCard: iden.CreditCard,
		Email:      iden.Email,
		Face:       iden.Face,
		IDCard:     iden.IDCard,
		Phone:      iden.Phone,
		Usernames:  iden.Usernames,
		Approved:   iden.Approved,
	}
	return save
}
