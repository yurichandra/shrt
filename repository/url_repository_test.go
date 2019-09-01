package repository

import (
	"testing"

	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/model"
)

func TestGetURLRepository(t *testing.T) {
	db.Reset()

	for i := 0; i < 5; i++ {
		db.Get().Create(&model.URL{
			OriginalURL: _testFaker.Lorem().Word(),
			Keys:        _testFaker.Lorem().Word(),
			UserID:      uint(1),
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

func TestFindURLRepository(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
		UserID:      uint(1),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	testURLRepo.Create(mockedURL)

	expected := testURLRepo.Find(mockedURL.ID)
	if expected.OriginalURL != mockedURL.OriginalURL {
		t.Errorf("Expected and actual data not matched")
	}

	if expected.Keys != mockedURL.Keys {
		t.Errorf("Expected and actual data not matched")
	}
}

func TestFindByURLRepository(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
		UserID:      uint(1),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	testURLRepo.Create(mockedURL)

	expected := testURLRepo.FindBy(mockedURL.OriginalURL, mockedURL.UserID)
	if expected.OriginalURL != mockedURL.OriginalURL {
		t.Errorf("Expected and actual data not matched")
	}

	if expected.UserID != mockedURL.UserID {
		t.Errorf("Expected and actual data not matched")
	}

	if expected.Keys != mockedURL.Keys {
		t.Errorf("Expected and actual data not matched")
	}
}

func TestCreateURLRepository(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
		UserID:      uint(1),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	expected := testURLRepo.Create(mockedURL)
	if expected != nil {
		t.Errorf("Error happened during calling Create() method")
	}
}

func TestUpdateURLRepository(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
		UserID:      uint(1),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	testURLRepo.Create(mockedURL)

	updatedMockURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
	}

	expected := testURLRepo.Update(updatedMockURL)
	if expected != nil {
		t.Errorf("Error happened during calling Update() method")
	}
}

func TestDeleteURLRepository(t *testing.T) {
	db.Reset()

	mockedURL := &model.URL{
		OriginalURL: _testFaker.Lorem().Word(),
		Keys:        _testFaker.Lorem().Word(),
		UserID:      uint(1),
	}

	testURLRepo := &URLRepository{
		db: _testDB,
	}

	testURLRepo.Create(mockedURL)

	expected := testURLRepo.Delete(mockedURL.ID)
	if expected != nil {
		t.Errorf("Error happened during calling Delete() method")
	}
}
