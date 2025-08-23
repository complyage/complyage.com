package service

import (
	"base/helpers"
	"base/storage"
	"log"
	"strings"

	"github.com/joho/godotenv"
)

//||------------------------------------------------------------------------------------------------||
//|| Global Storage Instance
//||------------------------------------------------------------------------------------------------||

var Storage *storage.Storage

//||------------------------------------------------------------------------------------------------||
//|| Init: Load Storage Backend From .env
//||------------------------------------------------------------------------------------------------||

func Init() {
	// Load environment variables (if any)
	_ = godotenv.Load("../.env")

	backend := strings.ToUpper(helpers.GetEnv("STORAGE_BACKEND", "LOCAL"))
	var cfg storage.StoreConfig

	switch backend {
	case "S3":
		cfg = storage.StoreConfig{
			Backend:   storage.BackendS3,
			Bucket:    helpers.GetEnv("STORAGE_BUCKET", ""),
			Region:    helpers.GetEnv("STORAGE_REGION", ""),
			AccessKey: helpers.GetEnv("STORAGE_ACCESS_KEY", ""),
			SecretKey: helpers.GetEnv("STORAGE_SECRET_KEY", ""),
		}
	case "MINIO":
		cfg = storage.StoreConfig{
			Backend:   storage.BackendMinIO,
			Bucket:    helpers.GetEnv("STORAGE_BUCKET", ""),
			Region:    helpers.GetEnv("STORAGE_REGION", ""),
			Endpoint:  helpers.GetEnv("STORAGE_ENDPOINT", ""),
			AccessKey: helpers.GetEnv("STORAGE_ACCESS_KEY", ""),
			SecretKey: helpers.GetEnv("STORAGE_SECRET_KEY", ""),
			UseSSL:    helpers.GetEnvBool("STORAGE_USE_SSL", false),
		}
	case "AZURE":
		cfg = storage.StoreConfig{
			Backend:     storage.BackendAzure,
			Bucket:      helpers.GetEnv("STORAGE_BUCKET", ""),
			AccountName: helpers.GetEnv("STORAGE_ACCOUNT_NAME", ""),
			AccountKey:  helpers.GetEnv("STORAGE_ACCOUNT_KEY", ""),
		}
	case "GCP":
		cfg = storage.StoreConfig{
			Backend:         storage.BackendGCP,
			Bucket:          helpers.GetEnv("STORAGE_BUCKET", ""),
			CredentialsJSON: helpers.GetEnv("STORAGE_CREDENTIALS_JSON", ""),
		}
	case "LOCAL":
		cfg = storage.StoreConfig{
			Backend:   storage.BackendLocal,
			Bucket:    "dummy",
			LocalPath: "./localdata",
		}
	default:
		log.Fatalf("[storage] Invalid STORAGE_BACKEND: %s", backend)
	}

	if missingConfig(cfg) {
		log.Fatalf("[storage] Missing config for backend %s", backend)
	}

	st := &storage.Storage{Config: cfg}
	if err := st.Init(); err != nil {
		log.Fatalf("[storage] Init failed: %v", err)
	}
	Storage = st
	log.Printf("[storage] Loaded backend: %s", backend)
}

//||------------------------------------------------------------------------------------------------||
//|| Helper: Check if Any Essential Config Is Missing
//||------------------------------------------------------------------------------------------------||

func missingConfig(cfg storage.StoreConfig) bool {
	switch cfg.Backend {
	case storage.BackendS3:
		return cfg.Bucket == "" || cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.Region == ""
	case storage.BackendMinIO:
		return cfg.Bucket == "" || cfg.AccessKey == "" || cfg.SecretKey == "" || cfg.Endpoint == ""
	case storage.BackendAzure:
		return cfg.Bucket == "" || cfg.AccountName == "" || cfg.AccountKey == ""
	case storage.BackendGCP:
		return cfg.Bucket == "" || cfg.CredentialsJSON == ""
	}
	return false
}
