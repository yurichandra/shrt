package repository

import "github.com/jinzhu/gorm"

// NewURLRepository return URLRepository.
func NewURLRepository(db *gorm.DB) *URLRepository {
	return &URLRepository{
		db: db,
	}
}
