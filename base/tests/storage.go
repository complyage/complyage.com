package tests

import (
	"fmt"
	"os"
	"testing"

	"github.com/ralphferrara/aria/storage"
)

//||------------------------------------------------------------------------------------------------||
//|| Helper :: Create a Temporary Storage
//||------------------------------------------------------------------------------------------------||

func CreateStorage(t *testing.T) *storage.Storage {
	t.Helper()
	tmp := t.TempDir()
	cfg := storage.StoreConfig{
		Backend:   storage.StorageLocal,
		LocalPath: tmp,
	}
	st := &storage.Storage{Config: cfg}
	if err := st.InitStorage(); err != nil {
		t.Fatalf("init storage: %v", err)
	}
	return st
}

//||------------------------------------------------------------------------------------------------||
//|| Fake Test but generates the Storage
//||------------------------------------------------------------------------------------------------||

func CreateStorageNoTest() *storage.Storage {
	tmp, _ := os.MkdirTemp("", "storage-test-*")
	cfg := storage.StoreConfig{
		Backend:   storage.StorageLocal,
		LocalPath: tmp,
	}
	st := &storage.Storage{Config: cfg}
	if err := st.InitStorage(); err != nil {
		panic(fmt.Sprintf("init storage: %v", err))
	}
	return st
}
