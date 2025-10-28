package identity

import (
	"encoding/json"
)

//||------------------------------------------------------------------------------------------------||
//|| toJSON - Converts Identity struct to JSON string
//||------------------------------------------------------------------------------------------------||

func (i Identity) toJSON() (string, error) {
	b, err := json.Marshal(i)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
