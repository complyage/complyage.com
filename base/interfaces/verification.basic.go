package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Basic Response
//||------------------------------------------------------------------------------------------------||

type VerificationBasicInitialResponse struct {
	UUID   string `json:"uuid"`
	Status string `json:"status,omitempty"`
	Type   string `json:"type,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Basic Media
//||------------------------------------------------------------------------------------------------||

type VerificationMediaResponse struct {
	Exists bool   `json:"exists"`
	Blob   string `json:"blob,omitempty"`
	Mime   string `json:"mime,omitempty"`
	Type   string `json:"type,omitempty"`
}
