package admin

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
	"strconv"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| VerificationsWidgetHandler - returns the count of verifications by area
//|| GET /admin/widget/verifications?area=US
//||------------------------------------------------------------------------------------------------||

func VerificationsWidgetHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Get Area
	//||------------------------------------------------------------------------------------------------||

	area := r.URL.Query().Get("filter")
	if area == "" {
		responses.Error(w, http.StatusBadRequest, "Missing area parameter")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Query Database
	//||------------------------------------------------------------------------------------------------||

	var count int64
	err := app.SQLDB["main"].DB.
		Debug(). // REMOVE in production
		Table("verify").
		Where("verify_status = ?", area). // FIXED filter field
		Count(&count).
		Error
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch verification count")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return Count
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]string{
		"count": strconv.FormatInt(count, 10),
		"area":  area,
	})
}
