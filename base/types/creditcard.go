package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| CreditCard
//||------------------------------------------------------------------------------------------------||

type CreditCard struct {
	LastFour      string  `json:"lastFour"`
	CardType      string  `json:"cardType"`
	ExpYear       string  `json:"expYear"`
	Address       Address `json:"address"`
	DOB           DOB     `json:"dob"`
	TransactionId string  `json:"transactionId"`
}

func (c *CreditCard) String() string {
	return fmt.Sprintf("%s ending in %s exp %s", c.CardType, c.LastFour, c.ExpYear)
}

func (c CreditCard) Mask() string {
	return fmt.Sprintf("%s ending in %s exp ****", c.CardType, c.LastFour)
}
