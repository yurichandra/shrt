package service

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/yurichandra/shrt/model"
	"github.com/yurichandra/shrt/repository"
)

var url model.URL

// ShortenerService represent service layer of URL model.
type ShortenerService struct {
	repo     repository.URLRepositoryContract
	cache    RedisServiceContract
	userRepo repository.UserRepositoryContract
}

// Find finds an url by shorturl.
func (s *ShortenerService) Find(shortURL string) (model.URL, error) {
	mappedURL, err := s.cache.Find(shortURL)
	if err != nil {
		return model.URL{}, errors.New("Shortened URL is not found")
	}

	urlData := []byte(mappedURL)
	err = json.Unmarshal(urlData, &url)
	if err != nil {
		return model.URL{}, err
	}

	return url, nil
}

// Shorten shortens original URL.
func (s *ShortenerService) Shorten(data map[string]string, auth bool) (model.URL, error) {
	originalURL := data["originalURL"]
	now := time.Now()
	key, err := s.cache.Generate()
	if err != nil {
		return model.URL{}, err
	}

	if !auth {
		url := model.URL{
			OriginalURL: originalURL,
			ShortURL:    key,
			ExpiredDate: now.Add(7 * 24),
			UserID:      0,
		}

		s.cache.Map(key, &url)
		s.repo.Create(&url)

		return url, nil
	}

	apiKey := data["apiKey"]

	return s.ShortenWithAuth(originalURL, apiKey)
}

// ShortenWithAuth doing shorten while authentication is true.
func (s *ShortenerService) ShortenWithAuth(originalURL string, apiKey string) (model.URL, error) {
	user := s.userRepo.FindByKey(apiKey)
	if user.ID == 0 {
		return model.URL{}, errors.New("Your api_key is invalid")
	}

	url := s.repo.FindBy(originalURL, user.ID)
	if url.ID != 0 {
		return url, nil
	}

	key, _ := s.cache.Generate()
	url = model.URL{
		OriginalURL: originalURL,
		ShortURL:    key,
		ExpiredDate: time.Now().Add(7 * 24),
		UserID:      user.ID,
	}

	s.cache.Map(key, &url)
	s.repo.Create(&url)

	return url, nil
}
