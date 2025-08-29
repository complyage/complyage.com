package models

import (
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Zone
//||------------------------------------------------------------------------------------------------||

type Zone struct {
	IDZone             uint       `gorm:"column:id_zone;primaryKey;autoIncrement" json:"id"`
	ZoneState          *string    `gorm:"column:zone_state;size:60"               json:"state,omitempty"`
	ZoneCountry        *string    `gorm:"column:zone_country;size:2"              json:"country,omitempty"`
	ZoneLaw            *string    `gorm:"column:zone_law;size:255"                json:"law,omitempty"`
	ZoneLawDescription *string    `gorm:"column:zone_law_description;type:text"   json:"description,omitempty"`
	ZoneRequirements   *string    `gorm:"column:zone_requirements;size:45"        json:"requirements,omitempty"`
	ZonePenalties      *string    `gorm:"column:zone_penalties;size:120"          json:"penalties,omitempty"`
	ZoneEffective      *time.Time `gorm:"column:zone_effective;type:date"         json:"effective,omitempty"`
	ZoneMeta           *string    `gorm:"column:zone_meta;type:text"              json:"meta,omitempty"`
	ZoneLatitude       *string    `gorm:"column:zone_latitude;type:text"          json:"latitude,omitempty"`
	ZoneLongitude      *string    `gorm:"column:zone_longitude;type:text"         json:"longitude,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Table
//||------------------------------------------------------------------------------------------------||

func (Zone) TableName() string {
	return "zones"
}
