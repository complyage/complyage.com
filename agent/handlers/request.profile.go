package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/agent_helpers"
	"agent/agent_interfaces"
	"base/loaders"
	"base/responses"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: Accept AgentRequest and Queue It
//||------------------------------------------------------------------------------------------------||

func RequestProfileHandler(w http.ResponseWriter, r *http.Request) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse JSON Body
	//||------------------------------------------------------------------------------------------------||

	var req agent_interfaces.APIAgentRequestProfile
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Pretty Print Request JSON to Console
	//||------------------------------------------------------------------------------------------------||

	logReq := req

	//||------------------------------------------------------------------------------------------------||
	//|| Clear up logs
	//||------------------------------------------------------------------------------------------------||

	logReq.Front.Base64 = ""
	logReq.Front.Blob = ""
	logReq.Back.Base64 = ""
	logReq.Back.Blob = ""
	logReq.Profile.Base64 = ""
	logReq.Profile.Blob = ""

	//||------------------------------------------------------------------------------------------------||
	//|| Pretty Bird
	//||------------------------------------------------------------------------------------------------||

	if prettyJSON, err := json.MarshalIndent(logReq, "", "  "); err == nil {
		fmt.Println("===== Incoming Request JSON =====")
		fmt.Println(string(prettyJSON))
		fmt.Println("=================================")
	} else {
		fmt.Println("❌ Failed to marshal request for logging:", err)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Get Site By API Key
	//||------------------------------------------------------------------------------------------------||

	site := loaders.GetSiteByAgentKey(req.AgentKey)
	if site == nil {
		responses.Error(w, http.StatusBadRequest, "Invalid agentKey")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Have Required Media?
	//||------------------------------------------------------------------------------------------------||

	if agent_helpers.ValidateMedia(req.Front) != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid identification front.")
		return
	}

	if agent_helpers.ValidateMedia(req.Back) != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid identification back.")
		return
	}

	if agent_helpers.ValidateMedia(req.Profile) != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid profile photo.")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Level
	//||------------------------------------------------------------------------------------------------||

	if req.Level < 0 || req.Level > 3 {
		responses.Error(w, http.StatusBadRequest, "Invalid level. Must be between 0 and 3.")
		return
	}

	if req.MaxLevel < 0 || req.MaxLevel > 3 {
		responses.Error(w, http.StatusBadRequest, "Invalid max_level. Must be between 0 and 3.")
		return
	}

	if req.Level > req.MaxLevel {
		responses.Error(w, http.StatusBadRequest, "Level cannot be greater than max_level.")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Callback
	//||------------------------------------------------------------------------------------------------||

	err := agent_helpers.ValidateCallbackURL(req.CallBack)
	if err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid callback URL: "+err.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Fields
	//||------------------------------------------------------------------------------------------------||

	sessionId := uuid.New().String()
	session := agent_interfaces.AgentRequestProfile{
		SiteId:    site.IDSite,
		Identity:  sessionId,
		Params:    req.Params,
		Front:     req.Front,
		Back:      req.Back,
		Profile:   req.Profile,
		Level:     req.Level,
		MaxLevel:  req.MaxLevel,
		CallBack:  req.CallBack,
		Timestamp: time.Now(),
		Elapsed:   0,
		Process: agent_interfaces.IdentificationProcess{
			Step:       0,
			RawText:    "",
			FaceMatch:  false,
			IDVerified: false,
		},
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Respond with Success
	//||------------------------------------------------------------------------------------------------||

	procErr := agent_helpers.AgentProfileProcessLevelOne(session)
	if procErr != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to process profile request: "+procErr.Error())
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Respond with Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"queued":   true,
		"identity": sessionId,
	})
}
