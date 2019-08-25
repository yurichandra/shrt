package service

import (
	"github.com/go-redis/redis"
	"github.com/yurichandra/shrt/repository"
)

// InitRedisService initialize redis service.
func InitRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
	}
}

// InitURLService initialize url service.
func InitURLService(repo repository.URLRepositoryContract, redis RedisServiceContract) *URLService {
	return &URLService{
		repo:  repo,
		cache: redis,
	}
}
