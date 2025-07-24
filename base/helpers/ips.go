package helpers

import (
	"base/db"
	"base/interfaces"
	"base/loaders"
	"base/models"
	"errors"
)

func GetLocationByIP(ipAddress string) (interfaces.Location, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| UseInMemory
	//||------------------------------------------------------------------------------------------------||
	var UseInMemory = IsProduction()
	//||------------------------------------------------------------------------------------------------||
	//|| Location
	//||------------------------------------------------------------------------------------------------||
	location := interfaces.Location{}
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	ipNum := IpToUint32(ipAddress)
	if ipNum == 0 {
		return interfaces.Location{}, errors.New("Invalid IP address - " + ipAddress)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Pull the Country and State
	//||------------------------------------------------------------------------------------------------||
	if UseInMemory {
		ipBlock, found := loaders.FindIPRange(ipNum)
		if !found {
			return interfaces.Location{}, errors.New("IP not found")
		}
		location.Country = ipBlock.Country
		location.State = ipBlock.State
		location.Latitude = ipBlock.Latitude
		location.Longitude = ipBlock.Longitude
	} else {
		var ipRecord models.IP
		if result := db.DB.
			Where("start_ip <= ? AND end_ip >= ?", ipNum, ipNum).
			Order("start_ip DESC").
			Limit(1).
			First(&ipRecord); result.Error != nil {
			return interfaces.Location{}, errors.New("IP not found")
		}
		location.Country = ipRecord.Country
		location.State = ipRecord.State
		location.Latitude = ipRecord.Latitude
		location.Longitude = ipRecord.Longitude
	}
	return location, nil
}
