package public

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"os"
	"strings"

	"github.com/complyage/base/adapters"

	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Structs
//||------------------------------------------------------------------------------------------------||

type ContactRequest struct {
	Name           string `json:"name"`
	Email          string `json:"email"`
	Message        string `json:"message"`
	TurnstileToken string `json:"turnstileToken"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Contact Form
//||------------------------------------------------------------------------------------------------||

func ContactHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Contact Request
	//||------------------------------------------------------------------------------------------------||

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Fields
	//||------------------------------------------------------------------------------------------------||

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	//||------------------------------------------------------------------------------------------------||
	//|| Basic Validation
	//||------------------------------------------------------------------------------------------------||

	if req.Name == "" || req.Email == "" || req.Message == "" {
		responses.Error(w, http.StatusBadRequest, "Name, email, and message are required")
		return
	}
	if !strings.Contains(req.Email, "@") {
		responses.Error(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Turnstile Captcha
	//||------------------------------------------------------------------------------------------------||

	ip := r.RemoteAddr
	if err := adapters.VerifyTurnstile(req.TurnstileToken, ip); err != nil {
		responses.Error(w, http.StatusBadRequest, "Captcha verification failed: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Generate the Message
	//||------------------------------------------------------------------------------------------------||

	message := "You have received a new contact form submission:\n\n"
	message += "Name: \n" + req.Name + "\n"
	message += "Email: \n" + req.Email + "\n"
	message += "Message:\n" + req.Message + "\n"

	//||------------------------------------------------------------------------------------------------||
	//|| Send Email
	//||------------------------------------------------------------------------------------------------||

	adapters.SendGridSendEmail(
		os.Getenv("CONTACT_EMAIL_ADDRESS"),
		"New Contact Form Submission : complyage.com",
		message,
	)

	//||------------------------------------------------------------------------------------------------||
	//|| Success Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Thank you for contacting us! We'll reply as soon as possible.",
	})
}
