package abstract

import (
	"base/db/models"
	"errors"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

func ReturnAllZones() ([]models.Zone, error) {
	var results []models.Zone
	err := app.SQLDB["main"].DB.Table("zones").Order("id_zone ASC").Find(&results).Error
	if err != nil {
		return make([]models.Zone, 0), errors.New("failed to load zones: " + err.Error())
	}
	return results, nil
}
