package communicate

import (
	"fmt"
	"net/http"
	"os"

	"github.com/complyage/base/adapters"
	"github.com/ralphferrara/aria/app"
	"github.com/ralphferrara/aria/base/validate"
)

func SendTwoFactorCode(r *http.Request, accountID int64, identifier string, code string) error {
	if validate.IsEmailOrPhone(identifier) == "email" {
		// Send via SendGrid Email
		from := os.Getenv("SENDGRID_FROM_EMAIL")
		subject := app.Config.App.Name + " - Your Two-Factor Authentication Code - " + code
		bodyText := fmt.Sprintf("Your two-factor authentication code is: %s", code)
		bodyHTML := fmt.Sprintf("<p>Your two-factor authentication code is: <strong>%s</strong></p>", code)
		_, err := adapters.SendGridSendMail(identifier, from, subject, bodyText, bodyHTML)
		if err != nil {
			return fmt.Errorf("failed to send two-factor code via email: %w", err)
		}
		return nil
	} else if validate.IsEmailOrPhone(identifier) == "phone" {
		// Send via Twilio SMS
		return nil
	}
	return fmt.Errorf("invalid identifier for two-factor code: %s", identifier)
}
