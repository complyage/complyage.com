package keeper

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"time"

	"github.com/complyage/base/enforce"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/random"
	ariaHTTP "github.com/ralphferrara/aria/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Initiate a Keeper Session
//||------------------------------------------------------------------------------------------------||

func Initiate(r *http.Request) (KeeperRecord, error) {
	app.Log.Info("Initiating a new record...")
	start := time.Now()
	//||------------------------------------------------------------------------------------------------||
	//|| Client ID
	//||------------------------------------------------------------------------------------------------||
	clientId := r.URL.Query().Get("client_id")
	returnURL := r.URL.Query().Get("return_url")
	hostName := ariaHTTP.GetOriginDomain(r)
	//||------------------------------------------------------------------------------------------------||
	//|| Decode ReturnURL
	//||------------------------------------------------------------------------------------------------||
	if returnURL != "" {
		decoded, err := base64.StdEncoding.DecodeString(returnURL)
		if err != nil {
			returnURL = hostName
		} else {
			returnURL = string(decoded)
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Load Site
	//||------------------------------------------------------------------------------------------------||
	site, sErr := enforce.LoadSite(clientId, hostName)
	if sErr != nil {
		return KeeperRecord{}, sErr
	}
	fmt.Printf("⏱ Site took %dms\n", time.Since(start).Milliseconds())
	//||------------------------------------------------------------------------------------------------||
	//|| Load Zone
	//||------------------------------------------------------------------------------------------------||
	zone := enforce.LoadZone(r, site)
	fmt.Printf("⏱ Zone took %dms\n", time.Since(start).Milliseconds())
	//||------------------------------------------------------------------------------------------------||
	//|| Generate Keeper
	//||------------------------------------------------------------------------------------------------||
	fmt.Printf("⏱ Random took %dms\n", time.Since(start).Milliseconds())
	keep := KeeperRecord{
		KeeperId:  random.RandomString(36),
		Enforced:  zone.Enforced,
		Verified:  false,
		Age:       0,
		UserId:    0,
		IPAddress: zone.IPAddress,
		ClientId:  clientId,
		Status:    "INIT",
		ReturnURL: returnURL,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Save
	//||------------------------------------------------------------------------------------------------||
	keep.Save()
	fmt.Printf("⏱ Redis took %dms\n", time.Since(start).Milliseconds())
	return keep, nil
}
