//||------------------------------------------------------------------------------------------------||
//|| Sites Package
//|| sites.go
//||------------------------------------------------------------------------------------------------||

package sites

import (
	"fmt"
	"sync"
	"time"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| In-memory cache of Site records
//||------------------------------------------------------------------------------------------------||

var (
	Sites      []models.Site
	sitesMutex sync.RWMutex
)

//||------------------------------------------------------------------------------------------------||
//|| loadSitesFromDB
//||------------------------------------------------------------------------------------------------||

func loadSitesFromDB() error {
	results, err := abstract.ReturnAllSites()
	if err != nil {
		return err
	}
	sitesMutex.Lock()
	Sites = results
	sitesMutex.Unlock()
	fmt.Printf("\033[32m[LOAD] - Loaded %d sites into memory\033[0m\n", len(results))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| LoadSites – initial load + refresh every 5 minutes
//||------------------------------------------------------------------------------------------------||

func LoadSites() {
	//||------------------------------------------------------------------------------------------------||
	//|| Initial Load
	//||------------------------------------------------------------------------------------------------||
	if err := loadSitesFromDB(); err != nil {
		fmt.Println("Site loader initial load error:", err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Setup Ticker
	//||------------------------------------------------------------------------------------------------||
	ticker := time.NewTicker(5 * time.Minute)
	go func() {
		for range ticker.C {
			if err := loadSitesFromDB(); err != nil {
				fmt.Println("Site loader refresh error:", err)
			}
		}
	}()
}

//||------------------------------------------------------------------------------------------------||
//|| GetSites
//||------------------------------------------------------------------------------------------------||

func GetSites() []models.Site {
	sitesMutex.RLock()
	defer sitesMutex.RUnlock()
	copySlice := make([]models.Site, len(Sites))
	copy(copySlice, Sites)
	return copySlice
}

//||------------------------------------------------------------------------------------------------||
//|| FetchSiteByClientId
//||------------------------------------------------------------------------------------------------||

func FetchSiteByClientId(clientId string) (models.Site, error) {
	if app.Config.App.Env != "production" {
		fmt.Println("Fetching site from database by clientId:", clientId)
		site, err := abstract.GetSiteByClientId(clientId)
		if err != nil {
			return models.Site{}, fmt.Errorf("site not found")
		}
		return site, nil
	} else {
		local, err := GetSiteByPublic(clientId)
		if err != nil {
			return models.Site{}, fmt.Errorf("site not found")
		}
		return local, nil
	}
}

//||------------------------------------------------------------------------------------------------||
//|| GetSiteByPublic
//||------------------------------------------------------------------------------------------------||

func GetSiteByPublic(publicKey string) (models.Site, error) {
	sitesMutex.RLock()
	defer sitesMutex.RUnlock()
	for i := range Sites {
		if Sites[i].Public == publicKey {
			return Sites[i], nil
		}
	}
	return models.Site{}, fmt.Errorf("site not found")
}

//||------------------------------------------------------------------------------------------------||
//|| GetSiteByAgent
//||------------------------------------------------------------------------------------------------||

func GetSiteByAgentKey(agentKey string) *models.Site {
	sitesMutex.RLock()
	defer sitesMutex.RUnlock()
	for i := range Sites {
		if Sites[i].AgentPrivate == agentKey {
			return &Sites[i]
		}
	}
	return nil
}
