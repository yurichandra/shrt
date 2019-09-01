package service

import (
	"time"

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
	Find(key string) (model.URL, error)
	Shorten(originalURL string, apiKey string, now time.Time, auth bool) (model.URL, error)
	ShortenWithAuth(originalURL string, key string, now time.Time, apiKey string) (model.URL, error)
}

// AuthServiceContract represent contract of
// AuthService.
type AuthServiceContract interface {
	Authenticate(email string, password string) (model.User, error)
	Authorize(email string, password string) (model.User, error)
}
