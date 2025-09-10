package models

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Zone
//||------------------------------------------------------------------------------------------------||

type Zone struct {
	ID           uint       `gorm:"column:id_zone;primaryKey;autoIncrement" json:"id"`
	Region       *string    `gorm:"column:zone_state;size:60"               json:"state,omitempty"`
	Country      *string    `gorm:"column:zone_country;size:2"              json:"country,omitempty"`
	Law          *string    `gorm:"column:zone_law;size:255"                json:"law,omitempty"`
	Description  *string    `gorm:"column:zone_law_description;type:text"   json:"description,omitempty"`
	Requirements *string    `gorm:"column:zone_requirements;size:45"        json:"requirements,omitempty"`
	Penalties    *string    `gorm:"column:zone_penalties;size:120"          json:"penalties,omitempty"`
	Effective    *time.Time `gorm:"column:zone_effective;type:date"         json:"effective,omitempty"`
	Meta         *string    `gorm:"column:zone_meta;type:text"              json:"meta,omitempty"`
	Latitude     *string    `gorm:"column:zone_latitude;type:text"          json:"latitude,omitempty"`
	Longitude    *string    `gorm:"column:zone_longitude;type:text"         json:"longitude,omitempty"`
	MinAge       int        `gorm:"column:zone_minage"                      json:"minAge,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table
//||------------------------------------------------------------------------------------------------||

func (Zone) TableName() string {
	return "zones"
}
