package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/yurichandra/shrt/model"
	"github.com/yurichandra/shrt/repository"
)

// ShortenerService represent service layer of URL model.
type ShortenerService struct {
	redisService RedisServiceContract
	urlRepo      repository.URLRepositoryContract
	userRepo     repository.UserRepositoryContract
	url          model.URL
}

// Find finds an url by key.
func (s *ShortenerService) Find(key string) (model.URL, error) {
	mappedURL, err := s.redisService.Find(key)
	if err != nil {
		return model.URL{}, errors.New("Shortened URL is not found")
	}

	urlData := []byte(mappedURL)
	err = json.Unmarshal(urlData, &s.url)
	if err != nil {
		return model.URL{}, err
	}

	return s.url, nil
}

// Shorten shortens original URL.
func (s *ShortenerService) Shorten(
	originalURL string,
	apiKey string,
	now time.Time,
	auth bool,
) (model.URL, error) {
	key, err := s.redisService.Generate()
	if err != nil {
		return model.URL{}, err
	}

	if !auth {
		url := model.URL{
			OriginalURL: originalURL,
			Keys:        key,
			ExpiredDate: now.Add(time.Hour * 24),
			UserID:      0,
		}

		s.redisService.Map(key, &url)
		s.urlRepo.Create(&url)

		return url, nil
	}

	return s.ShortenWithAuth(originalURL, key, time.Now(), apiKey)
}

// ShortenWithAuth doing shorten while authentication is true.
func (s *ShortenerService) ShortenWithAuth(
	originalURL string,
	key string,
	now time.Time,
	apiKey string,
) (model.URL, error) {
	user := s.userRepo.FindByKey(apiKey)
	if user.ID == 0 {
		return model.URL{}, errors.New("Your api_key is invalid")
	}

	url := s.urlRepo.FindBy(originalURL, user.ID)
	if url.ID != 0 {
		return url, nil
	}

	url = model.URL{
		OriginalURL: originalURL,
		Keys:        key,
		ExpiredDate: now.Add(time.Hour * 24),
		UserID:      user.ID,
	}

	s.redisService.Map(key, &url)
	s.urlRepo.Create(&url)

	return url, nil
}
