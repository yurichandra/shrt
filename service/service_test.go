package service

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"

	"github.com/go-redis/redis"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"

	"github.com/yurichandra/shrt/db"
)

var (
	_mockRedis *redis.Client
	_mockFaker faker.Faker
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("No .env file specified")

		return
	}

	_mockRedis = db.GetRedis()
	_mockFaker = faker.New()

	m.Run()
}
