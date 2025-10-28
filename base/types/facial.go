package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Facial
//||------------------------------------------------------------------------------------------------||

type Facial struct {
	DOB      DOB   `json:"dob,omitempty"`
	DOBMatch bool  `json:"dob_match,omitempty"`
	Selfie   Media `json:"selfie,omitempty"`
	Age      int   `json:"age,omitempty"`
	Min      int   `json:"min,omitempty"`
	Max      int   `json:"max,omitempty"`
}

func (f *Facial) String() string {
	return fmt.Sprintf("%04d-%02d-%02d", f.DOB.Year, f.DOB.Month, f.DOB.Day)
}

func (f *Facial) Mask() string {
	if f == nil {
		return ""
	}
	if f.DOB.Year != 0 {
		return fmt.Sprintf("%04d", f.DOB.Year)
	}
	if f.Age != 0 {
		return fmt.Sprintf("%d", f.Age)
	}
	return ""
}
