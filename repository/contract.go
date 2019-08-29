package repository

import (
	"github.com/yurichandra/shrt/model"
)

// URLRepositoryContract represent contract
// of URLRepository.
type URLRepositoryContract interface {
	Get() []model.URL
	Find(id uint) model.URL
	Create(data *model.URL) error
	Update(data *model.URL) error
	Delete(id uint) error
}

// UserRepositoryContract represent contract
// of UserRepository.
type UserRepositoryContract interface {
	Find(id uint) model.User
	FindByKey(key string) model.User
	FindByEmail(email string) model.User
	Create(data *model.User) error
}
