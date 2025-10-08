package verify_test

import (
	"os"
	"testing"

	"github.com/complyage/base/tests"
	"github.com/complyage/base/types"
	"github.com/complyage/base/verify"
)

// ||------------------------------------------------------------------------------------------------||
// || Helper :: setup test env for verify package
// ||------------------------------------------------------------------------------------------------||

func setupTestEnv(t *testing.T) {
	t.Helper()
	verify.Storage = tests.CreateStorage(t)
	verify.Database = tests.CreateTestGormWrapper(t, &verify.ModelVerification{}).DB
}

// ||------------------------------------------------------------------------------------------------||
// || E2E: Create + Save + Load
// ||------------------------------------------------------------------------------------------------||
func TestVerify_CreateSaveLoad(t *testing.T) {
	setupTestEnv(t)

	// make a fake account
	account := tests.GenerateAccountSessionChecked(t)

	// create verification
	vr := verify.Create(types.DataTypeADDR, account)
	if vr.UUID == "" {
		t.Errorf("expected UUID to be set")
	}
	if vr.Status != verify.StatusPending {
		t.Errorf("expected Pending status, got %v", vr.Status)
	}

	// ensure file storage saved JSON
	_, err := verify.Storage.Get(vr.Filename())
	if err != nil {
		t.Errorf("expected JSON file in storage, got error: %v", err)
	}

	// ensure DB row inserted
	var dbRec verify.ModelVerification
	if err := verify.Database.First(&dbRec, "verification_uuid = ?", vr.UUID).Error; err != nil {
		t.Errorf("expected DB record, got error: %v", err)
	}
}

// ||------------------------------------------------------------------------------------------------||
// || E2E: Status transitions update DB + Storage
// ||------------------------------------------------------------------------------------------------||
func TestVerify_StatusTransitions(t *testing.T) {
	setupTestEnv(t)

	vr := verify.Create(types.DataTypeMAIL, tests.GenerateAccountSessionChecked(t))

	// transition to PendingVerification
	if err := vr.UpdateStatusPendingVerification(); err != nil {
		t.Fatalf("UpdateStatusPendingVerification failed: %v", err)
	}
	if vr.Status != verify.StatusPendingVerification {
		t.Errorf("expected PendingVerification, got %v", vr.Status)
	}

	// transition to Verified
	if err := vr.UpdateStatusVerified("tester"); err != nil {
		t.Fatalf("UpdateStatusVerified failed: %v", err)
	}
	if vr.Status != verify.StatusVerified {
		t.Errorf("expected Verified, got %v", vr.Status)
	}

	// check DB updated
	var dbRec verify.ModelVerification
	if err := verify.Database.First(&dbRec, "verification_uuid = ?", vr.UUID).Error; err != nil {
		t.Fatalf("DB lookup failed: %v", err)
	}
	if dbRec.Status != vr.Status.String() {
		t.Errorf("expected DB status %s, got %s", vr.Status.String(), dbRec.Status)
	}

	// check file updated
	if _, err := os.Stat(vr.Filename()); err != nil {
		t.Errorf("expected JSON file written: %v", err)
	}
}

// ||------------------------------------------------------------------------------------------------||
// || E2E: Finalize writes FinalVerification
// ||------------------------------------------------------------------------------------------------||
func TestVerify_Finalize(t *testing.T) {
	setupTestEnv(t)

	vr := verify.Create(types.DataTypePHNE, tests.GenerateAccountSessionChecked(t))

	vr.IdentityUpdated = true // bypass stopgap
	vr.TransactionSaved = true

	if err := vr.Finalize(true); err != nil {
		t.Fatalf("Finalize failed: %v", err)
	}

	// file should exist
	if _, err := verify.Storage.Get(vr.Filename()); err != nil {
		t.Errorf("expected FinalVerification file, got: %v", err)
	}

	// DB should reflect last status
	var dbRec verify.ModelVerification
	if err := verify.Database.First(&dbRec, "verification_uuid = ?", vr.UUID).Error; err != nil {
		t.Fatalf("DB lookup failed: %v", err)
	}
	if dbRec.Status != vr.Status.String() {
		t.Errorf("expected DB status %s, got %s", vr.Status.String(), dbRec.Status)
	}
}
