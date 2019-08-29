package service

import (
	"github.com/yurichandra/shrt/repository"
)

// ShortenerService represent service layer of URL model.
type ShortenerService struct {
	repo  repository.URLRepositoryContract
	cache RedisServiceContract
}
