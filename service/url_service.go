package service

import (
	"github.com/yurichandra/shrt/repository"
)

const (
	keyCache  = "urls"
	maxString = 6
)

// URLService represent service layer of URL model.
type URLService struct {
	repo  repository.URLRepositoryContract
	cache RedisServiceContract
}
