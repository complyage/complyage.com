package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| CreditCard
//||------------------------------------------------------------------------------------------------||

type CreditCard struct {
	LastFour      string  `json:"lastFour,omitempty"`
	CardType      string  `json:"cardType,omitempty"`
	ExpYear       string  `json:"expYear,omitempty"`
	Address       Address `json:"address"`
	TransactionId string  `json:"transactionId,omitempty"`
}

func (c *CreditCard) String() string {
	return fmt.Sprintf("%s ending in %s exp %s", c.CardType, c.LastFour, c.ExpYear)
}

func (c CreditCard) Mask() string {
	return fmt.Sprintf("%s ending in %s exp ****", c.CardType, c.LastFour)
}
