package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/repository"
	"github.com/yurichandra/shrt/service"
)

var (
	shortenerService *service.ShortenerService
	redisService     *service.RedisService
	authService      *service.AuthService
	redisClient      *redis.Client

	userRepo repository.UserRepositoryContract
)

func boot() {
	var err error

	dbConn := db.Get()
	redisClient = db.GetRedis()

	userRepo := repository.NewUserRepository(dbConn)
	urlRepo := repository.NewURLRepository(dbConn)

	redisService = service.NewRedisService(redisClient)
	authService = service.NewAuthService(userRepo)
	shortenerService = service.NewShortenerService(urlRepo, userRepo, redisService)

	if err != nil {
		log.Fatal(err)
	}
}
