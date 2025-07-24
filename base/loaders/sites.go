package loaders

import (
	"fmt"
	"sync"
	"time"

	"base/db"
	"base/models"
)

//||------------------------------------------------------------------------------------------------||
//|| In‑memory cache of Site records
//||------------------------------------------------------------------------------------------------||

var (
	Sites      []models.Site
	sitesMutex sync.RWMutex
)

// refresh interval
const siteRefreshInterval = 5 * time.Minute

//||------------------------------------------------------------------------------------------------||
//|| loadSitesFromDB
//||
//|| Internal: queries the sites table and updates the in‑memory cache.
//||------------------------------------------------------------------------------------------------||

func loadSitesFromDB() error {
	var results []models.Site
	if err := db.DB.
		Table("sites").
		Where("site_status NOT IN ?", []string{"RMVD", "BNND"}).
		Find(&results).Error; err != nil {
		return fmt.Errorf("failed to load sites: %w", err)
	}

	sitesMutex.Lock()
	Sites = results
	sitesMutex.Unlock()

	fmt.Printf("Loaded %d sites into memory\n", len(results))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| StartSiteLoader
//||
//|| Call this once at application startup. It will immediately load the
//|| sites cache, then refresh every 5 minutes in the background.
//||------------------------------------------------------------------------------------------------||

func StartSiteLoader() {
	// initial load
	if err := loadSitesFromDB(); err != nil {
		fmt.Println("Site loader initial load error:", err)
	}

	// start periodic refresh
	ticker := time.NewTicker(siteRefreshInterval)
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
//||
//|| Returns a snapshot of the current cached sites.
//||------------------------------------------------------------------------------------------------||

func GetSites() []models.Site {
	sitesMutex.RLock()
	defer sitesMutex.RUnlock()

	// return a copy to avoid data races
	copySlice := make([]models.Site, len(Sites))
	copy(copySlice, Sites)
	return copySlice
}

// ||------------------------------------------------------------------------------------------------||
// || GetSiteByPublic
// ||
// || Looks up a Site by its `site_public` key in the in‑memory cache.
// || Returns the site and true if found, or nil and false otherwise.
// ||------------------------------------------------------------------------------------------------||

func GetSiteByPublic(publicKey string) *models.Site {
	sitesMutex.RLock()
	defer sitesMutex.RUnlock()

	for i := range Sites {
		if Sites[i].SitePublic == publicKey {
			// return pointer to the cached element
			return &Sites[i]
		}
	}
	return nil
}
