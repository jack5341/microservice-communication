package db

import (
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"os"
)

func Conn() *redis.Client {
	url := os.Getenv("REDIS_URL")
	pass := os.Getenv("REDIS_PASS")

	if url == "" {
		log.Fatal("REDIS_URL is not set")
	}

	client := redis.NewClient(&redis.Options{
		Addr:     url,
		Password: pass,
	})

	_, err := client.Ping().Result()

	if err != nil {
		log.Fatal(err)
	}

	log.Info("successfully connected to db")
	return client
}
