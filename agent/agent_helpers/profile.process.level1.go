package agent_helpers

import (
	"agent/agent_interfaces"
	"agent/prompts"
	"base/db"
	"base/helpers"
	"context"
	"encoding/json"
	"os"
	"time"
)

//||------------------------------------------------------------------------------------------------||
//|| AgentProfileProcess
//||------------------------------------------------------------------------------------------------||

func AgentProfileProcessLevelOne(session agent_interfaces.AgentRequestProfile) error {

	//||------------------------------------------------------------------------------------------------||
	//|| Start Request
	//||------------------------------------------------------------------------------------------------||

	request := agent_interfaces.AgentRequest{
		Identity:  session.Identity,
		Params:    session.Params,
		CallBack:  session.CallBack,
		Timestamp: helpers.UniversalNow(),
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serialize Session
	//||------------------------------------------------------------------------------------------------||

	switch session.Process.Step {
	case 0:
		request.Model = "ocr"
		request.Action = "extract"
		request.Prompt = ""
		request.Media = []agent_interfaces.AgentMedia{session.Front}
		request.CallBack = os.Getenv("VITE_COMPLYAGE_AGENT_URL") + "/v1/callback/profile/?step=0"
		break
	case 1:
		request.Model = "face"
		request.Action = "detect"
		request.Prompt = ""
		request.Media = []agent_interfaces.AgentMedia{session.Front}
		request.CallBack = os.Getenv("VITE_COMPLYAGE_AGENT_URL") + "/v1/callback/profile/?step=1"
		break
	case 2:
		request.Model = "vision"
		request.Action = "content"
		request.Prompt = prompts.PromptLevelOneStepTwo(session.Process.RawText)
		request.Media = []agent_interfaces.AgentMedia{}
		request.CallBack = os.Getenv("VITE_COMPLYAGE_AGENT_URL") + "/v1/callback/profile/?step=2"
		break
	case 3:
		request.Model = "face"
		request.Action = "match"
		request.Prompt = ""
		request.Media = []agent_interfaces.AgentMedia{session.Front, session.Profile}
		request.CallBack = os.Getenv("VITE_COMPLYAGE_AGENT_URL") + "/v1/callback/profile/?step=3"
		break
	case 4:
		if session.Process.Error != "" {
			return respondCallback(false, session.Process.Error, session)
		}
		if session.Process.FaceMatch == false {
			return respondCallback(false, "LEVEL1_FACEMISMATCH", session)
		}
		if session.Process.RawText == "" {
			return respondCallback(false, "LEVEL1_NOTEXT", session)
		}
		if session.Process.IDVerified == false {
			return respondCallback(false, "LEVEL1_ADDRESS_PARSE", session)
		}
		return respondCallback(true, "LEVEL1_SUCCESS", session)
	default:
		return respondCallback(false, "LEVEL1_INVALID_STEP", session)
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Serialize Session
	//||------------------------------------------------------------------------------------------------||

	sessionData, err := json.Marshal(session)
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Save to Redis
	//||------------------------------------------------------------------------------------------------||

	redisKey := "agent:session:" + session.Identity
	err = db.Redis.Set(context.Background(), redisKey, sessionData, 60*time.Minute).Err()
	if err != nil {
		return err
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Channel
	//||------------------------------------------------------------------------------------------------||

	channel := os.Getenv("RABBITMQ_CHANNEL")
	if session.Level >= 3 {
		channel = os.Getenv("RABBITMQ_HUMAN_CHANNEL")
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Push to Queue
	//||------------------------------------------------------------------------------------------------||

	pubErr := PublishAgentRequest(db.MQChan, channel, request)
	if pubErr != nil {
		return pubErr
	}

	//||------------------------------------------------------------------------------------------------||
	//|| Success
	//||------------------------------------------------------------------------------------------------||

	return nil

}

//||------------------------------------------------------------------------------------------------||
//|| RespondCallback
//||------------------------------------------------------------------------------------------------||

func respondCallback(success bool, message string, session agent_interfaces.AgentRequestProfile) error {
	return nil
}
