package models

//||------------------------------------------------------------------------------------------------||
//|| VPNRating Model
//||------------------------------------------------------------------------------------------------||

type VPNRating struct {
	IDRating int    `gorm:"column:id_rating;primaryKey;autoIncrement"    json:"id_rating"`
	FIDVPN   int    `gorm:"column:fid_vpn"                              json:"fid_vpn"`
	FIDUser  int64  `gorm:"column:fid_user"                             json:"fid_user"`
	Rating   int    `gorm:"column:rating"                               json:"rating"`
	Comment  string `gorm:"column:comment;size:512"                     json:"comment"`
}

// TableName overrides the table name used by GORM.
func (VPNRating) TableName() string {
	return "vpn_ratings"
}
