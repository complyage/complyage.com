package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/redis/go-redis/v9"
)

var (
	Redis *redis.Client
	Ctx   = context.Background()
)

func ConnectRedis() {
	godotenv.Load()
	Redis = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("REDIS_HOST"), os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"), // "" if no password
		DB:       0,                           // use default DB
	})

	// ✅ Test connection
	pong, err := Redis.Ping(Ctx).Result()
	if err != nil {
		log.Fatal("❌ Failed to connect to Redis:", err)
	}
	fmt.Println("✅ Connected to Redis! Ping:", pong)
}
