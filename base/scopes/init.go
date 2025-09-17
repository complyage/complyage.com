package scopes

import "strings"

//||------------------------------------------------------------------------------------------------||
//|| Scope
//||------------------------------------------------------------------------------------------------||

type BaseScope struct {
	Code        string
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
		{Code: "MAIL", Description: "Verified Email Address", Icon: "envelope", Level: 0},
		{Code: "PHNE", Description: "Verified Phone Number", Icon: "phone", Level: 1},
		{Code: "UAGE", Description: "Verified Age", Icon: "calendar", Level: 0},
		{Code: "BDAY", Description: "Verified Birthday", Icon: "cake", Level: 1},
		{Code: "CRCD", Description: "Verified Credit Card", Icon: "credit-card", Level: 0},
		{Code: "PROF", Description: "Verified Profile Photo", Icon: "user-circle", Level: 1},
		{Code: "UNAM", Description: "Verified Username", Icon: "id-card", Level: 0},
		{Code: "ADDR", Description: "Verified Mailing Address", Icon: "home", Level: 1},
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

func Icon(code string) string {
	if s, ok := FindScope(code); ok {
		return s.Icon
	}
	return "question"
}

func Description(code string) string {
	if s, ok := FindScope(code); ok {
		return s.Description
	}
	return "Unknown Scope"
}

func Level(code string) int {
	if s, ok := FindScope(code); ok {
		return s.Level
	}
	return -1
}
