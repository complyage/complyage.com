package adapters

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| SendGridSendTXT : Send SMS using Twilio API
//||------------------------------------------------------------------------------------------------||
//| Params: to (E.164 format, ex: +15551234567), from (your Twilio phone #), body (text message)
//| Returns: response body or error
//||------------------------------------------------------------------------------------------------||

func SendGridSendTXT(to string, from string, body string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Load Config
	//||------------------------------------------------------------------------------------------------||

	accountSID := os.Getenv("TWILIO_ACCOUNT_SID")
	authToken := os.Getenv("TWILIO_AUTH_TOKEN")

	if accountSID == "" || authToken == "" {
		return "", fmt.Errorf("missing Twilio credentials: TWILIO_ACCOUNT_SID or TWILIO_AUTH_TOKEN")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Build Request
	//||------------------------------------------------------------------------------------------------||

	twilioURL := fmt.Sprintf("https://api.twilio.com/2010-04-01/Accounts/%s/Messages.json", accountSID)

	msgData := url.Values{}
	msgData.Set("To", to)
	msgData.Set("From", from)
	msgData.Set("Body", body)

	req, err := http.NewRequest("POST", twilioURL, strings.NewReader(msgData.Encode()))
	if err != nil {
		return "", fmt.Errorf("twilio new request: %w", err)
	}
	req.SetBasicAuth(accountSID, authToken)
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//||------------------------------------------------------------------------------------------------||
	//|| Send Request
	//||------------------------------------------------------------------------------------------------||

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("twilio http: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)

	//||------------------------------------------------------------------------------------------------||
	//|| Check Response
	//||------------------------------------------------------------------------------------------------||

	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("twilio error %s: %s", resp.Status, string(respBody))
	}

	return string(respBody), nil
}
