package types

import (
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| DOB
//||------------------------------------------------------------------------------------------------||

type DOB struct {
	Month int `json:"month,omitempty"`
	Day   int `json:"day,omitempty"`
	Year  int `json:"year,omitempty"`
}

func (d *DOB) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
}

func (d *DOB) Mask() string {
	if d == nil {
		return ""
	}
	// mask year fully, keep month/day
	return fmt.Sprintf("%04d-**-**", d.Year)
}

func (d *DOB) Age() int {
	if d == nil || d.Year == 0 || d.Month <= 0 || d.Day <= 0 {
		return 0
	}
	now := time.Now()
	age := now.Year() - d.Year
	if int(now.Month()) < d.Month || (int(now.Month()) == d.Month && now.Day() < d.Day) {
		age--
	}
	if age < 0 {
		return 0
	}
	return age
}
