package helpers

import (
	"base/interfaces"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Email Helper, Replace with Mailer Call
//||------------------------------------------------------------------------------------------------||

func EmailBIPListToUser(email string, bipList interfaces.BIPList, privateKey string) error {
	// In real code, send via SMTP or your mail service.
	fmt.Println("=== EMAIL PRIVATE KEY ===")
	fmt.Println("TO:", email)
	fmt.Println("List:", bipList)
	fmt.Println("Key:", privateKey)
	fmt.Println("=========================")
	return nil
}

//||------------------------------------------------------------------------------------------------||
//|| Email Helper, Replace with Mailer Call
//||------------------------------------------------------------------------------------------------||

func EmailPrivateKeyToUser(email string, privateKey string) error {
	// In real code, send via SMTP or your mail service.
	fmt.Println("=== EMAIL PRIVATE KEY ===")
	fmt.Println("TO:", email)
	fmt.Println("KEY:", privateKey)
	fmt.Println("=========================")
	return nil
}
