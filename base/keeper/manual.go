package keeper

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/complyage/base/enforce"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/random"
	ariaHTTP "github.com/ralphferrara/aria/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func Manual(r *http.Request, ipAddress, sessionValue, clientId string) (KeeperRecord, error) {
	app.Log.Info("Loading Manual Keeper Record...")
	//||------------------------------------------------------------------------------------------------||
	//|| Return from Redis
	//||------------------------------------------------------------------------------------------------||
	app.Log.Info("Fetching session from record...")
	var record KeeperRecord
	recordJSON, rErr := SessionCache.Get(REDIS_PREFIX + sessionValue)
	if rErr != nil {
		start := time.Now()
		//||------------------------------------------------------------------------------------------------||
		//|| Client ID
		//||------------------------------------------------------------------------------------------------||
		hostName := ariaHTTP.GetOriginDomain(r)
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
			ReturnURL: "",
		}
		//||------------------------------------------------------------------------------------------------||
		//|| Save
		//||------------------------------------------------------------------------------------------------||
		keep.Save()
		record = keep
	} else {
		json.Unmarshal([]byte(recordJSON), &record)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||
	app.Log.Info("Verifying keeper record...")
	app.Log.Info(record.IPAddress, ipAddress)
	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||
	return record, nil
}
