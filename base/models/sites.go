package models

import (
	"base/types"
	"time"

	"gorm.io/datatypes"
)

//||------------------------------------------------------------------------------------------------||
//|| Site Model (maps to `sites` table)
//||------------------------------------------------------------------------------------------------||

type Site struct {
	IDSite     uint   `gorm:"column:id_site;primaryKey;autoIncrement" json:"id"`
	FidAccount string `gorm:"column:fid_account;size:36;index"          json:"fid_account"`

	SiteName        string `gorm:"column:site_name;size:128"                 json:"name"`
	SiteLogo        string `gorm:"column:site_logo;size:128"                 json:"logo"`
	SiteDescription string `gorm:"column:site_description;size:160"          json:"description"`
	SiteURL         string `gorm:"column:site_url;size:255;index"            json:"url"`
	SiteStatus      string `gorm:"column:site_status;size:4;index"           json:"status"`

	SiteEnforcement string            `gorm:"column:site_enforcement;size:4"            json:"enforcement"`
	SiteZones       datatypes.JSONMap `gorm:"column:site_zones;type:text"               json:"zones"`

	SiteDomains     string `gorm:"column:site_domains;type:text"             json:"domains"`
	SitePublic      string `gorm:"column:site_public;size:64;uniqueIndex"    json:"public"`
	SitePrivate     string `gorm:"column:site_private;size:64;uniqueIndex"   json:"private"`
	SiteRedirect    string `gorm:"column:site_redirect;size:256"             json:"redirect"`
	SitePermissions string `gorm:"column:site_permissions;type:text"         json:"permissions"`

	SiteTestMode bool `gorm:"column:site_testmode;default:1"            json:"testmode"`

	// new gate fields
	SiteGateSignup  types.IntBool `gorm:"column:site_gate_signup;default:1" 	json:"gateSignup"`
	SiteGateConfirm string        `gorm:"column:site_gate_confirm;size:256"       json:"gateConfirm"`
	SiteGateExit    string        `gorm:"column:site_gate_exit;size:256"          json:"gateExit"`

	SiteCreated time.Time `gorm:"column:site_created;autoCreateTime"        json:"created"`
	SiteUpdated time.Time `gorm:"column:site_updated;autoUpdateTime"        json:"updated"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (Site) TableName() string {
	return "sites"
}
