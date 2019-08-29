package service

import (
	"time"

	"github.com/yurichandra/shrt/model"
	"github.com/yurichandra/shrt/repository"
)

// ShortenerService represent service layer of URL model.
type ShortenerService struct {
	repo     repository.URLRepositoryContract
	cache    RedisServiceContract
	userRepo repository.UserRepositoryContract
}

// Find finds an url by shorturl.
func (s *ShortenerService) Find(shortURL string) (model.URL, error) {
	return model.URL{}, nil
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
			ExpiredDate: now.AddDate(0, 0, 7*1),
			UserID:      0,
		}

		s.cache.Map(key, &url)
		s.repo.Create(&url)

		return url, nil
	}

	apiKey := data["apiKey"]
	user := s.userRepo.FindByKey(apiKey)

	url := model.URL{
		OriginalURL: originalURL,
		ShortURL:    key,
		ExpiredDate: now.AddDate(0, 0, 7*1),
		UserID:      user.ID,
	}

	s.cache.Map(key, &url)
	s.repo.Create(&url)

	return url, nil
}
