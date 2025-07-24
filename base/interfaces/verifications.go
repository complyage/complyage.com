package interfaces

//||------------------------------------------------------------------------------------------------||
//|| OAuth Response Types
//||------------------------------------------------------------------------------------------------||

type UserVerification struct {
	ID     int    `json:"id"`
	Type   string `json:"type"`
	Data   string `json:"data"`
	Status string `json:"status"`
}
