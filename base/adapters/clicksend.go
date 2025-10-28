//||------------------------------------------------------------------------------------------------||
//|| ClickSendPostcard sends a postcard PNG to the recipient address
//||------------------------------------------------------------------------------------------------||
//| Returns: response body as string, or error
//||------------------------------------------------------------------------------------------------||

package adapters

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"path/filepath"

	"github.com/complyage/base/types"
)

//||------------------------------------------------------------------------------------------------||
//|| ClickSend Address struct
//||------------------------------------------------------------------------------------------------||

type ClickSendAddress struct {
	Name     string `json:"name"`
	Address1 string `json:"address_line_1"`
	Address2 string `json:"address_line_2"`
	City     string `json:"address_city"`
	State    string `json:"address_state"`
	Postal   string `json:"address_postal_code"`
	Country  string `json:"address_country"`
}

//||------------------------------------------------------------------------------------------------||
//|| ClickSend Address struct
//||------------------------------------------------------------------------------------------------||

type ClickSendSMSRequest struct {
	Messages []ClickSendSMSMessage `json:"messages"`
}

//||------------------------------------------------------------------------------------------------||
//|| ClickSend SMS Message struct
//||------------------------------------------------------------------------------------------------||

type ClickSendSMSMessage struct {
	Source       string `json:"source,omitempty"` // Optional
	From         string `json:"from,omitempty"`   // Optional
	Body         string `json:"body"`
	To           string `json:"to"` // E.164 format: +15551234567
	Schedule     int    `json:"schedule,omitempty"`
	CustomString string `json:"custom_string,omitempty"`
}

//||------------------------------------------------------------------------------------------------||
//|| Send Email Request
//||------------------------------------------------------------------------------------------------||

type ClickSendEmailRequest struct {
	To      string `json:"to"`   // destination email
	From    string `json:"from"` // your sending address
	Subject string `json:"subject"`
	Body    string `json:"body"`
}

//||------------------------------------------------------------------------------------------------||
//|| ClickSend Address struct
//||------------------------------------------------------------------------------------------------||

func ClickSendPostcard(toAddress types.Address, verificationUUID string, templatePath string, verifyURL string, checkCode string, clickSendUsername string, clickSendAPIKey string) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Create the Back PNG
	//||------------------------------------------------------------------------------------------------||

	backPng, err := GeneratePostcardBack(toAddress, verificationUUID, checkCode)
	if err != nil {
		return "", fmt.Errorf("generate back: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create the From Address
	//||------------------------------------------------------------------------------------------------||

	from := ClickSendAddress{
		Name:     os.Getenv("VERIFICATION_ADDRESS_RETURN_NAME"),
		Address1: os.Getenv("VERIFICATION_ADDRESS_RETURN_ADDRESS1"),
		Address2: os.Getenv("VERIFICATION_ADDRESS_RETURN_ADDRESS2"),
		City:     os.Getenv("VERIFICATION_ADDRESS_RETURN_CITY"),
		State:    os.Getenv("VERIFICATION_ADDRESS_RETURN_STATE"),
		Postal:   os.Getenv("VERIFICATION_ADDRESS_RETURN_POSTAL"),
		Country:  os.Getenv("VERIFICATION_ADDRESS_RETURN_COUNTRY"),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Convert the To Address
	//||------------------------------------------------------------------------------------------------||

	to := ClickSendAddress{
		Name:     "Current Resident",
		Address1: toAddress.Line1,
		Address2: toAddress.Line2,
		City:     toAddress.City,
		State:    toAddress.State,
		Postal:   toAddress.Postal,
		Country:  toAddress.Country,
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Open the front PNG file
	//||------------------------------------------------------------------------------------------------||

	frontPath := "./.assets/postcard-front.png"
	frontFile, err := os.Open(frontPath)
	if err != nil {
		return "", fmt.Errorf("open front png: %w", err)
	}
	defer frontFile.Close()

	//||------------------------------------------------------------------------------------------------||
	//|| Start a multipart writer
	//||------------------------------------------------------------------------------------------------||

	var body bytes.Buffer
	writer := multipart.NewWriter(&body)

	//||------------------------------------------------------------------------------------------------||
	//|| Add Front file
	//||------------------------------------------------------------------------------------------------||

	frontPart, err := writer.CreateFormFile("front", filepath.Base(frontPath))
	if err != nil {
		return "", fmt.Errorf("form front: %w", err)
	}
	if _, err := io.Copy(frontPart, frontFile); err != nil {
		return "", fmt.Errorf("copy front: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Back File
	//||------------------------------------------------------------------------------------------------||

	backPart, err := writer.CreateFormFile("back", "postcard-back.png")
	if err != nil {
		return "", fmt.Errorf("form back: %w", err)
	}
	if _, err := backPart.Write(backPng); err != nil {
		return "", fmt.Errorf("write back: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Create JSON
	//||------------------------------------------------------------------------------------------------||

	postcardPayload := map[string]interface{}{
		"to":   to,
		"from": from,
	}
	payloadBytes, _ := json.Marshal(postcardPayload)
	_ = writer.WriteField("postcard", string(payloadBytes))

	writer.Close()

	req, err := http.NewRequest("POST", "https://rest.clicksend.com/v3/post/postcards/send", &body)
	if err != nil {
		return "", fmt.Errorf("http.NewRequest: %w", err)
	}
	req.SetBasicAuth(clickSendUsername, clickSendAPIKey)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("http.Post: %w", err)
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("ClickSend error %s: %s", resp.Status, string(respBody))
	}
	return string(respBody), nil
}

//||------------------------------------------------------------------------------------------------||
//|| Generate the Back File
//||------------------------------------------------------------------------------------------------||

func SendText(toPhone string, body string) (string, error) {

	clickSendUsername := os.Getenv("CLICKSEND_USERNAME")
	clickSendAPIKey := os.Getenv("CLICKSEND_API_KEY")

	payload := ClickSendSMSRequest{
		Messages: []ClickSendSMSMessage{
			{
				Body: body,
				To:   toPhone,
			},
		},
	}
	b, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://rest.clicksend.com/v3/sms/send", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clickSendUsername, clickSendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("ClickSend SMS error %s: %s", resp.Status, string(respBody))
	}
	return string(respBody), nil
}

//||------------------------------------------------------------------------------------------------||
//|| Send Email
//||------------------------------------------------------------------------------------------------||

func SendEmail(toEmail, fromEmail, subject, body string) (string, error) {
	payload := map[string]interface{}{
		"to":      toEmail,
		"from":    fromEmail,
		"subject": subject,
		"body":    body,
	}

	clickSendUsername := os.Getenv("CLICKSEND_USERNAME")
	clickSendAPIKey := os.Getenv("CLICKSEND_API_KEY")

	b, _ := json.Marshal(payload)
	req, err := http.NewRequest("POST", "https://rest.clicksend.com/v3/email/send", bytes.NewBuffer(b))
	if err != nil {
		return "", err
	}
	req.SetBasicAuth(clickSendUsername, clickSendAPIKey)
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode >= 300 {
		return "", fmt.Errorf("ClickSend Email error %s: %s", resp.Status, string(respBody))
	}
	return string(respBody), nil
}
