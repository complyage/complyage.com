package vpns

import (
	"math/rand"
	"time"

	"github.com/complyage/base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| GetRandomVPN – return a random VPN from memory
//||------------------------------------------------------------------------------------------------||

func GetRandomVPN() *models.VPN {
	if len(VPNs) == 0 {
		return nil
	}
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(VPNs))
	return &VPNs[index]
}
