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

// InitShortenerService initialize url service.
func InitShortenerService(repo repository.URLRepositoryContract, userRepo repository.UserRepositoryContract, redis RedisServiceContract) *ShortenerService {
	return &ShortenerService{
		redisService: redis,
		urlRepo:      repo,
		userRepo:     userRepo,
	}
}

// InitAuthService initialize auth service.
func InitAuthService(repo repository.UserRepositoryContract) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
