package handlers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/complyage/base/keeper"
	"github.com/ralphferrara/aria/responses"
)

//||------------------------------------------------------------------------------------------------||
//|| Request
//||------------------------------------------------------------------------------------------------||

type ManualRequest struct {
	IP        string `json:"ip"`
	SessionId string `json:"session_id"`
	ClientId  string `json:"client_id"`
}

//||------------------------------------------------------------------------------------------------||
//|| Response
//||------------------------------------------------------------------------------------------------||

func ManualHandler(w http.ResponseWriter, r *http.Request) {
	//||------------------------------------------------------------------------------------------------||
	//|| Decode JSON Body
	//||------------------------------------------------------------------------------------------------||
	var req ManualRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Initiate the Zone
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("Manual check for session:", req.SessionId, "and IP:", req.IP, "and ClientID:", req.ClientId)
	keep, err := keeper.Manual(r, req.IP, req.SessionId, req.ClientId)
	if err != nil {
		fmt.Println("Error loading manual keeper:", err)
		responses.Error(w, http.StatusInternalServerError, "Error loading manual keeper")
		return
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Write the Cookie
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, keep)
}
