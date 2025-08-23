package verify

import (
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| ModerateType (iota-based enum)
//||------------------------------------------------------------------------------------------------||

type ModerateType int

const (
	ModerateTypeTwoFactor ModerateType = iota
	ModerateTypeAILevel1
	ModerateTypeAILevel2
	ModerateTypeHuman
)

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (t ModerateType) String() string {
	switch t {
	case ModerateTypeTwoFactor:
		return "TWO_FACTOR"
	case ModerateTypeAILevel1:
		return "AI_LEVEL_1"
	case ModerateTypeAILevel2:
		return "AI_LEVEL_2"
	case ModerateTypeHuman:
		return "HUMAN"
	default:
		return "UNKNOWN"
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Marshal/Unmarshal JSON
//||------------------------------------------------------------------------------------------------||

func (t ModerateType) MarshalJSON() ([]byte, error) {
	return json.Marshal(t.String())
}

//||------------------------------------------------------------------------------------------------||
//|| Unmarshal JSON
//||------------------------------------------------------------------------------------------------||

func (t *ModerateType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "TWO_FACTOR":
		*t = ModerateTypeTwoFactor
	case "AI_LEVEL_1":
		*t = ModerateTypeAILevel1
	case "AI_LEVEL_2":
		*t = ModerateTypeAILevel2
	case "HUMAN":
		*t = ModerateTypeHuman
	default:
		return fmt.Errorf("invalid ModerateType: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Namespace for dot notation
//||------------------------------------------------------------------------------------------------||

type nsModerateType struct {
	TwoFactor ModerateType
	AILevel1  ModerateType
	AILevel2  ModerateType
	Human     ModerateType
}

var MODERATE_TYPE = nsModerateType{
	TwoFactor: ModerateTypeTwoFactor,
	AILevel1:  ModerateTypeAILevel1,
	AILevel2:  ModerateTypeAILevel2,
	Human:     ModerateTypeHuman,
}
