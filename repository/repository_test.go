package repository

import (
	"fmt"
	"testing"

	"github.com/jaswdr/faker"
	"github.com/jinzhu/gorm"
	"github.com/joho/godotenv"
	"github.com/yurichandra/shrt/db"

	_ "github.com/joho/godotenv/autoload"
)

var (
	_testFaker faker.Faker
	_testDB    *gorm.DB
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		fmt.Println("No env file specified")
		return
	}

	_testFaker = faker.New()
	_testDB = db.Get()

	m.Run()
}
