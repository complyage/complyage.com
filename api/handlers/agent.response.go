package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"base/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| AgentResponse - Structure matching AI agent's postback payload
//||------------------------------------------------------------------------------------------------||

type AgentResponse struct {
	Status string                 `json:"status"` // "success" or "error"
	Action string                 `json:"action"` // e.g., NSFW, ID, FACIAL, AGE
	Data   map[string]interface{} `json:"data"`   // contains "params" and "result"
}

//||------------------------------------------------------------------------------------------------||
//|| AgentResponseHandler - Processes AI agent postback with shared key verification
//||------------------------------------------------------------------------------------------------||

func AgentResponseHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Check Shared Key
	//||------------------------------------------------------------------------------------------------||

	sharedKey := os.Getenv("AGENT_SHARED_KEY")
	authHeader := r.Header.Get("Authorization")

	if sharedKey == "" {
		responses.Error(w, http.StatusInternalServerError, "Shared key not configured")
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") || strings.TrimPrefix(authHeader, "Bearer ") != sharedKey {
		responses.Error(w, http.StatusUnauthorized, "Invalid or missing shared key")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate JSON
	//||------------------------------------------------------------------------------------------------||

	var agentResp AgentResponse
	if err := json.NewDecoder(r.Body).Decode(&agentResp); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON payload")
		return
	}

	fmt.Println(agentResp)

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Payload
	//||------------------------------------------------------------------------------------------------||

	if agentResp.Action == "" || agentResp.Status == "" {
		responses.Error(w, http.StatusBadRequest, "Missing required fields: action or status")
		return
	}
	params, _ := agentResp.Data["params"].(map[string]interface{})
	result, _ := agentResp.Data["result"].(map[string]interface{})

	//||------------------------------------------------------------------------------------------------||
	//|| Process Actions
	//||------------------------------------------------------------------------------------------------||

	switch agentResp.Action {
	case "NSFW":
		processNSFWResult(params, result)

	case "FACIAL":
		processFacialMatch(params, result)

	case "ID":
		processIdentification(params, result)

	case "AGE":
		processAgeEstimation(params, result)

	default:
		// CONTENT or other actions
		processGenericAction(agentResp.Action, params, result)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Response
	//||------------------------------------------------------------------------------------------------||

	resp := map[string]interface{}{
		"message": "Agent response processed",
		"action":  agentResp.Action,
		"status":  agentResp.Status,
		"time":    time.Now(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, resp)
}

// ||------------------------------------------------------------------------------------------------||
// || Stub Processing Functions (Implement as needed)
// ||------------------------------------------------------------------------------------------------||

func processNSFWResult(params map[string]interface{}, result map[string]interface{}) {
	fmt.Println(params)
	fmt.Println(result)
}

func processFacialMatch(params map[string]interface{}, result map[string]interface{}) {
	// TODO: save facial match results
}

func processIdentification(params map[string]interface{}, result map[string]interface{}) {
	// TODO: handle ID verification results
}

func processAgeEstimation(params map[string]interface{}, result map[string]interface{}) {
	// TODO: handle age estimation results
}

func processGenericAction(action string, params map[string]interface{}, result map[string]interface{}) {
	// TODO: default handler for other actions
}
