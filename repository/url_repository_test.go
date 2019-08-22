package repository

import (
	"testing"

	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/model"
)

func TestGet(t *testing.T) {
	db.Reset()

	for i := 0; i < 5; i++ {
		db.Get().Create(&model.URL{
			OriginalURL: _testFaker.Lorem().Word(),
			ShortURL:    _testFaker.Lorem().Word(),
		})
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	urls := testURLRepo.Get()

	if len(urls) != 5 {
		t.Errorf("Get() method fail to get all urls in repository")
	}
}

func TestFind(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		ShortURL:    _testFaker.Lorem().Word(),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	testURLRepo.Create(mockedURL)

	expected := testURLRepo.Find(mockedURL.ID)
	if expected.OriginalURL != mockedURL.OriginalURL {
		t.Errorf("Expected and actual data not matched")
	}

	if expected.ShortURL != mockedURL.ShortURL {
		t.Errorf("Expected and actual data not matched")
	}
}

func TestCreate(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		ShortURL:    _testFaker.Lorem().Word(),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	expected := testURLRepo.Create(mockedURL)
	if expected != nil {
		t.Errorf("Something happened during calling Create() method")
	}
}
