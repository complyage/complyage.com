package scopes

import "strings"

//||------------------------------------------------------------------------------------------------||
//|| Scope
//||------------------------------------------------------------------------------------------------||

type BaseScope struct {
	Code        string
	Title       string
	Description string
	Icon        string
	Level       int
}

//||------------------------------------------------------------------------------------------------||
//|| Globals
//||------------------------------------------------------------------------------------------------||

var (
	ScopesList []BaseScope
	ScopesMap  map[string]BaseScope
)

//||------------------------------------------------------------------------------------------------||
//|| Init: Register Scopes
//||------------------------------------------------------------------------------------------------||

func init() {
	ScopesList = []BaseScope{
		{
			Code:        "IDEN",
			Title:       "ID/Passport",
			Description: "Verified ID/Passport",
			Icon:        "user-id",
			Level:       1,
		},
		{
			Code:        "CRCD",
			Title:       "Credit Card",
			Description: "Verified Credit Card",
			Icon:        "credit-card",
			Level:       1,
		},
		{
			Code:        "FACE",
			Title:       "Facial Age",
			Description: "Verified Facial Age",
			Icon:        "smile",
			Level:       0,
		},
		{
			Code:        "MAIL",
			Title:       "Email",
			Description: "Verified Email Address",
			Icon:        "envelope",
			Level:       0,
		},
		{
			Code:        "PHNE",
			Title:       "Phone",
			Description: "Verified Phone Number",
			Icon:        "phone",
			Level:       0,
		},
		{
			Code:        "ADDR",
			Title:       "Mailing Address",
			Description: "Verified Mailing Address",
			Icon:        "home",
			Level:       1,
		},
	}

	ScopesMap = make(map[string]BaseScope, len(ScopesList))
	for _, s := range ScopesList {
		ScopesMap[strings.ToUpper(s.Code)] = s
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Lookup Helpers
//||------------------------------------------------------------------------------------------------||

func FindScope(code string) (BaseScope, bool) {
	s, ok := ScopesMap[strings.ToUpper(code)]
	return s, ok
}

func Title(code string) string {
	if s, ok := FindScope(code); ok {
		return s.Title
	}
	return "Unknown-" + code
}

func Description(code string) string {
	if s, ok := FindScope(code); ok {
		return s.Description
	}
	return "Unknown Scope - " + code
}

func Icon(code string) string {
	if s, ok := FindScope(code); ok {
		return s.Icon
	}
	return "question"
}

func Level(code string) int {
	if s, ok := FindScope(code); ok {
		return s.Level
	}
	return -1
}
