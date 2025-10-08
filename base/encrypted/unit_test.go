package encrypted

import (
	"testing"

	"github.com/complyage/base/tests"
	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| Address
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadADDR(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.Address{
		Line1:   "123 Main St",
		City:    "Metropolis",
		State:   "CA",
		Postal:  "90210",
		Country: "USA",
	}
	uuid := "test-uuid-1234"

	if err := SaveADDR(pub, uuid, orig); err != nil {
		t.Fatalf("SaveADDR failed: %v", err)
	}

	got, err := LoadADDR(uuid, priv)
	if err != nil {
		t.Fatalf("LoadADDR failed: %v", err)
	}

	if got.Line1 != orig.Line1 || got.City != orig.City || got.Postal != orig.Postal {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| CreditCard
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadCRCD(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.CreditCard{
		LastFour: "4242",
		CardType: "VISA",
		ExpYear:  "2030",
		Address: types.Address{
			Line1:   "456 Market St",
			City:    "Gotham",
			State:   "NY",
			Postal:  "10001",
			Country: "USA",
		},
		TransactionId: "txn-001",
	}
	uuid := "test-uuid-crcd"

	if err := SaveCRCD(pub, uuid, orig); err != nil {
		t.Fatalf("SaveCRCD failed: %v", err)
	}

	got, err := LoadCRCD(uuid, priv)
	if err != nil {
		t.Fatalf("LoadCRCD failed: %v", err)
	}

	if got.LastFour != orig.LastFour || got.CardType != orig.CardType || got.ExpYear != orig.ExpYear {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Facial
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadFACE(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.Facial{
		Age: 30,
		Min: 25,
		Max: 35,
		DOB: types.DOB{Year: 1995, Month: 5, Day: 20},
		Selfie: types.Media{
			Exists: true,
			Size:   12345,
			Mime:   "image/jpeg",
			Base64: "base64-selfie-data",
		},
		DOBMatch: true,
	}
	uuid := "test-uuid-face"

	if err := SaveFACE(pub, uuid, orig); err != nil {
		t.Fatalf("SaveFACE failed: %v", err)
	}

	got, err := LoadFACE(uuid, priv)
	if err != nil {
		t.Fatalf("LoadFACE failed: %v", err)
	}

	if got.Age != orig.Age || got.DOB.Year != orig.DOB.Year || got.Selfie.Mime != orig.Selfie.Mime {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Identification
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadIDEN(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.Identification{
		IDType: "Passport",
		Number: "A1234567",
		Front:  types.Media{Exists: true, Mime: "image/jpeg", Size: 1000, Base64: "front-image"},
		Back:   types.Media{Exists: true, Mime: "image/jpeg", Size: 2000, Base64: "back-image"},
		Selfie: types.Media{Exists: true, Mime: "image/jpeg", Size: 3000, Base64: "selfie-image"},
		Address: types.Address{
			Line1:   "789 Broadway",
			City:    "Star City",
			State:   "WA",
			Postal:  "98101",
			Country: "USA",
		},
		DOB:  &types.DOB{Year: 1988, Month: 12, Day: 5},
		Name: &types.Name{First: "Oliver", Last: "Queen"},
	}
	uuid := "test-uuid-iden"

	if err := SaveIDEN(pub, uuid, orig); err != nil {
		t.Fatalf("SaveIDEN failed: %v", err)
	}

	got, err := LoadIDEN(uuid, priv)
	if err != nil {
		t.Fatalf("LoadIDEN failed: %v", err)
	}

	if got.IDType != orig.IDType || got.Number != orig.Number || got.Address.City != orig.Address.City {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| EmailAddress
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadMAIL(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.EmailAddress{
		Email: "user@example.com",
	}
	uuid := "test-uuid-mail"

	if err := SaveMAIL(pub, uuid, orig); err != nil {
		t.Fatalf("SaveMAIL failed: %v", err)
	}

	got, err := LoadMAIL(uuid, priv)
	if err != nil {
		t.Fatalf("LoadMAIL failed: %v", err)
	}

	if got.Email != orig.Email {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}

//||------------------------------------------------------------------------------------------------||
//|| PhoneNumber
//||------------------------------------------------------------------------------------------------||

func TestSaveAndLoadPHNE(t *testing.T) {
	store := tests.CreateStorage(t)
	Init(store)
	priv, pub, _ := tests.GenerateKeyPair(t)

	orig := types.PhoneNumber{
		CountryCode: "1",
		Number:      "5551234567",
	}
	uuid := "test-uuid-phne"

	if err := SavePHNE(pub, uuid, orig); err != nil {
		t.Fatalf("SavePHNE failed: %v", err)
	}

	got, err := LoadPHNE(uuid, priv)
	if err != nil {
		t.Fatalf("LoadPHNE failed: %v", err)
	}

	if got.CountryCode != orig.CountryCode || got.Number != orig.Number {
		t.Errorf("decrypted value mismatch: got %+v, want %+v", got, orig)
	}
}
