package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Address
//||------------------------------------------------------------------------------------------------||

type Address struct {
	Line1   string `json:"line1,omitempty"`
	Line2   string `json:"line2,omitempty"`
	City    string `json:"city,omitempty"`
	State   string `json:"state,omitempty"`
	Postal  string `json:"postal,omitempty"`
	Country string `json:"country,omitempty"`
}

func (a *Address) String() string {
	return fmt.Sprintf("%s %s, %s, %s %s, %s",
		a.Line1, a.Line2, a.City, a.State, a.Postal, a.Country)
}

func (a *Address) Mask() string {
	if a == nil {
		return ""
	}
	return fmt.Sprintf("%s, %s, %s",
		maskString(a.City), maskString(a.State), maskString(a.Country))
}
