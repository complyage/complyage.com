package interfaces

//||------------------------------------------------------------------------------------------------||
//|| OAuth Response Types
//||------------------------------------------------------------------------------------------------||

type UserVerification struct {
	ID     int64  `json:"id"`
	Type   string `json:"type"`
	Data   string `json:"data"`
	Meta   string `json:"meta"`
	Status string `json:"status"`
}
