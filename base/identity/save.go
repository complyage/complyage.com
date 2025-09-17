package identity

import (
	"fmt"

	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Save to Minio / S3
//||------------------------------------------------------------------------------------------------||

func (i *Identity) Save() error {
	fmt.Println("Saving Identity Record")

	jsonStr, err := i.toJSON()
	if err != nil {
		fmt.Println("Error converting identity to JSON:", err)
		return err
	}

	data := []byte(jsonStr)
	if err := app.Storages["identity"].Put(i.Filename(), data); err != nil {
		fmt.Println("Error saving identity to storage:", err)
		return err
	}

	return nil
}
