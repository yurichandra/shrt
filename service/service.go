package service

import "github.com/go-redis/redis"

// InitRedisService initialize redis service.
func InitRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
	}
}
