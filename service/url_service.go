package service

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"

	"github.com/yurichandra/shrt/model"
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

// Get return all urls available.
func (srv *URLService) Get() []model.URL {
	urls := srv.repo.Get()

	return urls
}

// Find return single url.
func (srv *URLService) Find(id uint) model.URL {
	url := srv.repo.Find(id)

	return url
}

// Create will store original url into cache or database and return it.
func (srv *URLService) Create(originalURL string) (model.URL, error) {
	proceededURL := srv.processURL(originalURL)
	found, err := srv.cache.HGet(keyCache, proceededURL)
	if len(found) > 0 {
		var urlFound model.URL
		json.Unmarshal([]byte(found), &urlFound)

		return urlFound, nil
	}

	url := model.URL{
		OriginalURL: originalURL,
		ShortURL:    proceededURL,
	}

	err = srv.repo.Create(&url)
	if err != nil {
		return model.URL{}, err
	}

	val, err := json.Marshal(url)

	err = srv.cache.HSet(keyCache, proceededURL, string(val))
	if err != nil {
		return model.URL{}, err
	}

	return url, nil
}

// ProcessURL will process given URL from params and return it as hash.
func (srv *URLService) processURL(s string) string {
	hash := sha1.New()
	hash.Write([]byte(s))

	hashURL := hex.EncodeToString(hash.Sum(nil))

	return hashURL[:maxString]
}
