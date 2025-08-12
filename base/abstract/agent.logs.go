package abstract

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/models"
)

// ||------------------------------------------------------------------------------------------------||
// || Get User Verifications
// ||------------------------------------------------------------------------------------------------||

func CreateAgentLog(siteId int64, level int, authId string) error {
	log := models.AgentRequestLog{
		FidSite:      siteId,
		RequestLevel: level,
		RequestID:    authId,
	}

	if err := db.DB.Create(&log).Error; err != nil {
		return err
	}
	return nil
}
