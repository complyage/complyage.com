package verify

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/db"
	"github.com/ralphferrara/aria/storage"
)

//||------------------------------------------------------------------------------------------------||
//|| Load: load a Verification from storage by UUID
//||------------------------------------------------------------------------------------------------||

func AgentLoad(d *db.GormWrapper, s *storage.Storage, uuid string, vType DataType) (Verification, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Validate
	//||------------------------------------------------------------------------------------------------||
	if uuid == "" {
		return Verification{}, app.Err("verify").Error("UUID_EMPTY")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| V
	//||------------------------------------------------------------------------------------------------||
	v := Verification{}
	v.UUID = uuid
	//||------------------------------------------------------------------------------------------------||
	//|| Fetch Verification JSON
	//||------------------------------------------------------------------------------------------------||
	data, err := s.Get(v.ObjectName(false))
	if err != nil {
		fmt.Println("Error fetching verification:", v.ObjectName(false), err)
		vErr := app.Err("verify").Get("RECORD_NOT_FOUND")
		verifyRecord := panicReject(uuid, vType, vErr)
		return verifyRecord, errors.New(vErr.Message)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal
	//||------------------------------------------------------------------------------------------------||
	if err := json.Unmarshal(data, &v); err != nil {
		fmt.Println("Error unmarshaling verification:", v.ObjectName(false), err)
		vErr := app.Err("verify").Get("RECORD_CORRUPT")
		verifyRecord := panicReject(uuid, vType, vErr)
		return verifyRecord, errors.New(vErr.Message)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Storage
	//||------------------------------------------------------------------------------------------------||
	v.Database = d
	v.Storage = *s
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	return v, nil
}
