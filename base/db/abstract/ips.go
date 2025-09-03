package abstract

import (
	"base/db/models"
	"errors"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Location
//||------------------------------------------------------------------------------------------------||

func FetchIPFromDatabase(ipNum uint32) (city, state, country string, lat, long float64, err error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Cast to signed to match BIGINT (signed) columns
	//||------------------------------------------------------------------------------------------------||

	ipI64 := int64(ipNum)

	//||------------------------------------------------------------------------------------------------||
	//|| Cast to signed to match BIGINT (signed) columns
	//||------------------------------------------------------------------------------------------------||

	var rec models.IP
	db := app.SQLDB["main"].DB

	tx := db.
		Where("? BETWEEN start_ip AND end_ip", ipI64).
		Order("start_ip DESC").
		Limit(1).
		First(&rec)

	if tx.Error != nil {
		return "", "", "", 0, 0, tx.Error
	}
	// Defensive: verify range
	if ipI64 < int64(rec.StartIP) || ipI64 > int64(rec.EndIP) {
		return "", "", "", 0, 0, errors.New("ip not in returned range (sanity check failed)")
	}

	return rec.City, rec.State, rec.Country, rec.Latitude, rec.Longitude, nil
}
