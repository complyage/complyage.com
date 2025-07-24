package handlers

import (
	"fmt"
	"net/http"

	"base/helpers"
	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Logout Handler
//||------------------------------------------------------------------------------------------------||

func LogoutHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session Cookie
	//||------------------------------------------------------------------------------------------------||

	cookie, err := r.Cookie("session")
	if err != nil || cookie.Value == "" {
		responses.Success(w, http.StatusOK, map[string]any{
			"message": "Already logged out",
			"next":    "/login", // Redirect to login page
		})
		return
	}

	sessionToken := cookie.Value
	fmt.Println("[Logout] Incoming -> sessionToken:", sessionToken)

	//||------------------------------------------------------------------------------------------------||
	//|| Delete Session from Redis
	//||------------------------------------------------------------------------------------------------||

	err = helpers.DeleteSession(sessionToken)
	if err != nil {
		fmt.Printf("[Logout] Failed to delete session %s from Redis: %v\n", sessionToken, err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Clear the Session Cookie from the browser
	//||------------------------------------------------------------------------------------------------||

	helpers.ClearSessionCookie(w)

	//||------------------------------------------------------------------------------------------------||
	//|| Success Response
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"message": "Logged out successfully",
		"next":    "/login", // Redirect to login page after logout
	})
}
