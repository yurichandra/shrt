package repository

import "github.com/jinzhu/gorm"

// NewURLRepository return URLRepository.
func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}

// NewUserRepository return UserRepository.
func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{
		db: db,
	}
}
