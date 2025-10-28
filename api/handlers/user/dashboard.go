package user

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"fmt"
	"net/http"

	"github.com/complyage/base/identity"
	"github.com/complyage/base/ips"
	"github.com/complyage/base/types"
	"github.com/complyage/base/zones"

	ariaHTTP "github.com/ralphferrara/aria/http"
	"github.com/ralphferrara/aria/responses"

	"github.com/ralphferrara/aria/auth/actions"
	"github.com/ralphferrara/aria/base/str"
)

//||------------------------------------------------------------------------------------------------||
//|| Struct for Response
//||------------------------------------------------------------------------------------------------||

type UserDashboardData struct {
	IsVerified  bool              `json:"isVerified"`
	VerifiedAge int               `json:"verifiedAge"`
	IPAddress   string            `json:"ipAddress"`
	MinimumType types.DataType    `json:"minimumType"`
	Location    ips.Location      `json:"location"`
	Zone        *zones.ShortZone  `json:"zone"`
	Identity    identity.Identity `json:"identity"`
}

//||------------------------------------------------------------------------------------------------||
//|| UserVerifications - Returns latest status of each verification type for logged-in user
//||------------------------------------------------------------------------------------------------||

func UserDashboard(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	var response UserDashboardData

	//||------------------------------------------------------------------------------------------------||
	//|| Check Account
	//||------------------------------------------------------------------------------------------------||

	_, account, _, err := actions.LoadSessionAccount(r)
	if err != nil {
		responses.Error(w, http.StatusUnauthorized, "Invalid session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get IP
	//||------------------------------------------------------------------------------------------------||

	ipAddress := ariaHTTP.GetClientIP(r)
	fmt.Println("Client IP Address:", ipAddress)
	if ipAddress == "" {
		responses.Error(w, http.StatusBadRequest, "Unable to determine IP address")
		return
	}
	response.IPAddress = ipAddress

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch IP Location
	//||------------------------------------------------------------------------------------------------||

	location, err := ips.GetLocationByIP(ipAddress)
	if err != nil {
		fmt.Println("Error fetching location:", err)
		responses.Error(w, http.StatusBadRequest, "Unable to determine location from IP address")
		return
	}
	response.Location = location

	//||------------------------------------------------------------------------------------------------||
	//|| Get Zone
	//||------------------------------------------------------------------------------------------------||

	zone, ok := zones.FetchShortZoneByLocation(location.Region, location.Country)
	if !ok {
		responses.Error(w, http.StatusBadRequest, "Unable to determine legal zone from location")
		return
	}
	response.Zone = zone

	//||------------------------------------------------------------------------------------------------||
	//|| Get Identity
	//||------------------------------------------------------------------------------------------------||

	idn, err := identity.Load(account.ID)
	if err != nil {
		idn = identity.Identity{} // now correct
	}
	response.Identity = idn

	//||------------------------------------------------------------------------------------------------||
	//|| Figure out if we are and can be verified by CRCD
	//||------------------------------------------------------------------------------------------------||

	response.IsVerified = false
	response.VerifiedAge = 0
	response.MinimumType = types.DataTypeIDEN

	//||------------------------------------------------------------------------------------------------||
	//|| Figure out if we are and can be verified by CRCD
	//||------------------------------------------------------------------------------------------------||

	var requirements []string
	for _, r := range zone.Requirements {
		requirements = append(requirements, r.String())
	}

	minimumSet := false

	if str.Contains(requirements, "CRCD") {
		fmt.Println("Zone requires CRCD")
		if !minimumSet {
			response.MinimumType = types.DataTypeCRCD
			minimumSet = true
		}
		fmt.Println("CRCD Verified?", idn.CreditCard.Verified, "Age:", idn.CreditCard.DOB.Age(), "MinAge:", zone.MinAge)
		if idn.CreditCard.Verified && idn.CreditCard.DOB.Age() >= zone.MinAge {
			fmt.Println("CRCD meets requirements")
			response.IsVerified = true
			response.VerifiedAge = idn.CreditCard.DOB.Age()
		}
	}

	if str.Contains(requirements, "FACE") {
		if !minimumSet {
			response.MinimumType = types.DataTypeFACE
			minimumSet = true
		}
		if idn.Face.Verified && idn.Face.DOB.Age() >= zone.MinAge {
			response.IsVerified = true
			response.VerifiedAge = idn.Face.DOB.Age()
		}
	}

	if str.Contains(requirements, "IDEN") {
		if !minimumSet {
			response.MinimumType = types.DataTypeIDEN
			minimumSet = true
		}
		if idn.IDCard.Verified && idn.IDCard.DOB.Age() >= zone.MinAge {
			response.IsVerified = true
			response.VerifiedAge = idn.IDCard.DOB.Age()
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return JSON Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, response)
}
