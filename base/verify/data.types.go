package verify

import (
	"encoding/json"
	"fmt"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| DataType (iota-based enum)
//||------------------------------------------------------------------------------------------------||

type DataType int

const (
	DataTypeUAGE DataType = iota
	DataTypeMAIL
	DataTypePHNE
	DataTypeADDR
	DataTypeCRCD
	DataTypeIDEN
	DataTypeUSER
)

//||------------------------------------------------------------------------------------------------||
//|| Types List (all as slice of DataType and strings)
//||------------------------------------------------------------------------------------------------||

var (
	AllDataTypes = []DataType{
		DataTypeUAGE,
		DataTypeMAIL,
		DataTypePHNE,
		DataTypeADDR,
		DataTypeCRCD,
		DataTypeIDEN,
		DataTypeUSER,
	}

	AllDataTypeStrings = []string{
		"UAGE", "MAIL", "PHNE", "ADDR", "CRCD", "IDEN", "USER",
	}
)

func IsValidDataType(s string) bool {
	s = strings.ToUpper(s)
	for _, v := range AllDataTypeStrings {
		if s == v {
			return true
		}
	}
	return false
}

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (d DataType) String() string {
	switch d {
	case DataTypeUAGE:
		return "UAGE"
	case DataTypeMAIL:
		return "MAIL"
	case DataTypePHNE:
		return "PHNE"
	case DataTypeADDR:
		return "ADDR"
	case DataTypeCRCD:
		return "CRCD"
	case DataTypeIDEN:
		return "IDEN"
	case DataTypeUSER:
		return "USER"
	default:
		return "UNKNOWN"
	}
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (d DataType) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *DataType) UnmarshalJSON(data []byte) error {
	var val string
	if err := json.Unmarshal(data, &val); err != nil {
		return err
	}
	switch val {
	case "UAGE":
		*d = DataTypeUAGE
	case "MAIL":
		*d = DataTypeMAIL
	case "PHNE":
		*d = DataTypePHNE
	case "ADDR":
		*d = DataTypeADDR
	case "CRCD":
		*d = DataTypeCRCD
	case "IDEN":
		*d = DataTypeIDEN
	case "USER":
		*d = DataTypeUSER
	default:
		return fmt.Errorf("invalid DataType: %q", val)
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Dot notation namespace for DataType
//||------------------------------------------------------------------------------------------------||

type nsDataType struct {
	UAGE DataType
	MAIL DataType
	PHNE DataType
	ADDR DataType
	CRCD DataType
	IDEN DataType
	USER DataType
}

var DATATYPES = nsDataType{
	UAGE: DataTypeUAGE,
	MAIL: DataTypeMAIL,
	PHNE: DataTypePHNE,
	ADDR: DataTypeADDR,
	CRCD: DataTypeCRCD,
	IDEN: DataTypeIDEN,
	USER: DataTypeUSER,
}
