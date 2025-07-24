package models

type IP struct {
	ID         uint    `gorm:"column:id;primaryKey;autoIncrement"`
	StartIP    uint64  `gorm:"column:start_ip;not null"`
	EndIP      uint64  `gorm:"column:end_ip;not null"`
	City       string  `gorm:"column:city;size:200"`
	State      string  `gorm:"column:state;size:100"`
	Country    string  `gorm:"column:country;size:5"`
	PostalCode string  `gorm:"column:postal_code;size:20"`
	Latitude   float64 `gorm:"column:latitude;type:float(10,6)"`
	Longitude  float64 `gorm:"column:longitude;type:float(10,6)"`
}

func (IP) TableName() string {
	return "ips"
}
