package zones

import "github.com/complyage/base/verify"

type ShortZone struct {
	ID           int               `json:"id,omitempty"`
	Law          string            `json:"laws,omitempty"`
	Requirements []verify.DataType `json:"requirements,omitempty"`
	Effective    string            `json:"effective,omitempty"`
	MinAge       int               `json:"minAge,omitempty"`
}

type SiteZone struct {
	ID       int64 `json:"id,omitempty"`
	Enforced bool  `json:"enforced,omitempty"`
}
