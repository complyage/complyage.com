package vpns

import (
	"fmt"
	"log"
	"time"

	"github.com/complyage/base/db/abstract"
	"github.com/complyage/base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| In-memory slice of VPNs
//||------------------------------------------------------------------------------------------------||

var VPNs []models.VPN

//||------------------------------------------------------------------------------------------------||
//|| Package init – automatically start refresher on import
//||------------------------------------------------------------------------------------------------||

func Init() {
	LoadVPNs()
	StartVPNRefresher(30) // refresh every 30 minutes
}

//||------------------------------------------------------------------------------------------------||
//|| LoadVPNs – load all VPNs into memory
//||------------------------------------------------------------------------------------------------||

func LoadVPNs() error {
	var results []models.VPN
	results, err := abstract.ReturnAllVPNs()
	if err != nil {
		return err
	}

	VPNs = make([]models.VPN, len(results))
	copy(VPNs, results)

	fmt.Printf("\n\033[32m[LOAD] - Loaded %d VPNs into memory\033[0m\n", len(VPNs))
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| StartVPNRefresher – refresh VPNs in memory every X minutes
//||------------------------------------------------------------------------------------------------||

func StartVPNRefresher(minutes int) {
	ticker := time.NewTicker(time.Duration(minutes) * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				if err := LoadVPNs(); err != nil {
					log.Printf("[VPN] Failed to reload VPNs: %v\n", err)
				}
			}
		}
	}()
}
