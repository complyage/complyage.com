package loaders

import (
	"base/db"
	"base/models"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| IP Range entry
//||------------------------------------------------------------------------------------------------||

type IPRange struct {
	StartIP   uint32
	EndIP     uint32
	Country   string
	State     string
	Latitude  float64
	Longitude float64
}

//||------------------------------------------------------------------------------------------------||
//|| Array of IP Ranges
//||------------------------------------------------------------------------------------------------||

var IPRanges []IPRange

//||------------------------------------------------------------------------------------------------||
//|| Initial Load of IP ranges from the database into memory
//||------------------------------------------------------------------------------------------------||

func LoadIPRanges() error {
	var results []models.IP
	if err := db.DB.Order("start_ip ASC").Find(&results).Error; err != nil {
		return err
	}

	IPRanges = make([]IPRange, len(results))
	for i, row := range results {
		IPRanges[i] = IPRange{
			StartIP:   uint32(row.StartIP),
			EndIP:     uint32(row.EndIP),
			Country:   row.Country,
			State:     row.State,
			Latitude:  row.Latitude,
			Longitude: row.Longitude,
		}
	}

	fmt.Printf("Loaded %d IP ranges into memory\n", len(IPRanges))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Function to find the IP range for a given IP number using binary search
//||------------------------------------------------------------------------------------------------||

func FindIPRange(ipNum uint32) (*IPRange, bool) {
	low := 0
	high := len(IPRanges) - 1

	for low <= high {
		mid := (low + high) / 2
		block := IPRanges[mid]

		if ipNum < block.StartIP {
			high = mid - 1
		} else if ipNum > block.StartIP {
			low = mid + 1
		} else {
			// Exact match
			if ipNum <= block.EndIP {
				return &block, true
			}
			break
		}
	}

	// Check final candidate
	if high >= 0 {
		block := IPRanges[high]
		if ipNum >= block.StartIP && ipNum <= block.EndIP {
			return &block, true
		}
	}

	return nil, false
}
