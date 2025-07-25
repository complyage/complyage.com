package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"base/db"
	"base/helpers"
	"base/interfaces"
	"base/responses"
	"encoding/json"
	"net/http"
	"os"
	"strings"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Serve
//||------------------------------------------------------------------------------------------------||

func ServePrivateKeyForm(w http.ResponseWriter, r *http.Request) {
	b, err := os.ReadFile("assets/private.html")
	if err != nil {
		responses.ErrorHTML(w, "template not found")
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(b)
}

//||------------------------------------------------------------------------------------------------||
//|| Submit Private Key Handler
//||------------------------------------------------------------------------------------------------||

func SubmitPrivateKeyHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse Form
	//||------------------------------------------------------------------------------------------------||

	err := r.ParseForm()
	if err != nil {
		responses.ErrorHTML(w, "Invalid form data")
		return
	}

	oauth := r.FormValue("oauth")
	privateKey := strings.TrimSpace(r.FormValue("private"))

	if oauth == "" || privateKey == "" {
		responses.ErrorHTML(w, "Missing OAuth session or private key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Load OAuth Session from Redis
	//||------------------------------------------------------------------------------------------------||

	val, err := db.Redis.Get(r.Context(), "oauth:"+oauth).Result()
	if err != nil {
		responses.ErrorHTML(w, "Invalid or expired OAuth session")
		return
	}

	var session interfaces.OAuthSession
	err = json.Unmarshal([]byte(val), &session)
	if err != nil {
		responses.ErrorHTML(w, "Failed to decode OAuth session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Private Key
	//||------------------------------------------------------------------------------------------------||

	pkErr := helpers.CheckPrivateKey(privateKey, session.PrivateCheck)
	if pkErr != nil {
		responses.ErrorHTML(w, "Invalid private key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Update Session With Validated Private Key
	//||------------------------------------------------------------------------------------------------||

	session.Private = privateKey

	updated, err := json.Marshal(session)
	if err != nil {
		responses.ErrorHTML(w, "Failed to encode updated session")
		return
	}

	err = db.Redis.Set(r.Context(), "oauth:"+oauth, updated, 60*time.Minute).Err()
	if err != nil {
		responses.ErrorHTML(w, "Failed to update session")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Redirect to Approval
	//||------------------------------------------------------------------------------------------------||

	http.Redirect(w, r, "/v1/approve?oauth="+oauth, http.StatusSeeOther)
}
