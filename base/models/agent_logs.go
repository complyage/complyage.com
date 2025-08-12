package models

import "time"

type AgentRequestLog struct {
	IDRequest        int64     `gorm:"column:id_request;primaryKey;autoIncrement"`
	FidSite          int64     `gorm:"column:fid_site"`
	RequestLevel     int       `gorm:"column:request_level"`
	RequestID        string    `gorm:"column:request_id"`
	RequestTimestamp time.Time `gorm:"column:request_timestamp;autoCreateTime"`
}

// TableName overrides default pluralization
func (AgentRequestLog) TableName() string {
	return "requests"
}
