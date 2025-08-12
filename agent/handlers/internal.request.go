package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/agent_helpers"
	"agent/agent_interfaces"
	"base/db"
	"base/helpers"
	"base/responses"
	"encoding/json"
	"net/http"
	"os"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: Accept AgentRequest and Queue It
//||------------------------------------------------------------------------------------------------||

func InternalRequestHandler(w http.ResponseWriter, r *http.Request, model string, action string) {

	//||------------------------------------------------------------------------------------------------||
	//|| Parse JSON Body
	//||------------------------------------------------------------------------------------------------||

	var req agent_interfaces.AgentRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		responses.Error(w, http.StatusBadRequest, "Invalid JSON body")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Model
	//||------------------------------------------------------------------------------------------------||

	if !agent_helpers.IsValidModel(model) {
		responses.Error(w, http.StatusBadRequest, "Invalid model: "+model)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Validate Action
	//||------------------------------------------------------------------------------------------------||

	if !agent_helpers.IsValidAction(model, action) {
		responses.Error(w, http.StatusBadRequest, "Invalid action: "+model+"-"+action)
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Have Required Media?
	//||------------------------------------------------------------------------------------------------||

	if len(req.Media) > agent_helpers.RequiresImage(model, action) {
		responses.Error(w, http.StatusBadRequest, "Invalid media count for action: "+model+"-"+action+". Requires "+string(agent_helpers.RequiresImage(model, action))+" media items.")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Add Fields
	//||------------------------------------------------------------------------------------------------||

	req.Action = action
	req.Model = model
	req.Timestamp = helpers.UniversalNow()

	//||------------------------------------------------------------------------------------------------||
	//|| Publish to RabbitMQ (using global agent.MQChan)
	//||------------------------------------------------------------------------------------------------||

	channelName := os.Getenv("RABBITMQ_CHANNEL")
	if err := agent_helpers.PublishAgentRequest(db.MQChan, channelName, req); err != nil {
		responses.Error(w, http.StatusInternalServerError, "Failed to enqueue AgentRequest")
		return
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Respond with Success
	//||------------------------------------------------------------------------------------------------||

	responses.Success(w, http.StatusOK, map[string]any{
		"queued":   true,
		"action":   req.Action,
		"callback": req.CallBack,
	})
}
