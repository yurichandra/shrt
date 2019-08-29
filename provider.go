package main

import (
	"log"

	"github.com/go-redis/redis"
	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/service"
)

var (
	urlService   *service.URLService
	redisService *service.RedisService
	redisClient  *redis.Client
)

func boot() {
	var err error

	// dbConn := db.Get()
	redisClient = db.GetRedis()

	redisService = service.InitRedisService(redisClient)

	if err != nil {
		log.Fatal(err)
	}
}
