package send

import (
	"fmt"

	"github.com/ralphferrara/aria/base/bip39"
)

//||------------------------------------------------------------------------------------------------||
//|| Email Helper, Replace with Mailer Call
//||------------------------------------------------------------------------------------------------||

func EmailBIPListToUser(email string, bipList bip39.BIP39WordList, privateKey string) error {
	// In real code, send via SMTP or your mail service.
	fmt.Println("=== EMAIL PRIVATE KEY ===")
	fmt.Println("TO:", email)
	fmt.Println("List:", bipList)
	fmt.Println("Key:", privateKey)
	fmt.Println("=========================")
	return nil
}
