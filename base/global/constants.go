package constants

//||------------------------------------------------------------------------------------------------||
//|| Constant Entry
//||------------------------------------------------------------------------------------------------||

type Constant struct {
	Code  string `json:"code"`
	Value any    `json:"value"`
	Info  string `json:"info,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Constant Library
//||------------------------------------------------------------------------------------------------||

type ConstantLibrary struct {
	Reference string
	Entries   map[string]Constant
}

//||------------------------------------------------------------------------------------------------||
//|| Register Constant Helper
//||------------------------------------------------------------------------------------------------||

func (lib *ConstantLibrary) RegisterConstant(code string, value any, info string) {
	lib.Entries[code] = Constant{
		Code:  code,
		Value: value,
		Info:  info,
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Example Libraries
//||------------------------------------------------------------------------------------------------||

var (
	STEP = &ConstantLibrary{
		Reference: "STEP",
		Entries:   map[string]Constant{},
	}
	STATUS = &ConstantLibrary{
		Reference: "STATUS",
		Entries:   map[string]Constant{},
	}
)
