package types

import (
	"encoding/json"
	"fmt"
)

type IntBool int

func (ib *IntBool) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		if s == "1" {
			*ib = 1
		} else {
			*ib = 0
		}
		return nil
	}

	var i int
	if err := json.Unmarshal(data, &i); err == nil {
		*ib = IntBool(i)
		return nil
	}

	return fmt.Errorf("invalid IntBool input: %s", string(data))
}
