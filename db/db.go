package db

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //Postgres dialect.
)

var db *gorm.DB

// InitDB initiate connection to database.
func InitDB() (*gorm.DB, error) {
	stringConnection := fmt.Sprintf(
		"host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_DATABASE"),
		os.Getenv("DB_PASSWORD"),
	)

	return gorm.Open("postgres", stringConnection)
}

// Get returning DB instance.
func Get() *gorm.DB {
	if db == nil {
		conn, err := InitDB()
		if err != nil {
			panic(err.Error())
		}

		db = conn
	}

	return db
}
