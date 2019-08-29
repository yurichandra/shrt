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
	Find(key string) (string, error)
}

// ShortenerServiceContract represent contract of
// Shorten service.
type ShortenerServiceContract interface {
	Find(shortURL string) (model.URL, error)
	Shorten(data map[string]string, auth bool) (model.URL, error)
}

// AuthServiceContract represent contract of
// AuthService.
type AuthServiceContract interface {
	Authenticate(email string, password string) (model.User, error)
	Authorize(email string, password string) (model.User, error)
}
