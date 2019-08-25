package service

import (
	"github.com/yurichandra/shrt/model"
)

// RedisServiceContract represent contract of
// redis service.
type RedisServiceContract interface {
	HGet(hash string, field string) (string, error)
	HSet(hash string, field string, value string) error
}

// URLServiceContract represent contract of
// URL service.
type URLServiceContract interface {
	Get() []model.URL
	Find(id uint) model.URL
	Create(originalURL string) (model.URL, error)
}
