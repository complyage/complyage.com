package verify

import (
	"base/db/models"

	"github.com/ralphferrara/aria/app"
)

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
