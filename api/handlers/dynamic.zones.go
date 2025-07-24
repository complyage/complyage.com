package handlers

import (
	"base/models"
	"fmt"
	"net/http"
	"sync"
	"time"

	// Updated import path to base/models
	"base/db"        // Import the db package to access db.DB
	"base/responses" // Import the responses package
)

// ZoneOutput defines the structure for the JSON response
// with simplified field names as requested.
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

var (
	cachedZones     []ZoneOutput // Cache now stores ZoneOutput objects
	zoneCacheExpiry time.Time
	zoneCacheMutex  sync.Mutex
)

// fetchZonesFromDB fetches all zones from the database using the global db.DB instance.
// This function still returns models.Zone, as it's directly from the DB.
func fetchZonesFromDB() ([]models.Zone, error) {
	var zones []models.Zone
	// Use the global DB instance from the base/dbb package
	result := db.DB.Find(&zones)
	if result.Error != nil {
		return nil, fmt.Errorf("failed to fetch zones from database: %w", result.Error)
	}
	return zones, nil
}

// ZoneHandler handles requests for zone data, with caching.
// It is now a direct http.HandlerFunc, suitable for mux.HandleFunc.
func ZoneHandler(w http.ResponseWriter, r *http.Request) {
	zoneCacheMutex.Lock()
	defer zoneCacheMutex.Unlock()

	// Check if cache is valid and not expired
	if time.Now().Before(zoneCacheExpiry) && cachedZones != nil {
		responses.Success(w, http.StatusOK, cachedZones)
		return
	}

	// If cache is expired or empty, fetch new data from DB
	dbZones, err := fetchZonesFromDB() // Renamed local variable to avoid conflict
	if err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to fetch zones")
		return
	}

	// Transform models.Zone to ZoneOutput for the desired JSON structure
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

	// Update cache with fresh data (now of type []ZoneOutput)
	cachedZones = outputZones
	zoneCacheExpiry = time.Now().Add(15 * time.Minute)

	// Return the zones using the standardized success response
	responses.Success(w, http.StatusOK, outputZones)
}
