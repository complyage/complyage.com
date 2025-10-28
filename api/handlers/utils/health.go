package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| HealthResponse Structure
//||------------------------------------------------------------------------------------------------||

type HealthResponse struct {
	Status     string                 `json:"status"`
	Time       string                 `json:"time"`
	Components map[string]interface{} `json:"components"`
}

//||------------------------------------------------------------------------------------------------||
//|| Health Handler
//||------------------------------------------------------------------------------------------------||

func HealthHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------||
	//|| Build base response
	//||------------------------------------------------------------------------------------------||
	health := HealthResponse{
		Status:     "ok",
		Time:       time.Now().Format(time.RFC3339),
		Components: make(map[string]interface{}),
	}

	//||------------------------------------------------------------------------------------------||
	//|| 1. Database check
	//||------------------------------------------------------------------------------------------||
	db := app.SQLDB["main"].DB
	if db == nil {
		health.Components["database"] = "not initialized"
		health.Status = "error"
	} else if err := db.Exec("SELECT 1").Error; err != nil {
		health.Components["database"] = fmt.Sprintf("error: %v", err)
		health.Status = "error"
	} else {
		health.Components["database"] = "ok"
	}

	//||------------------------------------------------------------------------------------------||
	//|| 2. Redis checks
	//||------------------------------------------------------------------------------------------||
	redisStatus := map[string]string{}
	for name, cache := range app.CacheRedis {
		if cache == nil {
			redisStatus[name] = "not initialized"
			health.Status = "error"
			continue
		}
		if err := cache.Ping(); err != nil {
			redisStatus[name] = fmt.Sprintf("error: %v", err)
			health.Status = "error"
		} else {
			redisStatus[name] = "ok"
		}
	}
	health.Components["redis"] = redisStatus

	//||------------------------------------------------------------------------------------------||
	//|| 3. Storage (MinIO/S3) checks
	//||------------------------------------------------------------------------------------------||
	storageStatus := map[string]string{}
	for name, store := range app.Storages {
		if store == nil {
			storageStatus[name] = "not initialized"
			health.Status = "error"
			continue
		}
		if err := store.Ping(); err != nil {
			storageStatus[name] = fmt.Sprintf("error: %v", err)
			health.Status = "error"
		} else {
			storageStatus[name] = "ok"
		}
	}
	health.Components["storages"] = storageStatus

	//||------------------------------------------------------------------------------------------||
	//|| Return JSON response
	//||------------------------------------------------------------------------------------------||
	w.Header().Set("Content-Type", "application/json")
	if health.Status == "ok" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusServiceUnavailable)
	}
	json.NewEncoder(w).Encode(health)
}
