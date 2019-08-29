package service

import (
	"github.com/yurichandra/shrt/model"
)

// RedisServiceContract represent contract of
// redis service.
type RedisServiceContract interface {
	Init() error
	Generate() (string, error)
	Map(key string, url *model.URL) error
}

// URLServiceContract represent contract of
// URL service.
type URLServiceContract interface {
	Get() []model.URL
	Find(id uint) model.URL
	Create(originalURL string) (model.URL, error)
}

// AuthServiceContract represent contract of
// AuthService.
type AuthServiceContract interface {
	Authenticate(email string, password string) (model.URL, error)
	Authorize(email string, password string) (model.URL, error)
}
