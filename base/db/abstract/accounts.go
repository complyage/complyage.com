package abstract

import (
	"base/db/models"
	"base/verify"
	"encoding/json"
	"fmt"

	"github.com/ralphferrara/aria/app"
	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByID(id string) (*models.Account, error) {
	var account models.Account

	result := app.SQLDB["main"].DB.First(&account, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}

	return &account, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on Email
//||------------------------------------------------------------------------------------------------||

func GetAccountByEmail(hashedEmail string) (*models.Account, error) {
	var account models.Account

	result := app.SQLDB["main"].DB.
		Where("account_email = ?", hashedEmail).
		Limit(1).
		Find(&account)

	if result.Error != nil {
		return nil, result.Error
	}
	if result.RowsAffected == 0 {
		return nil, nil
	}

	return &account, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByVerificationUUID(uuid string) (*models.Account, error) {
	var account models.Account

	// Raw SQL query: join verify and accounts, filter by verify_uuid
	const q = `
            SELECT accounts.* FROM accounts
            INNER JOIN verify ON verify.fid_account = accounts.id_account
            WHERE verify.verify_uuid = ?
            LIMIT 1
      `

	result := app.SQLDB["main"].DB.Raw(q, uuid).Scan(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	if account.ID == 0 {
		return nil, nil // Not found
	}
	return &account, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Update User Identity
//||------------------------------------------------------------------------------------------------||

func UpdateUserIdentity(idAccount int64, ident verify.Identity) error {
	blob, err := json.Marshal(ident)
	if err != nil {
		return err
	}

	tx := app.SQLDB["main"].DB.
		Model(&models.Account{}).
		Where("id_account = ?", idAccount).
		Update("account_identity", string(blob))

	if tx.Error != nil {
		return tx.Error
	}
	if tx.RowsAffected == 0 {
		return gorm.ErrRecordNotFound
	}
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| DeleteAccount – deletes an account and cascades deletions where necessary
//||------------------------------------------------------------------------------------------------||

func DeleteAccount(accountID int64) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Start transaction to ensure atomic delete
	//||------------------------------------------------------------------------------------------------||

	tx := app.SQLDB["main"].DB.Begin()
	if tx.Error != nil {
		return fmt.Errorf("failed to start transaction: %w", tx.Error)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Explicitly delete any shared records referencing verifications tied to this account
	//||------------------------------------------------------------------------------------------------||

	if err := tx.Exec(`
        DELETE s FROM shared s
        JOIN verifications v ON v.id_verification = s.fid_verification
        WHERE v.fid_account = ?`, accountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete shared records: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Delete the account itself (verifications cascade via FK ON DELETE CASCADE)
	//||------------------------------------------------------------------------------------------------||

	if err := tx.Exec(`DELETE FROM accounts WHERE id_account = ?`, accountID).Error; err != nil {
		tx.Rollback()
		return fmt.Errorf("failed to delete account: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Commit transaction
	//||------------------------------------------------------------------------------------------------||
	if err := tx.Commit().Error; err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}
