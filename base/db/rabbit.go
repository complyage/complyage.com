package db

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	MQConn *amqp.Connection
	MQChan *amqp.Channel
)

//||------------------------------------------------------------------------------------------------||
//|| Connect to RabbitMQ
//||------------------------------------------------------------------------------------------------||

func ConnectMQ() {
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to RabbitMQ
	//||------------------------------------------------------------------------------------------------||
	dotErr := godotenv.Load("../.env")
	if dotErr != nil {
		fmt.Println("No .env file found, continuing...")
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Env
	//||------------------------------------------------------------------------------------------------||
	rabbitHost := os.Getenv("RABBITMQ_HOST")
	rabbitPort := os.Getenv("RABBITMQ_PORT")
	rabbitUser := os.Getenv("RABBITMQ_USER")
	rabbitPass := os.Getenv("RABBITMQ_PASSWORD")
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/", rabbitUser, rabbitPass, rabbitHost, rabbitPort)
	fmt.Println("Connecting to RabbitMQ at", connStr)
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to RabbitMQ
	//||------------------------------------------------------------------------------------------------||
	var err error
	MQConn, err = amqp.Dial(connStr)
	if err != nil {
		log.Fatal("❌ Failed to connect to RabbitMQ:", err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Connect to Channel
	//||------------------------------------------------------------------------------------------------||
	MQChan, err = MQConn.Channel()
	if err != nil {
		log.Fatal("❌ Failed to open RabbitMQ channel:", err)
	}
	//||------------------------------------------------------------------------------------------------||
	//|| Connected
	//||------------------------------------------------------------------------------------------------||
	fmt.Println("✅ Connected to RabbitMQ!")
}

//||------------------------------------------------------------------------------------------------||
//|| Close RabbitMQ
//||------------------------------------------------------------------------------------------------||

func CloseMQ() {
	if MQChan != nil {
		MQChan.Close()
	}
	if MQConn != nil {
		MQConn.Close()
	}
	fmt.Println("✅ RabbitMQ connection closed.")
}
