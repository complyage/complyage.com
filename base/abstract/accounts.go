package abstract

import (
	"base/db"
	"base/models"
	"fmt"

	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByID(id string) (*models.Account, error) {
	var account models.Account

	result := db.DB.First(&account, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // Not found, not a DB error
		}
		return nil, result.Error
	}

	return &account, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByVerificationUUID(uuid string) (*models.Account, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Prepare Account
	//||------------------------------------------------------------------------------------------------||

	var account models.Account

	//||------------------------------------------------------------------------------------------------||
	//|| Query: Join verifications and accounts by fid_account, filter by uuid
	//||------------------------------------------------------------------------------------------------||

	q := `
		SELECT 
			A.id_account, 
			A.account_username,
			A.account_public,
			A.account_private,
			A.account_private_hash
		FROM accounts A
		LEFT JOIN verifications V ON A.id_account = V.fid_account
		WHERE V.verification_uuid = ?
		LIMIT 1
	`

	//||------------------------------------------------------------------------------------------------||
	//|| Query DB
	//||------------------------------------------------------------------------------------------------||

	result := db.DB.Raw(q, uuid).Scan(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	// Not found if zero ID
	if account.IDAccount == 0 {
		return nil, nil
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	return &account, nil
}

//||------------------------------------------------------------------------------------------------||
//|| DeleteAccount – deletes an account and cascades deletions where necessary
//||------------------------------------------------------------------------------------------------||

func DeleteAccount(accountID int64) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Start transaction to ensure atomic delete
	//||------------------------------------------------------------------------------------------------||

	tx := db.DB.Begin()
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
