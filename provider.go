package main

import (
	"log"

	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/repository"
	"github.com/yurichandra/shrt/service"
)

var (
	urlService   *service.URLService
	redisService *service.RedisService
)

func boot() {
	var err error

	dbConn := db.Get()
	redisConn := db.GetRedis()

	redisService = service.InitRedisService(redisConn)

	urlService = service.InitURLService(
		repository.NewURLRepository(dbConn),
		redisService,
	)

	if err != nil {
		log.Fatal(err)
	}
}
