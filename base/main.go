package main

import (
	"fmt"
	"os"

	"github.com/complyage/base/adapters"
	"github.com/complyage/base/types"

	"github.com/joho/godotenv"
	"github.com/ralphferrara/aria/app"
)

// ||------------------------------------------------------------------------------------------------||
// || Main
// ||------------------------------------------------------------------------------------------------||
func main() {

	//||------------------------------------------------------------------------------------------------||
	//|| Load environment variables
	//||------------------------------------------------------------------------------------------------||
	if err := godotenv.Load("../.env"); err != nil {
		fmt.Println("No .env file found, continuing...")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Initialize application and Stripe
	//||------------------------------------------------------------------------------------------------||
	app.Init("../config.json")
	adapters.InitStripe()

	//||------------------------------------------------------------------------------------------------||
	//|| CLI Handler
	//||------------------------------------------------------------------------------------------------||
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	cmd := os.Args[1]

	switch cmd {

	//----------------------------------------------------------------------------------------------//
	// Help
	//----------------------------------------------------------------------------------------------//

	case "help", "--help", "-h":
		printHelp()

	//----------------------------------------------------------------------------------------------//
	// Test ClickSend SMS
	//----------------------------------------------------------------------------------------------//

	case "test-sms":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go test-sms <toPhone> <message>")
			return
		}
		to := os.Args[2]
		body := os.Args[3]
		resp, err := adapters.SendText(to, body)
		exitWith(resp, err)

	//----------------------------------------------------------------------------------------------//
	// Test ClickSend Email
	//----------------------------------------------------------------------------------------------//

	case "test-email":
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go test-email <to> <from> <subject>")
			return
		}
		to := os.Args[2]
		from := os.Args[3]
		subject := os.Args[4]
		resp, err := adapters.SendEmail(to, from, subject, "Test email body from ComplyAge CLI")
		exitWith(resp, err)

	//----------------------------------------------------------------------------------------------//
	// Test SendGrid Email
	//----------------------------------------------------------------------------------------------//

	case "test-sendgrid":
		fmt.Println("Sending SendGrid test email...")

		resp, err := adapters.SendGridSendMail(
			os.Args[2],
			os.Getenv("SENDGRID_FROM_EMAIL"),
			"ComplyAge SendGrid Test",
			"This is a SendGrid test email from ComplyAge CLI.",
			"This is a SendGrid test email from ComplyAge CLI.",
		)
		exitWith(resp, err)

	//----------------------------------------------------------------------------------------------//
	// Test Twilio TXT
	//----------------------------------------------------------------------------------------------//

	case "test-txt":
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go test-txt <to> <from> <message>")
			return
		}
		to := os.Args[2]
		from := os.Args[3]
		body := os.Args[4]
		resp, err := adapters.SendGridSendTXT(to, from, body)
		exitWith(resp, err)

	//----------------------------------------------------------------------------------------------//
	// Test Turnstile
	//----------------------------------------------------------------------------------------------//
	case "test-turnstile":
		if len(os.Args) < 4 {
			fmt.Println("Usage: go run main.go test-turnstile <token> <ip>")
			return
		}
		token := os.Args[2]
		ip := os.Args[3]
		err := adapters.VerifyTurnstile(token, ip)
		if err != nil {
			fmt.Println("❌ Turnstile failed:", err)
		} else {
			fmt.Println("✅ Turnstile verified successfully.")
		}

	//----------------------------------------------------------------------------------------------//
	// Test Stripe Intent
	//----------------------------------------------------------------------------------------------//
	case "test-stripe":
		if len(os.Args) < 5 {
			fmt.Println("Usage: go run main.go test-stripe <uuid> <amount> <currency>")
			return
		}
		uuid := os.Args[2]
		amount := os.Args[3]
		currency := os.Args[4]

		var amt float64
		fmt.Sscanf(amount, "%f", &amt)

		intent, err := adapters.StripeIntent(uuid, amt, currency, "TEST")
		if err != nil {
			fmt.Println("❌ Stripe error:", err)
		} else {
			fmt.Println("✅ Stripe intent created:", intent.ID)
		}

	//----------------------------------------------------------------------------------------------//
	// Test ClickSend Postcard
	//----------------------------------------------------------------------------------------------//
	case "test-postcard":
		fmt.Println("Testing ClickSend postcard...")
		addr := types.Address{
			Line1:   "123 Main St",
			Line2:   "Apt 4B",
			City:    "Phoenix",
			State:   "AZ",
			Postal:  "85022",
			Country: "US",
		}

		resp, err := adapters.ClickSendPostcard(
			addr,
			"test-uuid",
			"./.assets/postcard-template.png",
			"https://verify.com/",
			"123456",
			os.Getenv("CLICKSEND_USERNAME"),
			os.Getenv("CLICKSEND_API_KEY"),
		)
		exitWith(resp, err)

	default:
		fmt.Printf("Unknown command: %s\n\n", cmd)
		printHelp()
	}
}

//||------------------------------------------------------------------------------------------------||
//|| Helpers
//||------------------------------------------------------------------------------------------------||

func printHelp() {
	fmt.Println("ComplyAge Base Adapters CLI")
	fmt.Println("---------------------------------------------------")
	fmt.Println("Usage:")
	fmt.Println("  go run main.go <command> [arguments]")
	fmt.Println()
	fmt.Println("Commands:")
	fmt.Println("  help                      Show this help message")
	fmt.Println("  test-sms <to> <msg>       Send SMS via ClickSend")
	fmt.Println("  test-email <to> <from> <subject>   Send email via ClickSend")
	fmt.Println("  test-sendgrid <to> <from> <subject>  Send email via SendGrid")
	fmt.Println("  test-txt <to> <from> <msg>         Send SMS via Twilio")
	fmt.Println("  test-turnstile <token> <ip>        Verify Turnstile token manually")
	fmt.Println("  test-stripe <uuid> <amount> <currency>  Create a Stripe PaymentIntent")
	fmt.Println("  test-postcard             Send ClickSend postcard to test address")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println("  go run main.go test-sendgrid test@example.com noreply@domain.com \"Subject\"")
	fmt.Println("  go run main.go test-sms +15551234567 \"Hello from CLI\"")
	fmt.Println("---------------------------------------------------")
}

func exitWith(resp string, err error) {
	if err != nil {
		fmt.Println("❌ Error:", err)
	} else {
		fmt.Println("✅ Response:", resp)
	}
}
