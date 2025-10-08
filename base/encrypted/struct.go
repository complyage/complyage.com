package encrypted

import "github.com/complyage/base/types"

type Encrypted struct {
	UUID   string         `json:"uuid"`
	Type   types.DataType `json:"type"`
	Cipher []byte         `json:"cipher"`
}
