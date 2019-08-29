package service

import (
	"github.com/yurichandra/shrt/repository"
)

// URLService represent service layer of URL model.
type URLService struct {
	repo  repository.URLRepositoryContract
	cache RedisServiceContract
}
