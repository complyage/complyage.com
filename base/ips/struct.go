package ips

import (
	"time"

	"github.com/complyage/base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| IP Range entry
//||------------------------------------------------------------------------------------------------||

type IPRange struct {
	StartIP   uint32
	EndIP     uint32
	Country   string
	State     string
	City      string
	Latitude  float64
	Longitude float64
}

//||------------------------------------------------------------------------------------------------||
//|| Struct for Response
//||------------------------------------------------------------------------------------------------||

type IPLocationVerificationResponse struct {
	IPAddress string `json:"ipAddress"`
	City      string `json:"city"`
	Region    string `json:"region"`
	Country   string `json:"country"`
	Types     string `json:"types"`
	MinAge    int    `json:"minAge,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Location
//||------------------------------------------------------------------------------------------------||

type Location struct {
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Country   string  `json:"country"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

//||------------------------------------------------------------------------------------------------||
//|| Optimized Zone
//||------------------------------------------------------------------------------------------------||

type OptimizedZone struct {
	Region      *string    `json:"region,omitempty"`
	Country     *string    `json:"country,omitempty"`
	Law         *string    `json:"law,omitempty"`
	Description *string    `json:"description,omitempty"`
	Effective   *time.Time `json:"effective,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Optimized Site
//||------------------------------------------------------------------------------------------------||

type OptimizedSite struct {
	Name        string             `json:"name"`
	Logo        string             `json:"logo"`
	Description string             `json:"description"`
	URL         string             `json:"url"`
	Redirect    string             `json:"redirect"`
	Scopes      []models.SiteScope `json:"permissions"`
}
