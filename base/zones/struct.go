package zones

import "base/verify"

type ShortZone struct {
	Law          string            `json:"laws,omitempty"`
	Requirements []verify.DataType `json:"requirements,omitempty"`
	Effective    string            `json:"effective,omitempty"`
	MinAge       int               `json:"minAge,omitempty"`
}
