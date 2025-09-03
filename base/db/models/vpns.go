package models

//||------------------------------------------------------------------------------------------------||
//|| VPN Model
//||------------------------------------------------------------------------------------------------||

type VPN struct {
	IDVPN      int    `gorm:"column:id_vpn;primaryKey;autoIncrement"       json:"id_vpn"`
	Name       string `gorm:"column:vpn_name;size:64"                      json:"vpn_name"`
	URL        string `gorm:"column:vpn_url;size:512"                      json:"vpn_url"`
	Blurb      string `gorm:"column:vpn_blurb;size:256"                    json:"vpn_blurb"`
	Highlights string `gorm:"column:vpn_highlights;size:256"               json:"vpn_highlights"`
	Region     string `gorm:"column:vpn_region;size:64" json:"vpn_region"`
	Price      string `gorm:"column:vpn_price;size:16"                     json:"vpn_price"`
	Rating     int    `gorm:"column:vpn_rating"                            json:"vpn_rating"`
}

//||------------------------------------------------------------------------------------------------||
//|| TableName
//||------------------------------------------------------------------------------------------------||

func (VPN) TableName() string {
	return "vpns"
}
