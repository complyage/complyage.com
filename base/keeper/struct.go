package keeper

//||------------------------------------------------------------------------------------------------||
//|| Keeper Record
//||------------------------------------------------------------------------------------------------||

type KeeperRecord struct {
	KeeperId  string `json:"keeperId"`
	Enforced  bool   `json:"enforced"`
	Verified  bool   `json:"verified"`
	Age       int    `json:"age"`
	UserId    int64  `json:"userId"`
	IPAddress string `json:"ipAddress"`
	ClientId  string `json:"clientId"`
	Status    string `json:"status"`
	ReturnURL string `json:"returnUrl,omitempty"`
}
