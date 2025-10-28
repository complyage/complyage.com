package ips

import (
	"fmt"

	"github.com/complyage/base/db/models"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Array of IP Ranges
//||------------------------------------------------------------------------------------------------||

var IPRanges []IPRange
var IPRangesLoaded bool = false

//||------------------------------------------------------------------------------------------------||
//|| Initial Load of IP ranges from the database into memory
//||------------------------------------------------------------------------------------------------||

func LoadIPRanges() error {
	fmt.Printf("Loading IP Ranges...Please wait..\n")
	var results []models.IP
	if err := app.SQLDB["main"].DB.Order("start_ip ASC").Find(&results).Error; err != nil {
		return err
	}

	IPRanges = make([]IPRange, len(results))
	for i, row := range results {
		IPRanges[i] = IPRange{
			StartIP:   uint32(row.StartIP), // IPv4 only; cast is fine
			EndIP:     uint32(row.EndIP),
			Country:   row.Country,
			State:     row.State,
			City:      row.City, // <<< add this
			Latitude:  row.Latitude,
			Longitude: row.Longitude,
		}
	}
	fmt.Printf("\033[32m[LOAD] - Loaded %d IPs into memory\033[0m\n", len(IPRanges))
	IPRangesLoaded = true
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Function to find the IP range for a given IP number using binary search
//||------------------------------------------------------------------------------------------------||

func FindIPRange(ipNum uint32) (*IPRange, bool) {
	low, high := 0, len(IPRanges)-1
	for low <= high {
		mid := (low + high) / 2
		block := IPRanges[mid]

		if ipNum < block.StartIP {
			high = mid - 1
		} else if ipNum > block.EndIP { // correct comparison
			low = mid + 1
		} else {
			// ipNum is within [StartIP, EndIP]
			return &IPRanges[mid], true
		}
	}
	return nil, false
}
