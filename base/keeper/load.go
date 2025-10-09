package keeper

import (
	"encoding/base64"
	"encoding/json"
	"net/http"

	"github.com/ralphferrara/aria/app"
	ariaHTTP "github.com/ralphferrara/aria/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Load
//||------------------------------------------------------------------------------------------------||

func Load(r *http.Request) (KeeperRecord, error) {
	app.Log.Info("Loading Keeper Record...")
	//||------------------------------------------------------------------------------------------------||
	//|| Get Variables
	//||------------------------------------------------------------------------------------------------||
	var record KeeperRecord
	ipAddress := ariaHTTP.GetClientIP(r)
	cookie, err := r.Cookie(GATE_COOKIE_NAME)
	if err != nil || cookie.Value == "" {
		record, err := Initiate(r)
		if err != nil {
			return KeeperRecord{}, err
		}
		return record, nil
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Return from Redis
	//||------------------------------------------------------------------------------------------------||
	app.Log.Info("Fetching session from record...")
	recordJSON, re := SessionCache.Get(REDIS_PREFIX + cookie.Value)
	if re != nil {
		record, err := Initiate(r)
		if err != nil {
			return KeeperRecord{}, err
		}
		return record, nil
	}
	json.Unmarshal([]byte(recordJSON), &record)
	//||------------------------------------------------------------------------------------------------||
	//|| Verify
	//||------------------------------------------------------------------------------------------------||
	app.Log.Info("Verifying keeper record...")
	app.Log.Info(record.IPAddress, ipAddress)
	if record.IPAddress != ipAddress {
		return KeeperRecord{}, app.Err("Gate").Error("RECORD_MISMATTCH")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Update ReturnURL
	//||------------------------------------------------------------------------------------------------||
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
	if returnURL != "" && record.ReturnURL != returnURL {
		record.ReturnURL = returnURL
		record.Save()
	}
	app.Log.Info("Successfully loaded keeper record...")
	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||
	return record, nil
}
