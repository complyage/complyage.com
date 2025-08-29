package ips

import (
	"base/db/abstract"
	"errors"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/convert"
)

//||------------------------------------------------------------------------------------------------||
//|| Convert IP
//||------------------------------------------------------------------------------------------------||

func GetLocationByIP(ipAddress string) (Location, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	ipNum := convert.IpToUint32(ipAddress)
	if ipNum == 0 {
		return Location{}, errors.New("Invalid IP address - " + ipAddress)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Dev mode - always hit the database
	//||------------------------------------------------------------------------------------------------||
	if app.Config.App.Env != "production" {
		city, state, lat, long, err := abstract.FetchIPFromDatabase(ipNum)
		if err != nil {
			return Location{}, err
		}
		var location Location
		location.Country = city
		location.State = state
		location.Latitude = lat
		location.Longitude = long
		return location, nil
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Country and State
	//||------------------------------------------------------------------------------------------------||
	ipBlock, found := FindIPRange(ipNum)
	if !found {
		return Location{}, errors.New("IP not found")
	}
	var location Location
	location.Country = ipBlock.Country
	location.State = ipBlock.State
	location.Latitude = ipBlock.Latitude
	location.Longitude = ipBlock.Longitude
	return location, nil
}
