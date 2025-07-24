package helpers

import "fmt"

//||------------------------------------------------------------------------------------------------||
//|| Email Helper, Replace with Mailer Call
//||------------------------------------------------------------------------------------------------||

func EmailPrivateKeyToUser(email, privateKey string) error {
	// In real code, send via SMTP or your mail service.
	fmt.Println("=== EMAIL PRIVATE KEY ===")
	fmt.Println("TO:", email)
	fmt.Println("KEY:", privateKey)
	fmt.Println("=========================")
	return nil
}
