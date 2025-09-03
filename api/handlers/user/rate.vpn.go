package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db/models"
	"errors"
	"net/http"
	"strconv"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/responses"
	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: VPN Rating (Insert/Overwrite/Remove)
//||------------------------------------------------------------------------------------------------||

func VPNRatingUserHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse VPN ID
	//||------------------------------------------------------------------------------------------------||

	vpnStr := r.URL.Query().Get("vpn")
	vpnID, err := strconv.Atoi(vpnStr)
	if err != nil || vpnID <= 0 {
		responses.Error(w, http.StatusBadRequest, "Invalid VPN ID")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Rating
	//||------------------------------------------------------------------------------------------------||

	ratingStr := r.URL.Query().Get("rating")
	var rating int
	if ratingStr == "" {
		rating = 0
	} else {
		if ratingStr != "1" && ratingStr != "2" && ratingStr != "3" && ratingStr != "4" && ratingStr != "5" {
			responses.Error(w, http.StatusBadRequest, "Invalid rating")
			return
		}
		rating, err = strconv.Atoi(ratingStr)
		if err != nil || rating < 1 || rating > 5 {
			responses.Error(w, http.StatusBadRequest, "Invalid rating")
			return
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Session Cookie
	//||------------------------------------------------------------------------------------------------||

	_, account, _, uErr := actions.LoadSessionAccount(r)
	if uErr != nil {
		responses.Error(w, http.StatusUnauthorized, "No session cookie")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| DB
	//||------------------------------------------------------------------------------------------------||

	db := app.SQLDB["main"].DB

	//||------------------------------------------------------------------------------------------------||
	//|| Delete Rating if Blank or 0, Otherwise Upsert
	//||------------------------------------------------------------------------------------------------||

	if rating <= 0 {
		db.Where("fid_vpn = ? AND fid_user = ?", vpnID, account.ID).Delete(&models.VPNRating{})
	} else {
		var vr models.VPNRating
		res := db.Where("fid_vpn = ? AND fid_user = ?", vpnID, account.ID).First(&vr)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			vr = models.VPNRating{
				FIDVPN:  vpnID,
				FIDUser: account.ID,
				Rating:  rating,
			}
			db.Create(&vr)
		} else if res.Error == nil {
			vr.Rating = rating
			db.Save(&vr)
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Recalculate Rating as Percent (0-100)
	//||------------------------------------------------------------------------------------------------||

	var sum, count int
	db.Model(&models.VPNRating{}).
		Where("fid_vpn = ? AND rating IS NOT NULL", vpnID).
		Select("SUM(rating), COUNT(1)").Row().Scan(&sum, &count)

	var avgRating int
	if count > 0 {
		avg := float64(sum) / float64(count)
		avgRating = int((avg/5.0)*100.0 + 0.5)
	} else {
		avgRating = 0
	}

	db.Model(&models.VPN{}).Where("id_vpn = ?", vpnID).
		Updates(map[string]interface{}{
			"vpn_rating": avgRating,
		})

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]interface{}{
		"vpn_rating": avgRating,
		"count":      count,
	})
}
