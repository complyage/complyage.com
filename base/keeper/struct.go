package keeper

//||------------------------------------------------------------------------------------------------||
//|| Keeper Record
//||------------------------------------------------------------------------------------------------||

type KeeperRecord struct {
	KeeperId        string `json:"keeperId"`
	Enforced        bool   `json:"enforced"`
	Verified        bool   `json:"verified"`
	Age             int    `json:"age"`
	UserId          int64  `json:"userId"`
	IPAddress       string `json:"ipAddress"`
	Status          string `json:"status"`
	ClientId        string `json:"clientId"`
	ReturnURL       string `json:"returnUrl,omitempty"`
	UserAccountId   int64  `json:"userAccountId,omitempty"`
	UserAccountName string `json:"userAccountName,omitempty"`
}
