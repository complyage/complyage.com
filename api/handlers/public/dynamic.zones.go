package public

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||
import (
	"base/db/models"
	"fmt"
	"net/http"
	"sync"
	"time"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/responses"
)

// ||------------------------------------------------------------------------------------------------||
// || Interface: Zone Output (Public)
// ||------------------------------------------------------------------------------------------------||
type ZoneOutput struct {
	ID             uint       `json:"id"`
	State          *string    `json:"state"`
	Country        *string    `json:"country"`
	Law            *string    `json:"law"`
	LawDescription *string    `json:"description"`
	Requirements   *string    `json:"requirements"`
	Penalties      *string    `json:"penalties"`
	Effective      *time.Time `json:"effective"`
	Meta           *string    `json:"meta"`
	Latitude       *string    `json:"lat"`
	Longitude      *string    `json:"long"`
}

// ||------------------------------------------------------------------------------------------------||
// || Cache
// ||------------------------------------------------------------------------------------------------||
var (
	cachedZones     []ZoneOutput
	zoneCacheExpiry time.Time
	zoneCacheMutex  sync.Mutex
)

// ||------------------------------------------------------------------------------------------------||
// || Fetch: Zones From DB
// ||------------------------------------------------------------------------------------------------||
func fetchZonesFromDB() ([]models.Zone, error) {
	var zones []models.Zone

	result := app.SQLDB["main"].DB.Where("id_zone <> ?", 9999).Find(&zones)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch zones from database: %w", result.Error)
	}
	return zones, nil
}

// ||------------------------------------------------------------------------------------------------||
// || Handler: Zones (Cached)
// ||------------------------------------------------------------------------------------------------||
func ZoneHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Create Mutex Lock
	//||------------------------------------------------------------------------------------------------||

	zoneCacheMutex.Lock()
	defer zoneCacheMutex.Unlock()

	//||------------------------------------------------------------------------------------------------||
	//|| Cache: Serve if Fresh
	//||------------------------------------------------------------------------------------------------||

	if time.Now().Before(zoneCacheExpiry) && cachedZones != nil {
		responses.Success(w, http.StatusOK, cachedZones)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fetch: From DB
	//||------------------------------------------------------------------------------------------------||

	dbZones, err := fetchZonesFromDB()
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch zones")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Transform: models.Zone -> ZoneOutput
	//||------------------------------------------------------------------------------------------------||

	outputZones := make([]ZoneOutput, len(dbZones))
	for i, z := range dbZones {
		outputZones[i] = ZoneOutput{
			ID:             z.IDZone,
			State:          z.ZoneState,
			Country:        z.ZoneCountry,
			Law:            z.ZoneLaw,
			LawDescription: z.ZoneLawDescription,
			Requirements:   z.ZoneRequirements,
			Penalties:      z.ZonePenalties,
			Effective:      z.ZoneEffective,
			Meta:           z.ZoneMeta,
			Latitude:       z.ZoneLatitude,
			Longitude:      z.ZoneLongitude,
		}
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Cache: Update
	//||------------------------------------------------------------------------------------------------||
	cachedZones = outputZones
	zoneCacheExpiry = time.Now().Add(15 * time.Minute)

	//||------------------------------------------------------------------------------------------------||
	//|| Respond
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, outputZones)
}
