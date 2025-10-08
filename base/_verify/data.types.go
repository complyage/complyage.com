package verify

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import (
	"encoding/json"
	"fmt"
	"strings"
)

// ||------------------------------------------------------------------------------------------------||
// || DataType (iota-based enum)
// ||------------------------------------------------------------------------------------------------||
type DataType int

const (
	DataTypeMAIL DataType = iota
	DataTypePHNE
	DataTypeADDR
	DataTypeCRCD
	DataTypeIDEN
	DataTypeUSER
	DataTypeFACE
)

// ||------------------------------------------------------------------------------------------------||
// || Lookup Tables
// ||------------------------------------------------------------------------------------------------||
var (
	// list of all values (index matches iota)
	AllDataTypes = []DataType{
		DataTypeMAIL,
		DataTypePHNE,
		DataTypeADDR,
		DataTypeCRCD,
		DataTypeIDEN,
		DataTypeUSER,
		DataTypeFACE,
	}

	// canonical strings for marshal/display
	AllDataTypeStrings = []string{"MAIL", "PHNE", "ADDR", "CRCD", "IDEN", "USER", "FACE"}

	// fast maps for (de)serialization
	dataTypeFromString = map[string]DataType{
		"MAIL": DataTypeMAIL,
		"PHNE": DataTypePHNE,
		"ADDR": DataTypeADDR,
		"CRCD": DataTypeCRCD,
		"IDEN": DataTypeIDEN,
		"USER": DataTypeUSER,
		"FACE": DataTypeFACE,
	}
)

func IsValidDataType(s string) bool {
	_, ok := dataTypeFromString[strings.ToUpper(strings.TrimSpace(s))]
	return ok
}

// ||------------------------------------------------------------------------------------------------||
// || String
// ||------------------------------------------------------------------------------------------------||
func (d DataType) String() string {
	if int(d) >= 0 && int(d) < len(AllDataTypeStrings) {
		return AllDataTypeStrings[d]
	}
	return "UNKNOWN"
}

// ||------------------------------------------------------------------------------------------------||
// || JSON Marshal/Unmarshal
// ||------------------------------------------------------------------------------------------------||
func (d DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DataType) UnmarshalJSON(data []byte) error {
	// 1) Try string (case-insensitive, trims)
	var s string
	if err := json.Unmarshal(data, &s); err == nil {
		key := strings.ToUpper(strings.TrimSpace(s))
		if v, ok := dataTypeFromString[key]; ok {
			*d = v
			return nil
		}
		return fmt.Errorf("invalid DataType: %q (allowed: %s)", s, strings.Join(AllDataTypeStrings, ", "))
	}

	// 2) Try numeric (fallback for legacy numeric encodings)
	var n int
	if err := json.Unmarshal(data, &n); err == nil {
		if n >= 0 && n < len(AllDataTypes) {
			*d = AllDataTypes[n]
			return nil
		}
		return fmt.Errorf("invalid DataType (numeric): %d", n)
	}

	// 3) Unknown type
	return fmt.Errorf("invalid DataType: %s", string(data))
}

// ||------------------------------------------------------------------------------------------------||
// || Text Marshal/Unmarshal (useful for DB/YAML/env)
// ||------------------------------------------------------------------------------------------------||
func (d DataType) MarshalText() ([]byte, error) {
	return []byte(d.String()), nil
}

func (d *DataType) UnmarshalText(text []byte) error {
	key := strings.ToUpper(strings.TrimSpace(string(text)))
	if v, ok := dataTypeFromString[key]; ok {
		*d = v
		return nil
	}
	return fmt.Errorf("invalid DataType: %q (allowed: %s)", string(text), strings.Join(AllDataTypeStrings, ", "))
}

//||------------------------------------------------------------------------------------------------||
//|| Get DataType from String
//||------------------------------------------------------------------------------------------------||

func StringToDataType(s string) (DataType, bool) {
	key := strings.ToUpper(strings.TrimSpace(s))
	v, ok := dataTypeFromString[key]
	return v, ok
}

// ||------------------------------------------------------------------------------------------------||
// || Dot notation namespace for DataType
// ||------------------------------------------------------------------------------------------------||
type nsDataType struct {
	MAIL DataType
	PHNE DataType
	ADDR DataType
	CRCD DataType
	IDEN DataType
	USER DataType
	FACE DataType
}

var DATATYPES = nsDataType{
	MAIL: DataTypeMAIL,
	PHNE: DataTypePHNE,
	ADDR: DataTypeADDR,
	CRCD: DataTypeCRCD,
	IDEN: DataTypeIDEN,
	USER: DataTypeUSER,
	FACE: DataTypeFACE,
}
