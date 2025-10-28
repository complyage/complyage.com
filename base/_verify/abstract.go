package verify

import (
	"fmt"

	"github.com/complyage/base/db/models"
)

//||------------------------------------------------------------------------------------------------||
//|| Update Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseUpdate() error {
	LogInfo("DATABASE :: UPDATE DATABASE")
	err := v.DatabaseSaveVerify()
	if err != nil {
		fmt.Println("Failed to update verification:", err)
		return err
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Insert
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseSaveInsert() error {
	LogInfo("DATABASE :: INSERT VERIFY :: COMPLETE")
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

//||------------------------------------------------------------------------------------------------||
//|| Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseSaveVerify() error {
	LogInfo("DATABASE :: SAVE VERIFY :: COMPLETE")
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
	return v.Database.DB.Model(&models.Verify{}).Where("verify_uuid = ?", v.UUID).Updates(model).Error
}
