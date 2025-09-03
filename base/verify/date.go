package verify

import (
	"encoding/json"
	"fmt"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Layout
//||------------------------------------------------------------------------------------------------||

const universalDateLayout = "2006-01-02 15:04:05"

//||------------------------------------------------------------------------------------------------||
//|| UniversalDate Type
//||------------------------------------------------------------------------------------------------||

type UniversalDate struct {
	time.Time
}

//||------------------------------------------------------------------------------------------------||
//|| To Universal
//||------------------------------------------------------------------------------------------------||

func ToUniversal(t time.Time) UniversalDate {
	return UniversalDate{Time: t.UTC()}
}

//||------------------------------------------------------------------------------------------------||
//|| From Universal
//||------------------------------------------------------------------------------------------------||

func FromUniversal(t UniversalDate) time.Time {
	return t.Time
}

//||------------------------------------------------------------------------------------------------||
//|| FromUniversalString: Create a UniversalDate from a string
//||------------------------------------------------------------------------------------------------||

func FromUniversalString(str string) UniversalDate {
	t, err := time.Parse(universalDateLayout, str)
	if err != nil {
		return UniversalDate{Time: time.Unix(0, 0).UTC()}
	}
	return UniversalDate{Time: t.UTC()}
}

//||------------------------------------------------------------------------------------------------||
//|| To UniversalString
//||------------------------------------------------------------------------------------------------||

func ToUniversalString(u UniversalDate) string {
	return u.UTC().Format(universalDateLayout)
}

//||------------------------------------------------------------------------------------------------||
//|| JSON Marshal/Unmarshal
//||------------------------------------------------------------------------------------------------||

func (u UniversalDate) MarshalJSON() ([]byte, error) {
	str := u.UTC().Format(universalDateLayout)
	return json.Marshal(str)
}

//||------------------------------------------------------------------------------------------------||
//|| Now
//||------------------------------------------------------------------------------------------------||

func UniversalNow() UniversalDate {
	return UniversalDate{Time: time.Now().UTC()}
}

//||------------------------------------------------------------------------------------------------||
//|| UnmarshalJSON
//||------------------------------------------------------------------------------------------------||

func (u *UniversalDate) UnmarshalJSON(data []byte) error {
	var s string
	if err := json.Unmarshal(data, &s); err != nil {
		return err
	}
	t, err := time.Parse(universalDateLayout, s)
	if err != nil {
		return fmt.Errorf("invalid universal date: %w", err)
	}
	u.Time = t.UTC()
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| String
//||------------------------------------------------------------------------------------------------||

func (u UniversalDate) String() string {
	return u.UTC().Format(universalDateLayout)
}
