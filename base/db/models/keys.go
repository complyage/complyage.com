package models

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| ModelKey
//||------------------------------------------------------------------------------------------------||

type ModelKey struct {
	IDKey      uint      `gorm:"column:id_key;primaryKey;autoIncrement" json:"id_key"`
	FidAccount uint      `gorm:"column:fid_account;not null"            json:"fid_account"`
	Level      int       `gorm:"column:key_level;type:text;not null"  json:"level"`
	Private    string    `gorm:"column:key_private;type:text;not null"  json:"-"`
	Public     string    `gorm:"column:key_public;type:text;not null"   json:"public_key"`
	CheckKey   string    `gorm:"column:key_check;size:64;not null"      json:"check_key"`
	CreatedAt  time.Time `gorm:"column:key_created;autoCreateTime"       json:"created_at"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (ModelKey) TableName() string {
	return "keys"
}
