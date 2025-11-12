package abstract

import (
	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

func ReturnAllVPNs() ([]models.VPN, error) {
	var vpns []models.VPN
	result := app.SQLDB["main"].DB.Find(&vpns)
	return vpns, result.Error
}
