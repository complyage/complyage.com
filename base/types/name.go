package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Name
//||------------------------------------------------------------------------------------------------||

type Name struct {
	First  string `json:"first"`
	Last   string `json:"last"`
	Middle string `json:"middle,omitempty"`
}

func (n *Name) String() string {
	return fmt.Sprintf("%s %s %s", n.First, n.Middle, n.Last)
}

func (n *Name) Mask() string {
	if n == nil {
		return ""
	}
	return fmt.Sprintf("%s %s %s", maskString(n.First), maskString(n.Middle), maskString(n.Last))
}
