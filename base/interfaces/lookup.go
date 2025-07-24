package interfaces

//||------------------------------------------------------------------------------------------------||
//|| Request Struct
//||------------------------------------------------------------------------------------------------||

type LookupRequest struct {
	APIKey    string `json:"api_key"`
	IPAddress string `json:"ip_address"`
}

//||------------------------------------------------------------------------------------------------||
//|| Response Struct
//||------------------------------------------------------------------------------------------------||

type LookupResponse struct {
	Requirements string `json:"requirements"`
	Country      string `json:"country"`
	State        string `json:"state"`
}
