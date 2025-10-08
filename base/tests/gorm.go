package tests

import (
	"testing"

	"github.com/ralphferrara/aria/db"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ||------------------------------------------------------------------------------------------------||
// || CreateFakeGormWrapper: returns an in-memory GormWrapper for testing
// ||------------------------------------------------------------------------------------------------||

func CreateTestGormWrapper(t *testing.T, models ...interface{}) *db.GormWrapper {
	t.Helper()

	// modernc.org/sqlite works with CGO_ENABLED=0
	dial := sqlite.Open("file::memory:?cache=shared&_pragma=foreign_keys(ON)")

	dbConn, err := gorm.Open(dial, &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open in-memory sqlite: %v", err)
	}

	// Auto-migrate schema for test models
	if err := dbConn.AutoMigrate(models...); err != nil {
		t.Fatalf("failed to migrate models: %v", err)
	}

	return &db.GormWrapper{
		Name: "test",
		DB:   dbConn,
	}
}
