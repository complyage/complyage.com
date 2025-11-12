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
	IP        string `json:"ip_address"`
	SessionId string `json:"session_id"`
	ClientId  string `json:"client_id"`
}

//||------------------------------------------------------------------------------------------------||
//|| Manual Response
//||------------------------------------------------------------------------------------------------||

type ManualResponse struct {
	SessionID string `json:"session_id"`
	Enforced  bool   `json:"enforced"`
	Verified  bool   `json:"verified"`
	Age       int    `json:"age"`
	UserID    int64  `json:"user_id"`
	IPAddress string `json:"ip_address"`
	ClientID  string `json:"client_id"`
	Status    string `json:"status"`
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
	//|| Manual
	//||------------------------------------------------------------------------------------------------||
	manualResponse := ManualResponse{
		SessionID: keep.KeeperId,
		Enforced:  keep.Enforced,
		Verified:  keep.Verified,
		Age:       keep.Age,
		UserID:    keep.UserId,
		IPAddress: keep.IPAddress,
		ClientID:  keep.ClientId,
		Status:    keep.Status,
	}
	fmt.Println("Manual Response:", manualResponse)
	//||------------------------------------------------------------------------------------------------||
	//|| Write the Cookie
	//||------------------------------------------------------------------------------------------------||
	responses.Success(w, http.StatusOK, keep)
}
