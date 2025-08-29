package abstract

import (
	"base/db/models"
	"errors"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Location
//||------------------------------------------------------------------------------------------------||

func FetchIPFromDatabase(ipInteger uint32) (string, string, float64, float64, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Convert IP
	//||------------------------------------------------------------------------------------------------||
	var ipRecord models.IP
	err := app.SQLDB["main"].DB.Where("start_ip <= ? AND end_ip >= ?", ipInteger, ipInteger).Order("start_ip DESC").Limit(1).First(&ipRecord)
	if err != nil {
		return "", "", 0, 0, errors.New("IP not found")
	}
	return ipRecord.Country, ipRecord.State, ipRecord.Latitude, ipRecord.Longitude, nil
}
