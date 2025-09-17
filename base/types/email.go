package types

import "strings"

//||------------------------------------------------------------------------------------------------||
//|| EmailAddress
//||------------------------------------------------------------------------------------------------||

type EmailAddress struct {
	Email string `json:"email"`
}

func (e *EmailAddress) String() string {
	return e.Email
}

func (e *EmailAddress) Mask() string {
	if e == nil || e.Email == "" {
		return ""
	}
	parts := strings.SplitN(e.Email, "@", 2)
	if len(parts) != 2 {
		return maskString(e.Email)
	}
	return maskString(parts[0]) + "@****"
}
