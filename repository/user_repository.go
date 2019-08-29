package repository

import (
	"github.com/jinzhu/gorm"
	"github.com/yurichandra/shrt/model"
)

// UserRepository represent repository of User.
type UserRepository struct {
	db *gorm.DB
}

// Find return a user.
func (repo *UserRepository) Find(id uint) model.User {
	user := model.User{}
	repo.db.Where("id = ?", id).First(&user)

	return user
}

// FindByKey return a user find by key.
func (repo *UserRepository) FindByKey(key string) model.User {
	user := model.User{}
	repo.db.Where("key = ?", key).First(&user)

	return user
}

// Create creates a new user and return error if occured.
func (repo *UserRepository) Create(data *model.User) error {
	return repo.db.Create(data).Error
}
