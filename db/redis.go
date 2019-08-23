package db

import (
	"fmt"
	"os"

	"github.com/go-redis/redis"
)

var conn *redis.Client

// InitRedis initialize redis connection.
func InitRedis() *redis.Client {
	address := fmt.Sprintf(
		"%s:%s",
		os.Getenv("REDIS_HOST"),
		os.Getenv("REDIS_PORT"),
	)

	conn := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: "",
		DB:       0,
	})

	return conn
}

// GetRedis return redis client.
func GetRedis() *redis.Client {
	if conn == nil {
		conn = InitRedis()
	}

	return conn
}
