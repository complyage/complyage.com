package main

import (
	"agent/consumers"
	"agent/prompts"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/ralphferrara/aria/app"
)

//||------------------------------------------------------------------------------------------------||
//|| Main
//||------------------------------------------------------------------------------------------------||

func main() {
	//||------------------------------------------------------------------------------------------------||
	//|| Load Env
	//||------------------------------------------------------------------------------------------------||
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println("No .env file found, continuing...")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Starting switch-over to aria
	//||------------------------------------------------------------------------------------------------||
	if err := app.Init("../config.json"); err != nil {
		app.Log.Error("main", "Startup failed: %v", err)
		os.Exit(1)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Start Consumer
	//||------------------------------------------------------------------------------------------------||
	rabbit := app.QueueRabbit["agent"]
	if rabbit == nil {
		panic("RabbitMQ instance 'agent' not found")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Consume Level 1
	//||------------------------------------------------------------------------------------------------||
	errL1 := rabbit.ConsumeQueue("AgentLevelOne", func(msg []byte) {
		if err := consumers.Level1Router(msg); err != nil {
			app.Log.Error("main", "Level1Router error: %v", err)
		}
	})
	if errL1 != nil {
		app.Log.Error("main", "Failed to start Level 1 consumer: %v", errL1)
		os.Exit(1)
	}
	app.Log.Info("main", "Listening for Level 1 messages on 'AgentLevelOne' queue")
	//||------------------------------------------------------------------------------------------------||
	//|| Consume Level 1
	//||------------------------------------------------------------------------------------------------||
	errL2 := rabbit.ConsumeQueue("AgentLevelTwo", func(msg []byte) {
		if err := consumers.Level2Router(msg); err != nil {
			app.Log.Error("main", "Level2Router error: %v", err)
		}
	})
	if errL2 != nil {
		app.Log.Error("main", "Failed to start Level 2 consumer: %v", errL1)
		os.Exit(1)
	}
	app.Log.Info("main", "Listening for Level 2 messages on 'AgentLevelTwo' queue")
	//||------------------------------------------------------------------------------------------------||
	//|| Started
	//||------------------------------------------------------------------------------------------------||
	app.Log.Info("main", "Agent has started")
	//||------------------------------------------------------------------------------------------------||
	//|| Template
	//||------------------------------------------------------------------------------------------------||
	prompts.RegisterPrompts()
	//||------------------------------------------------------------------------------------------------||
	//|| Keep Running
	//||------------------------------------------------------------------------------------------------||
	select {}
}
