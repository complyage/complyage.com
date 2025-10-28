package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Username
//||------------------------------------------------------------------------------------------------||

type Username struct {
	Username  string `json:"username"`
	FidSite   string `json:"fidSite"`
	Reference Media  `json:"reference"`
	Signed    Media  `json:"signed"`
}

func (u *Username) String() string {
	return fmt.Sprintf("%s@%s", u.Username, u.FidSite)
}

func (u *Username) Mask() string {
	if u == nil {
		return ""
	}
	return fmt.Sprintf("%s@%s", maskString(u.Username), maskString(u.FidSite))
}
