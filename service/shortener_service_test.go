package service

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/yurichandra/shrt/mock"
	"github.com/yurichandra/shrt/model"
)

func TestFindShortenerService(t *testing.T) {
	mockRedisService := &mock.RedisServiceMock{}
	mockURLRepository := &mock.URLRepositoryMock{}
	mockUserRepository := &mock.UserRepositoryMock{}

	key := _mockFaker.Lorem().Word()

	url := model.URL{
		OriginalURL: _mockFaker.Lorem().Word(),
		Keys:        key,
		UserID:      uint(1),
	}

	marshalled, _ := json.Marshal(url)

	mockRedisService.On("Find", key).Return(string(marshalled), nil)

	shortenerService := ShortenerService{
		redisService: mockRedisService,
		urlRepo:      mockURLRepository,
		userRepo:     mockUserRepository,
	}

	expected, err := shortenerService.Find(key)
	if err != nil {
		t.Errorf("TestShortenerService find was failed")
	}

	if expected.Keys != url.Keys {
		t.Errorf("Expected and actual data is not equal")
	}
}

func TestShortenShortenerService(t *testing.T) {
	mockRedisService := &mock.RedisServiceMock{}
	mockURLRepository := &mock.URLRepositoryMock{}
	mockUserRepository := &mock.UserRepositoryMock{}

	key := _mockFaker.Lorem().Word()
	now := time.Now()

	url := model.URL{
		OriginalURL: _mockFaker.Lorem().Word(),
		Keys:        key,
		ExpiredDate: now.Add(time.Hour * 24),
		UserID:      uint(0),
	}

	mockRedisService.On("Generate").Return(key, nil)
	mockRedisService.On("Map", key, &url).Return(nil)
	mockURLRepository.On("Create", &url).Return(nil)

	shortenerService := ShortenerService{
		redisService: mockRedisService,
		urlRepo:      mockURLRepository,
		userRepo:     mockUserRepository,
	}

	actual, err := shortenerService.Shorten(url.OriginalURL, "", now, false)
	if err != nil {
		t.Errorf("TestShortenerService shorten was failed")
	}

	if actual.Keys != url.Keys {
		t.Errorf("Shorten new url was failed")
	}
}

func TestShorthenWithAuthShortenerService(t *testing.T) {
	mockRedisService := &mock.RedisServiceMock{}
	mockURLRepository := &mock.URLRepositoryMock{}
	mockUserRepository := &mock.UserRepositoryMock{}

	key := _mockFaker.Lorem().Word()
	apiKey := _mockFaker.Lorem().Word()
	originalURL := _mockFaker.Lorem().Word()
	now := time.Now()

	preURL := model.URL{
		ID:          0,
		OriginalURL: originalURL,
		UserID:      1,
	}

	url := model.URL{
		OriginalURL: originalURL,
		Keys:        key,
		ExpiredDate: now.Add(time.Hour * 24),
		UserID:      uint(1),
	}

	user := model.User{
		ID:       1,
		Email:    _mockFaker.Lorem().Word(),
		Password: _mockFaker.Lorem().Word(),
		Key:      apiKey,
	}

	mockUserRepository.On("FindByKey", apiKey).Return(user)
	mockURLRepository.On("FindBy", preURL.OriginalURL, user.ID).Return(preURL)
	mockRedisService.On("Generate").Return(key, nil)
	mockRedisService.On("Map", key, &url).Return(nil)
	mockURLRepository.On("Create", &url).Return(nil)

	shortenerService := ShortenerService{
		redisService: mockRedisService,
		urlRepo:      mockURLRepository,
		userRepo:     mockUserRepository,
	}

	actual, err := shortenerService.ShortenWithAuth(originalURL, key, now, apiKey)
	if err != nil {
		t.Errorf("TestShortenWithAuth is failed")
	}

	if actual.OriginalURL != url.OriginalURL {
		t.Errorf("Actual and expected data is not equal")
	}
}
