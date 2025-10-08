package verify

import (
	"encoding/json"

	"github.com/complyage/base/types"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Finalize and persist a Verification to storage as <uuid>.json
//||------------------------------------------------------------------------------------------------||

func (v *Verification) Finalize(confirmUpdates bool) error {
	//||------------------------------------------------------------------------------------------------||
	//|| Validate
	//||------------------------------------------------------------------------------------------------||
	if v == nil {
		return app.Err("Verify").Error("NIL_VERIFICATION")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| StopGap to ensure that Identity and Transaction have been updated before finalize
	//||------------------------------------------------------------------------------------------------||
	if confirmUpdates {
		if v.IdentityUpdated == false {
			panic("Identity not updated before finalize")
		}
		if v.EncryptedSaved == false {
			panic("Encrypted data not saved before finalize")
		}
		if (v.Type == types.DataTypeCRCD || v.Type == types.DataTypeADDR) && !v.TransactionSaved {
			panic("Transaction not saved before finalize")
		}
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Touch updated timestamp
	//||------------------------------------------------------------------------------------------------||
	v.Updated()
	//||------------------------------------------------------------------------------------------------||
	//|| Final
	//||------------------------------------------------------------------------------------------------||
	final := v.ConvertToFinalVerification()
	//||------------------------------------------------------------------------------------------------||
	//|| Serialize
	//||------------------------------------------------------------------------------------------------||
	data, mErr := json.MarshalIndent(final, "", "  ")
	if mErr != nil {
		app.Log.Error("Verification Final Marshal Error:", mErr)
		return app.Err("Verify").Error("VERIFY_SAVE_MARSHAL")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Store
	//||------------------------------------------------------------------------------------------------||
	err := Storage.Put(v.Filename(), data)
	if err != nil {
		return app.Err("Verify").Error("VERIFY_SAVE_PUT")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Saved
	//||------------------------------------------------------------------------------------------------||
	app.Log.Data("Saved Final Verification:", v.Filename())
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	return nil

}

//||------------------------------------------------------------------------------------------------||
//|| Convert to Final Verification Struct
//||------------------------------------------------------------------------------------------------||

func (v *Verification) ConvertToFinalVerification() VerificationFinal {
	return VerificationFinal{
		UUID:      v.UUID,
		Status:    v.Status,
		Type:      v.Type,
		Display:   v.Display,
		Level:     v.Level,
		Step:      v.Step,
		Steps:     v.Steps,
		Moderate:  v.Moderate,
		Timestamp: v.Timestamp,
		Completed: v.Completed,
	}
}
