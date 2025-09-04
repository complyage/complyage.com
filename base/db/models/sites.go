package models

import (
	"time"

	"gorm.io/datatypes"
)

//||------------------------------------------------------------------------------------------------||
//|| Site Model (maps to `sites` table)
//||------------------------------------------------------------------------------------------------||

type Site struct {
	ID           uint              `gorm:"column:id_site;primaryKey;autoIncrement"       json:"id"`
	FidAccount   int64             `gorm:"column:fid_account;size:36;index"              json:"fid_account"`
	Name         string            `gorm:"column:site_name;size:128"                     json:"name"`
	Logo         string            `gorm:"column:site_logo;size:128"                     json:"logo"`
	Description  string            `gorm:"column:site_description;size:160"              json:"description"`
	URL          string            `gorm:"column:site_url;size:255;index"                json:"url"`
	Status       string            `gorm:"column:site_status;size:4;index"               json:"status"`
	Enforcement  string            `gorm:"column:site_enforcement;size:4"                json:"enforcement"`
	Zones        datatypes.JSONMap `gorm:"column:site_zones;type:text"                   json:"zones"`
	Domains      string            `gorm:"column:site_domains;type:text"                 json:"domains"`
	Public       string            `gorm:"column:site_public;size:64;uniqueIndex"        json:"public"`
	Private      string            `gorm:"column:site_private;size:64;uniqueIndex"       json:"private"`
	Redirect     string            `gorm:"column:site_redirect;size:256"                 json:"redirect"`
	Permissions  string            `gorm:"column:site_permissions;type:text"             json:"permissions"`
	TestMode     bool              `gorm:"column:site_testmode;"                         json:"testmode"`
	GateSignup   int64             `gorm:"column:site_gate_signup;"                      json:"gateSignup"`
	GateConfirm  string            `gorm:"column:site_gate_confirm;size:256"             json:"gateConfirm"`
	GateExit     string            `gorm:"column:site_gate_exit;size:256"                json:"gateExit"`
	Created      time.Time         `gorm:"column:site_created;autoCreateTime"            json:"created"`
	Updated      time.Time         `gorm:"column:site_updated;autoUpdateTime"            json:"updated"`
	AgentPrivate string            `gorm:"column:site_agent_private;size:64"             json:"agentPrivate"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (Site) TableName() string {
	return "sites"
}
