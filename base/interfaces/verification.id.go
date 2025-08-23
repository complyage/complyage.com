package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Struct Matching TypeScript Interface (Partial)
//||------------------------------------------------------------------------------------------------||

type VerificationSteps struct {
	ParsedTextFront bool `json:"parsedTextFront"`
	ParsedTextBack  bool `json:"parsedTextBack"`
	DataParsed      bool `json:"dataParsed"`
	HasDOB          bool `json:"hasDOB"`
	HasName         bool `json:"hasName"`
	HasAddress      bool `json:"hasAddress"`
	FaceMatch       bool `json:"faceMatch"`
	Verified        bool `json:"verified"`
}

type VerificationProcessID struct {
	Identifier string            `json:"identifier"`
	Status     string            `json:"status"`
	Level      string            `json:"level"`
	IDType     string            `json:"idType"`
	Error      *string           `json:"error"`
	AccountID  int64             `json:"accountID"`
	Front      any               `json:"front"` // omit for now
	Back       any               `json:"back"`
	Selfie     any               `json:"selfie"`
	Steps      VerificationSteps `json:"steps"`
}
