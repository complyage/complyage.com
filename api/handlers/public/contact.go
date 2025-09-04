package public

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Structs
//||------------------------------------------------------------------------------------------------||

type ContactRequest struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Message string `json:"message"`
}

//||------------------------------------------------------------------------------------------------||
//|| Handler: Contact Form
//||------------------------------------------------------------------------------------------------||

func ContactHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		responses.Error(w, http.StatusMethodNotAllowed, "Invalid request method")
		return
	}

	var req ContactRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	// Basic validation
	if req.Name == "" || req.Email == "" || req.Message == "" {
		responses.Error(w, http.StatusBadRequest, "Name, email, and message are required")
		return
	}
	if !strings.Contains(req.Email, "@") {
		responses.Error(w, http.StatusBadRequest, "Invalid email address")
		return
	}

	// TODO: Implement your mail sending logic here
	// err := sendMail(req.Name, req.Email, req.Message)
	// if err != nil {
	//     responses.Error(w, http.StatusInternalServerError, "Failed to send message")
	//     return
	// }

	responses.Success(w, http.StatusOK, map[string]any{
		"success": true,
		"message": "Thank you for contacting us! We'll reply as soon as possible.",
	})
}
