package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"net/http"
)

//||------------------------------------------------------------------------------------------------||
//|| Handler: Accept AgentRequest and Queue It
//||------------------------------------------------------------------------------------------------||

func AgentRequestHandler(w http.ResponseWriter, r *http.Request) {

	// //||------------------------------------------------------------------------------------------------||
	// //|| Parse JSON Body
	// //||------------------------------------------------------------------------------------------------||

	// var req agent.AgentRequest
	// if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
	// 	responses.Error(w, http.StatusBadRequest, "Invalid JSON body")
	// 	return
	// }
	// req.Timestamp = time.Now().UTC().Truncate(time.Second)

	// //||------------------------------------------------------------------------------------------------||
	// //|| Validate Action
	// //||------------------------------------------------------------------------------------------------||

	// if !agent.IsValidAction(req.Action) {
	// 	responses.Error(w, http.StatusBadRequest, "Invalid action: "+req.Action)
	// 	return
	// }

	// //||------------------------------------------------------------------------------------------------||
	// //|| Publish to RabbitMQ (using global agent.MQChan)
	// //||------------------------------------------------------------------------------------------------||

	// channelName := os.Getenv("RABBITMQ_CHANNEL")
	// if err := agent.PublishAgentRequest(agent.MQChan, channelName, req); err != nil {
	// 	responses.Error(w, http.StatusInternalServerError, "Failed to enqueue AgentRequest")
	// 	return
	// }

	// //||------------------------------------------------------------------------------------------------||
	// //|| Respond with Success
	// //||------------------------------------------------------------------------------------------------||

	// responses.Success(w, http.StatusOK, map[string]any{
	// 	"queued":   true,
	// 	"action":   req.Action,
	// 	"callback": req.CallBack,
	// })
}
