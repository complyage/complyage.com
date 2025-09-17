package types

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Media
//||------------------------------------------------------------------------------------------------||

type Media struct {
	Exists bool   `json:"exists"`
	Size   int64  `json:"size,omitempty"`
	Base64 string `json:"blob,omitempty"`
	Mime   string `json:"mime,omitempty"`
}

func (m *Media) String() string {
	return fmt.Sprintf("Media(mime=%s, size=%d)", m.Mime, m.Size)
}

func (m *Media) Mask() string {
	if m == nil {
		return ""
	}
	return fmt.Sprintf("Media(mime=%s, size=%d)", m.Mime, m.Size)
}
