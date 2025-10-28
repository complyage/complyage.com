package identity

import (
	"fmt"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Save to Minio / S3
//||------------------------------------------------------------------------------------------------||

func Create(accountId int64) (Identity, error) {
	fmt.Println("Creating Identity Record")
	i := Identity{}
	i.ID = accountId
	err := i.Save()
	if err != nil {
		app.Log.Info("Error creating identity record:", err)
		return Identity{}, err
	}
	return i, nil
}
