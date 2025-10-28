package models

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| ModelSecret
//||------------------------------------------------------------------------------------------------||

type ModelSecrets struct {
	IDSecret   uint      `gorm:"column:id_secret;primaryKey;autoIncrement" json:"id_secret"`
	FidAccount uint      `gorm:"column:fid_account;not null"               json:"fid_account"`
	Level      int       `gorm:"column:secret_level;not null"              json:"level"`
	Private    string    `gorm:"column:secret_private;type:text;not null"  json:"-"`
	Public     string    `gorm:"column:secret_public;type:text;not null"   json:"public_key"`
	CheckKey   string    `gorm:"column:secret_check;size:64;not null"      json:"check_key"`
	CreatedAt  time.Time `gorm:"column:secret_created;autoCreateTime"      json:"created_at"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table Name
//||------------------------------------------------------------------------------------------------||

func (ModelSecrets) TableName() string {
	return "secrets"
}
