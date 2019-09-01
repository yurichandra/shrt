package service

import (
	"github.com/go-redis/redis"
	"github.com/yurichandra/shrt/repository"
)

// NewRedisService returns redis service instance.
func NewRedisService(client *redis.Client) *RedisService {
	return &RedisService{
		client: client,
	}
}

// NewShortenerService returns url service instance.
func NewShortenerService(repo repository.URLRepositoryContract, userRepo repository.UserRepositoryContract, redis RedisServiceContract) *ShortenerService {
	return &ShortenerService{
		redisService: redis,
		urlRepo:      repo,
		userRepo:     userRepo,
	}
}

// NewAuthService returns auth service instance.
func NewAuthService(repo repository.UserRepositoryContract) *AuthService {
	return &AuthService{
		repo: repo,
	}
}
