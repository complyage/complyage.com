package verify

import (
	"github.com/complyage/base/db/models"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Insert
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseInsert() error {
	app.Log.Info("Inserting Into Database - ", v.UUID)
	//||------------------------------------------------------------------------------------------------||
	//|| Create the database model
	//||------------------------------------------------------------------------------------------------||
	model := models.Verify{
		UUID:       v.UUID,
		Type:       v.Type.String(),
		Display:    v.Display,
		FidAccount: v.Account.ID,
		Status:     v.Status.String(),
		UpdatedAt:  v.UpdatedAt,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	err := Database.Save(&model).Where("verify_uuid = ?", v.UUID).Error
	if err != nil {
		app.Log.Error("Failed to insert verification:", err)
		return app.Err("Verify").Error("DATABASE_INSERT_FAILED")
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Update Database
//||------------------------------------------------------------------------------------------------||

func (v *Verification) DatabaseUpdate() error {
	app.Log.Info("Updating Database - ", v.UUID)
	//||------------------------------------------------------------------------------------------------||
	//|| Create the database model
	//||------------------------------------------------------------------------------------------------||
	model := models.Verify{
		UUID:       v.UUID,
		Type:       v.Type.String(),
		Display:    v.Display,
		FidAccount: v.Account.ID,
		Status:     v.Status.String(),
		UpdatedAt:  v.UpdatedAt,
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||
	err := Database.Model(&models.Verify{}).Where("verify_uuid = ?", v.UUID).Updates(model).Error
	if err != nil {
		app.Log.Error("Failed to update verification:", err)
		return app.Err("Verify").Error("DATABASE_UPDATE_FAILED")
	}
	return nil
}
