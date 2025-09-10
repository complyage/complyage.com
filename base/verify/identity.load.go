package verify

import (
	"encoding/json"
)

//||------------------------------------------------------------------------------------------------||
//|| LoadIdentityFromJSON
//||------------------------------------------------------------------------------------------------||

func LoadIdentityFromJSON(data string) (Identity, error) {
	var iden Identity
	err := json.Unmarshal([]byte(data), &iden)
	if err != nil {
		return Identity{}, err
	}
	return iden, nil
}
