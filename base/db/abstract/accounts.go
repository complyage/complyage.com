package abstract

import (
	"fmt"

	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/auth/db"
	"gorm.io/gorm"
)

//||------------------------------------------------------------------------------------------------||
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByID(id string) (*db.ModelAccount, error) {
	var account db.ModelAccount

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
//|| Get Account Based on ID
//||------------------------------------------------------------------------------------------------||

func GetAccountByVerificationUUID(uuid string) (*db.ModelAccount, error) {
	var account db.ModelAccount

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
