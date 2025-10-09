package webhook

import "time"

type WebhookPayload struct {
	ID            string    `json:"id"`
	ClientID      string    `json:"clientId"`
	CheckToken    string    `json:"verifier"`
	SessionID     string    `json:"sessionId"`
	Username      string    `json:"username"`
	IPAddress     string    `json:"ipAddress"`
	Region        string    `json:"region"`
	Country       string    `json:"country"`
	DOB           string    `json:"dob"`
	Age           int       `json:"age"`
	Method        string    `json:"method"`
	Selfie        string    `json:"selfie"`
	IDFront       string    `json:"idFront"`
	TransactionID string    `json:"transactionId"`
	LastFour      string    `json:"lastFour"`
	Timestamp     time.Time `json:"timestamp"`
}
