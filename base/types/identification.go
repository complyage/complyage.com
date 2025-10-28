package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Identification
//||------------------------------------------------------------------------------------------------||

type Identification struct {
	IDType  string  `json:"idType,omitempty"`
	Number  string  `json:"number,omitempty"`
	Front   Media   `json:"front,omitempty"`
	Back    Media   `json:"back,omitempty"`
	Selfie  Media   `json:"selfie,omitempty"`
	Address Address `json:"address,omitempty"`
	DOB     *DOB    `json:"dob,omitempty"`
	Name    *Name   `json:"name,omitempty"`
}

func (i *Identification) String() string {
	return fmt.Sprintf("%s %s (%s)", i.IDType, i.Number, i.Name.String())
}

func (i *Identification) Mask() string {
	if i == nil {
		return ""
	}
	return fmt.Sprintf("%s %s (%s)", i.IDType, maskString(i.Number), i.Name.Mask())
}
