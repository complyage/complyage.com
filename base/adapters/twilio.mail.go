package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| SendGridSendEmail : Send an email using Twilio SendGrid API
//||------------------------------------------------------------------------------------------------||

func SendGridSendEmail(toEmail, subject, body string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request Payload
	//||------------------------------------------------------------------------------------------------||

	request := SendGridRequest{
		Personalizations: []SendGridPersonalization{
			{
				To: []SendGridEmail{
					{Email: toEmail},
				},
			},
		},
		From:    SendGridEmail{Email: os.Getenv("SENDGRID_FROM_EMAIL")},
		Subject: subject,
		Content: []SendGridContent{
			{Type: "text/plain", Value: body},
		},
	}

	b, _ := json.Marshal(request)

	//||------------------------------------------------------------------------------------------------||
	//|| Create HTTP Request
	//||------------------------------------------------------------------------------------------------||

	req, err := http.NewRequest("POST", "https://api.sendgrid.com/v3/mail/send", bytes.NewBuffer(b))
	if err != nil {
		return "", fmt.Errorf("sendgrid new request: %w", err)
	}
	req.Header.Set("Authorization", "Bearer "+os.Getenv("SENDGRID_API_KEY"))
	req.Header.Set("Content-Type", "application/json")

	//||------------------------------------------------------------------------------------------------||
	//|| Perform Request
	//||------------------------------------------------------------------------------------------------||

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("sendgrid http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	//||------------------------------------------------------------------------------------------------||
	//|| Check Response
	//||------------------------------------------------------------------------------------------------||

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("sendgrid error %s: %s", resp.Status, string(respBody))
	}

	return string(respBody), nil
}
