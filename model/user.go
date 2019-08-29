package model

import (
	"time"
)

// User represent model of user.
type User struct {
	ID        uint
	Email     string
	Password  string
	Key       string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	URL       []URL `gorm:"foreignKey:ID"`
}
