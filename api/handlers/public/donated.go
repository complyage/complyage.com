package public

import (
	"base/db/models"
	"net/http"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

type MonthlyTotalsResponse struct {
	Donated float64 `json:"donated"`
	Bills   float64 `json:"bills"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler
//||------------------------------------------------------------------------------------------------||

func DonatedHandler(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	firstOfMonth := time.Date(now.Year(), now.Month(), 1, 0, 0, 0, 0, now.Location())

	// Donations = sum of transaction_amount this month
	var donated float64
	if err := app.SQLDB["main"].DB.
		Model(&models.Transactions{}).
		Where("transaction_timestamp >= ? AND transaction_timestamp <= ?", firstOfMonth, now).
		Select("COALESCE(SUM(transaction_amount),0)").
		Scan(&donated).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch donations")
		return
	}

	// Bills = sum of bill_amount this month
	var bills float64
	if err := app.SQLDB["main"].DB.
		Model(&models.Bills{}).
		Where("bill_timestamp >= ? AND bill_timestamp <= ?", firstOfMonth, now).
		Select("COALESCE(SUM(bill_amount),0)").
		Scan(&bills).Error; err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch bills")
		return
	}

	responses.Success(w, http.StatusOK, MonthlyTotalsResponse{
		Donated: donated,
		Bills:   bills,
	})
}
