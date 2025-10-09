package models

import (
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Site Scope
//||------------------------------------------------------------------------------------------------||

type SiteScope struct {
	Code    string `json:"code"`
	Status  string `json:"status"`
	Enabled bool   `json:"enabled"`
}

// Slice wrapper for []SiteScope
type SiteScopes []SiteScope

// Scan implements sql.Scanner
func (s *SiteScopes) Scan(value interface{}) error {
	if value == nil {
		*s = SiteScopes{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan SiteScopes: %v", value)
	}
	return json.Unmarshal(bytes, s)
}

// Value implements driver.Valuer
func (s SiteScopes) Value() (driver.Value, error) {
	if len(s) == 0 {
		return "[]", nil
	}
	return json.Marshal(s)
}

//||------------------------------------------------------------------------------------------------||
//|| Site Zone
//||------------------------------------------------------------------------------------------------||

type SiteZone struct {
	Zone     int  `json:"zone"`
	Enforced bool `json:"enforced"`
}

// Slice wrapper for []SiteZone
type SiteZones []SiteZone

// Scan implements sql.Scanner
func (z *SiteZones) Scan(value interface{}) error {
	if value == nil {
		*z = SiteZones{}
		return nil
	}
	bytes, ok := value.([]byte)
	if !ok {
		return fmt.Errorf("failed to scan SiteZones: %v", value)
	}
	return json.Unmarshal(bytes, z)
}

// Value implements driver.Valuer
func (z SiteZones) Value() (driver.Value, error) {
	if len(z) == 0 {
		return "[]", nil
	}
	return json.Marshal(z)
}

//||------------------------------------------------------------------------------------------------||
//|| Site Model (maps to `sites` table)
//||------------------------------------------------------------------------------------------------||

type Site struct {
	ID         uint  `gorm:"column:id_site;primaryKey;autoIncrement" json:"id"`
	FidAccount int64 `gorm:"column:fid_account;size:36;index"        json:"fid_account"`
	//||------------------------------------------------------------------------------------------------||
	//|| Basic Data
	//||------------------------------------------------------------------------------------------------||
	ClientID string `gorm:"column:site_client_id;size:32;unique"    json:"clientId"`
	TestMode bool   `gorm:"column:site_testmode;"                   json:"testMode"`
	Status   string `gorm:"column:site_status;size:4;index"         json:"status"`
	//||------------------------------------------------------------------------------------------------||
	//|| Identity
	//||------------------------------------------------------------------------------------------------||
	Name        string `gorm:"column:site_name;size:128"               json:"name"`
	Logo        string `gorm:"column:site_logo;size:128"               json:"logo"`
	Description string `gorm:"column:site_description;size:160"        json:"description"`
	URL         string `gorm:"column:site_url;size:255;index"          json:"url"`
	Domains     string `gorm:"column:site_domains;type:text"           json:"domains"`
	//||------------------------------------------------------------------------------------------------||
	//|| Age Gate
	//||------------------------------------------------------------------------------------------------||
	Enforcement string    `gorm:"column:site_enforcement;size:4"          json:"enforcement"`
	Zones       SiteZones `gorm:"column:site_zones;type:text"             json:"zones"`
	Webhook     string    `gorm:"column:site_gate_webhook;size:512"       json:"webhook"`
	CheckKey    string    `gorm:"column:site_gate_checkkey;size:32"       json:"checkKey"`
	//||------------------------------------------------------------------------------------------------||
	//|| OAuth
	//||------------------------------------------------------------------------------------------------||
	Scopes   SiteScopes `gorm:"column:site_scopes;type:text"            json:"scopes"`
	Redirect string     `gorm:"column:site_redirect;size:512"           json:"redirect"`
	//||------------------------------------------------------------------------------------------------||
	//|| Timestamps
	//||------------------------------------------------------------------------------------------------||
	Created time.Time `gorm:"column:site_created;autoCreateTime"      json:"created"`
	Updated time.Time `gorm:"column:site_updated;autoUpdateTime"      json:"updated"`
	//||------------------------------------------------------------------------------------------------||
	//|| Deprecated or Future Use
	//||------------------------------------------------------------------------------------------------||
	ScopeAuto    bool   `gorm:"column:site_scope_auto;type:tinyint(1)"  json:"scopeAuto"`
	AgentPrivate string `gorm:"column:site_agent_private;size:36;uniqueIndex"       json:"agentPrivate"`
	Private      string `gorm:"column:site_private;size:500"            json:"private"`
	Public       string `gorm:"column:site_public;size:300"             json:"public"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (Site) TableName() string {
	return "sites"
}
