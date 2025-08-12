package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/agent_interfaces"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

//||------------------------------------------------------------------------------------------------||
//|| Agent Callback Handler :: Receives data from Python agent and displays all fields
//||------------------------------------------------------------------------------------------------||

func AgentCallbackHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("Call to Callback handler")

	//||------------------------------------------------------------------------------------------------||
	//|| Header
	//||------------------------------------------------------------------------------------------------||

	authHeader := r.Header.Get("Authorization")
	if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
		fmt.Println("⚠️ WARNING: Authorization header is missing or malformed")
		http.Error(w, "Unauthorizated", http.StatusUnauthorized)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Shared Token / Verify
	//||------------------------------------------------------------------------------------------------||

	token := strings.TrimPrefix(authHeader, "Bearer ")
	expectedKey := os.Getenv("AGENT_SHARED_KEY")
	if expectedKey == "" {
		fmt.Println("⚠️ WARNING: AGENT_SHARED_KEY is not set in the environment!")
	}

	if token != expectedKey {
		fmt.Println("⚠️ WARNING: Authorization code doesn't match")
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Agent Callback Handler :: Receives data from Python agent and displays all fields
	//||------------------------------------------------------------------------------------------------||

	if r.Method != http.MethodPost {
		fmt.Println("⚠️ WARNING: Method not allowed")
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Agent Callback Handler :: Receives data from Python agent and displays all fields
	//||------------------------------------------------------------------------------------------------||

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("⚠️ WARNING: Failed to read request body:", err)
		http.Error(w, "Failed to read request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()

	//||------------------------------------------------------------------------------------------------||
	//|| Raw
	//||------------------------------------------------------------------------------------------------||

	var pretty map[string]interface{}
	if err := json.Unmarshal(body, &pretty); err != nil {
		fmt.Println("⚠️ WARNING: Failed to parse JSON for display:", err)
		fmt.Println("===== RAW POSTED BODY (unformatted) =====")
		fmt.Println(string(body))
		fmt.Println("=========================================")
	} else {
		formatted, _ := json.MarshalIndent(pretty, "", "  ")
		fmt.Println("===== Pretty Posted Body =====")
		fmt.Println(string(formatted))
		fmt.Println("================================")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Agent Callback Handler :: Receives data from Python agent and displays all fields
	//||------------------------------------------------------------------------------------------------||

	var payload agent_interfaces.AgentResponse
	if err := json.Unmarshal(body, &payload); err != nil {
		fmt.Println("⚠️ WARNING: Failed to parse JSON:", err)
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Spit it all out
	//||------------------------------------------------------------------------------------------------||

	success := "FAILED"
	if payload.Success {
		success = "SUCCESS"
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Spit it all out
	//||------------------------------------------------------------------------------------------------||

	fmt.Println("||--------------------------------------------------------------------------------||")
	fmt.Println("|| 📥 Received Agent Callback")
	fmt.Println("||--------------------------------------------------------------------------------||")
	fmt.Printf("|| Status: %s\n", success)
	fmt.Printf("|| Message: %s\n", payload.Message)
	fmt.Println("|| Data:")
	prettyJSON, _ := json.MarshalIndent(payload.Data, "||   ", "  ")
	fmt.Println("||   " + string(prettyJSON))
	fmt.Println("||--------------------------------------------------------------------------------||")
	fmt.Println("|| Request:")
	prettyRequest, _ := json.MarshalIndent(payload.Request, "||   ", "  ")
	fmt.Println("||   " + string(prettyRequest))

	//||------------------------------------------------------------------------------------------------||
	//|| We're good
	//||------------------------------------------------------------------------------------------------||

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"received":true}`))
}
