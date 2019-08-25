package service

import (
	"encoding/json"
	"testing"

	"github.com/yurichandra/shrt/mock"
	"github.com/yurichandra/shrt/model"
)

func TestURLServiceGet(t *testing.T) {
	mockRedisService := new(mock.RedisServiceMock)
	mockURLRepository := new(mock.URLRepositoryMock)

	mockURLRepository.On("Get").Return(make([]model.URL, 0))

	mockURLService := URLService{
		repo:  mockURLRepository,
		cache: mockRedisService,
	}

	expected := mockURLService.Get()

	if len(expected) != 0 {
		t.Errorf("Expected data should be 0")
	}
}

func TestURLServiceFind(t *testing.T) {
	mockRedisService := new(mock.RedisServiceMock)
	mockURLRepository := new(mock.URLRepositoryMock)

	expected := model.URL{
		OriginalURL: _mockFaker.Lorem().Word(),
		ShortURL:    _mockFaker.Lorem().Word(),
	}

	mockURLRepository.On("Find", expected.ID).Return(expected)

	mockURLService := URLService{
		repo:  mockURLRepository,
		cache: mockRedisService,
	}

	actual := mockURLService.Find(uint(expected.ID))
	if actual.ID != expected.ID {
		t.Errorf("Expected ID and actual ID is not equal")
	}
}

func TestURLServiceCreate(t *testing.T) {
	mockRedisService := new(mock.RedisServiceMock)
	mockURLRepository := new(mock.URLRepositoryMock)

	mockKey := "urls"
	mockProceededURL := "d41d8c"

	expected := model.URL{
		OriginalURL: _mockFaker.Lorem().Word(),
		ShortURL:    mockProceededURL,
	}

	urlCasted, _ := json.Marshal(expected)

	mockRedisService.On("HGet", mockKey, mockProceededURL).Return("", nil)
	mockRedisService.On("HSet", mockKey, mockProceededURL, string(urlCasted)).Return(nil)
	mockURLRepository.On("Create", &expected).Return(nil)

	mockURLService := URLService{
		repo:  mockURLRepository,
		cache: mockRedisService,
	}

	actual, err := mockURLService.Create(expected.OriginalURL)
	if err != nil {
		t.Errorf("Error happened during call Create() method on URLService")
	}

	if actual.OriginalURL != expected.OriginalURL {
		t.Errorf("Expected and actual data is not equal")
	}
}
