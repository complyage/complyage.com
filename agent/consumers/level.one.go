package consumers

//||------------------------------------------------------------------------------------------------||
//|| Import
//||------------------------------------------------------------------------------------------------||

import (
	"agent/handlers"
	"agent/publish"
	"encoding/json"
	"errors"
	"fmt"
)

//||------------------------------------------------------------------------------------------------||
//|| Level 1
//||------------------------------------------------------------------------------------------------||

func Level1Router(msg []byte) error {

	fmt.Println("Level1Router: Received message for Level 1 processing")

	//||------------------------------------------------------------------------------------------------||
	//|| Unmarshal
	//||------------------------------------------------------------------------------------------------||

	var av publish.AgentVerification
	if err := json.Unmarshal(msg, &av); err != nil {
		return fmt.Errorf("failed to unmarshal AgentVerification: %w", err)
	}

	fmt.Printf("Parsed AgentVerification: %+v\n", av)

	//||------------------------------------------------------------------------------------------------||
	//|| Route by Process
	//||------------------------------------------------------------------------------------------------||

	switch av.Process {
	case publish.ProcessVerifyID:
		return handlers.HandleVerifyID(av)
	case publish.ProcessFacialAge:
		return handlers.HandleFacialAge(av)
	default:
		return errors.New("unknown process type: " + string(av.Process))
	}
}
