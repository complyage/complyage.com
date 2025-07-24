package helpers

import (
	"base/db"
	"base/interfaces"
	"base/models"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| Get the Account Record
//||------------------------------------------------------------------------------------------------||

func SessionCreate(email string, account models.Account) (string, error) {

	//||------------------------------------------------------------------------------------------------||
	//|| Generate a Random Token
	//||------------------------------------------------------------------------------------------------||

	rawToken, err := GenerateRandom()
	if err != nil {
		fmt.Println("[Session] Failed to generate session token:", err)
		return "", err
	}

	sessionToken := base64.URLEncoding.EncodeToString(rawToken)

	//||------------------------------------------------------------------------------------------------||
	//|| Check if Exists
	//||------------------------------------------------------------------------------------------------||

	session := interfaces.SessionRecord{
		Email:    email,
		Username: account.AccountUsername,
		ID:       account.IDAccount,
		Status:   account.AccountStatus,
		Type:     account.AccountType,
		Level:    DerefInt8(account.AccountLevel),
		Advanced: account.AccountAdvanced != nil && *account.AccountAdvanced == 1,
		Created:  time.Now().Unix(),
		Expires:  time.Now().Add(30 * 24 * time.Hour).Unix(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Marshal
	//||------------------------------------------------------------------------------------------------||

	sessionJSON, err := json.Marshal(session)
	if err != nil {
		fmt.Println("[Session] Failed to marshal session:", err)
		return "", err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save to Redis
	//||------------------------------------------------------------------------------------------------||

	err = db.Redis.Set(context.Background(), "session::"+sessionToken, sessionJSON, 30*24*time.Hour).Err()
	if err != nil {
		fmt.Println("[Session] Failed to save session to Redis:", err)
		return "", err
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Done
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("[Session] Token set for account:", sessionToken)
	return sessionToken, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Fetch Session
//||------------------------------------------------------------------------------------------------||

func FetchSession(sessionID string) (interfaces.SessionRecord, error) {
	//||------------------------------------------------------------------------------------------------||
	//|| Get the Session from the Database
	//||------------------------------------------------------------------------------------------------||

	sessionJSON, err := db.Redis.Get(context.Background(), "session::"+sessionID).Result()
	if err != nil {
		return interfaces.SessionRecord{}, fmt.Errorf("failed to fetch session: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return the Session
	//||------------------------------------------------------------------------------------------------||

	var session interfaces.SessionRecord
	if err := json.Unmarshal([]byte(sessionJSON), &session); err != nil {
		return interfaces.SessionRecord{}, fmt.Errorf("failed to unmarshal session: %w", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Check Session
	//||------------------------------------------------------------------------------------------------||

	if time.Unix(session.Expires, 0).Before(time.Now()) {
		return interfaces.SessionRecord{}, fmt.Errorf("session expired")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Return the Session
	//||------------------------------------------------------------------------------------------------||

	return session, nil
}

//||------------------------------------------------------------------------------------------------||
//|| Create the
//||------------------------------------------------------------------------------------------------||

func WriteSessionCookie(w http.ResponseWriter, sessionToken string) {
	//||------------------------------------------------------------------------------------------------||
	//|| Set the Cookie
	//||------------------------------------------------------------------------------------------------||

	if sessionToken == "" {
		fmt.Println("[Session] No session token provided")
		return
	}

	fmt.Println("[Session] Setting cookie with token:", sessionToken)

	//||------------------------------------------------------------------------------------------------||
	//|| Create and set the cookie
	//||------------------------------------------------------------------------------------------------||

	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    sessionToken,
		Path:     "/",
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteLaxMode,
		MaxAge:   86400 * 30,
	})

}

//||------------------------------------------------------------------------------------------------||
//|| Delete Session
//||------------------------------------------------------------------------------------------------||

func DeleteSession(sessionToken string) error {
	ctx := context.Background()
	// Delete the session token from Redis
	return db.Redis.Del(ctx, "session::"+sessionToken).Err()
}

//||------------------------------------------------------------------------------------------------||
//|| Clear Session Cookie
//||------------------------------------------------------------------------------------------------||

func ClearSessionCookie(w http.ResponseWriter) {
	// Create a cookie that is already expired
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Path:     "/",
		Expires:  time.Unix(0, 0), // Set expiry to the Unix epoch (already expired)
		MaxAge:   -1,              // Instructs the browser to delete the cookie immediately
		HttpOnly: true,
		Secure:   true, // Should be true in production (HTTPS)
		SameSite: http.SameSiteLaxMode,
	})
}
