package ips

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/abstract"
	"errors"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/convert"
)

//||------------------------------------------------------------------------------------------------||
//|| GetLocationByIP
//||------------------------------------------------------------------------------------------------||

func GetLocationByIP(ipAddress string) (Location, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||

	ipNum := convert.IpToUint32(ipAddress)
	if ipNum == 0 {
		return Location{}, errors.New("invalid IPv4 address: " + ipAddress)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Dev Mode: DB Lookup
	//||------------------------------------------------------------------------------------------------||

	if app.Config.App.Env != "production" {
		city, state, country, lat, long, err := abstract.FetchIPFromDatabase(ipNum)
		if err != nil {
			return Location{}, err
		}
		return Location{
			City:      city,
			State:     state,
			Country:   country,
			Latitude:  lat,
			Longitude: long,
		}, nil
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Prod Mode: In-Memory Ranges
	//||------------------------------------------------------------------------------------------------||

	ipBlock, ok := FindIPRange(ipNum)
	if !ok {
		return Location{}, errors.New("ip not found in range set")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return
	//||------------------------------------------------------------------------------------------------||

	return Location{
		City:      ipBlock.City,
		State:     ipBlock.State,
		Country:   ipBlock.Country,
		Latitude:  ipBlock.Latitude,
		Longitude: ipBlock.Longitude,
	}, nil
}
