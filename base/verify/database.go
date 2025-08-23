package verify

import (
	"base/models"
	"encoding/json"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Update Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseUpdate() error {
	logInfo("DATABASE :: UPDATE DATABASE")
	err := v.DatabaseSaveVerify()
	if err != nil {
		fmt.Println("Failed to update verification:", err)
		return err
	}
	err = v.DatabaseSaveIdentity()
	if err != nil {
		fmt.Println("Failed to update identity:", err)
		return err
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseLoadIdentity() error {
	logInfo("DATABASE :: LOAD IDENTITY")
	//||------------------------------------------------------------------------------------------------||
	//|| Pull from DB
	//||------------------------------------------------------------------------------------------------||
	var identityJSON string
	result := v.Database.DB.Raw("SELECT account_identity FROM accounts WHERE id_account = ?", v.FidAccount).Scan(&identityJSON)
	if result.Error != nil {
		logInfo("Verify: Failed to load identity from database, resetting to empty")
		return result.Error
	}
	if identityJSON == "" {
		logInfo("Verify: No identity found, resetting to empty")
		v.Identity = Identity{}
		return nil
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal the JSON string into v.Identity
	//||------------------------------------------------------------------------------------------------||
	err := json.Unmarshal([]byte(identityJSON), &v.Identity)
	if err != nil {
		logInfo("Verify: Identity is invalid JSON, resetting to empty")
		v.Identity = Identity{}
		return err
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseSaveIdentity() error {
	logInfo("DATABASE :: SAVE IDENTITY")
	bytes, err := json.Marshal(v.Identity)
	if err != nil {
		logInfo("Failed to marshal identity:", err)
		return err
	}
	return v.Database.DB.Exec(
		"UPDATE accounts SET account_identity=? WHERE id_account=?",
		string(bytes), v.FidAccount,
	).Error
}

//||------------------------------------------------------------------------------------------------||
//|| Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseSaveVerify() error {
	logInfo("DATABASE :: SAVE VERIFY :: COMPLETE")
	//||------------------------------------------------------------------------------------------------||
	//|| Create the database model
	//||------------------------------------------------------------------------------------------------||
	model := models.Verify{
		UUID:       v.UUID,
		Type:       v.Type.String(),
		Display:    v.Display,
		FidAccount: v.FidAccount,
		Status:     v.Status.String(),
		UpdatedAt:  v.UpdatedAt,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	return v.Database.DB.Save(&model).Where("verify_uuid = ?", v.UUID).Error
}
