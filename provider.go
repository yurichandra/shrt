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

	redisService = service.InitRedisService(redisClient)
	authService = service.InitAuthService(userRepo)

	if err != nil {
		log.Fatal(err)
	}
}
