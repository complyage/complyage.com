package interfaces

import "time"

//||------------------------------------------------------------------------------------------------||
//|| Optimized Zone
//||------------------------------------------------------------------------------------------------||

type OptimizedZone struct {
	State       *string    `json:"state,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Law         *string    `json:"law,omitempty"`
	Description *string    `json:"description,omitempty"`
	Effective   *time.Time `json:"effective,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Optimized Site
//||------------------------------------------------------------------------------------------------||

type OptimizedSite struct {
	Name        string `json:"name"`
	Logo        string `json:"logo"`
	Description string `json:"description"`
	URL         string `json:"url"`
	Redirect    string `json:"redirect"`
	Permissions string `json:"permissions"`
}
