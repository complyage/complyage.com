//==================================================================================================
// adapters/sendgrid.go
//==================================================================================================

package adapters

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

//||------------------------------------------------------------------------------------------------||
//|| SendGridSendMail : Send email using SendGrid API
//||------------------------------------------------------------------------------------------------||
//| Params: to (recipient email), from (sender email), subject (email subject), bodyText (plain text),
//|          bodyHTML (HTML content, optional)
//| Returns: response body or error
//||------------------------------------------------------------------------------------------------||

func SendGridSendMail(to string, from string, subject string, bodyText string, bodyHTML string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Load Config
	//||------------------------------------------------------------------------------------------------||

	apiKey := os.Getenv("SENDGRID_API_KEY")
	if apiKey == "" {
		return "", fmt.Errorf("missing SendGrid credentials: SENDGRID_API_KEY")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Message
	//||------------------------------------------------------------------------------------------------||

	fromEmail := mail.NewEmail("", from)
	toEmail := mail.NewEmail("", to)

	// Use plain text as fallback if HTML not provided
	if bodyHTML == "" {
		bodyHTML = fmt.Sprintf("<pre>%s</pre>", bodyText)
	}

	message := mail.NewSingleEmail(fromEmail, subject, toEmail, bodyText, bodyHTML)

	//||------------------------------------------------------------------------------------------------||
	//|| Send Message
	//||------------------------------------------------------------------------------------------------||

	client := sendgrid.NewSendClient(apiKey)
	response, err := client.Send(message)
	if err != nil {
		return "", fmt.Errorf("sendgrid send: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Handle Response
	//||------------------------------------------------------------------------------------------------||

	if response.StatusCode >= 300 {
		return "", fmt.Errorf("sendgrid error %d: %s", response.StatusCode, response.Body)
	}

	return response.Body, nil
}
