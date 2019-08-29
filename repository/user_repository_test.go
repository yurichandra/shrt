package repository

import (
	"testing"

	"github.com/yurichandra/shrt/db"
	"github.com/yurichandra/shrt/model"
)

func TestFindUserRepository(t *testing.T) {
	db.Reset()

	expectedID := 1
	expected := model.User{
		Email:    _testFaker.Lorem().Word(),
		Password: _testFaker.Lorem().Word(),
		Key:      _testFaker.Lorem().Word(),
	}

	userRepo := UserRepository{
		db: _testDB,
	}

	actual := userRepo.Find(uint(expectedID))

	if actual.ID != expected.ID {
		t.Errorf("Expected and actual data is not equal")
	}
}

func TestFindByKeyUserRepository(t *testing.T) {
	db.Reset()

	expected := model.User{
		Email:    _testFaker.Lorem().Word(),
		Password: _testFaker.Lorem().Word(),
		Key:      _testFaker.Lorem().Word(),
	}

	userRepo := UserRepository{
		db: _testDB,
	}

	userRepo.Create(&expected)

	actual := userRepo.FindByKey(expected.Key)

	if actual.Key != expected.Key {
		t.Errorf("Expected and actual data is not equal")
	}
}

func TestFindByEmailUserRepository(t *testing.T) {
	db.Reset()

	expected := model.User{
		Email:    _testFaker.Lorem().Word(),
		Password: _testFaker.Lorem().Word(),
		Key:      _testFaker.Lorem().Word(),
	}

	userRepo := UserRepository{
		db: _testDB,
	}

	userRepo.Create(&expected)

	actual := userRepo.FindByEmail(expected.Email)

	if actual.Email != expected.Email {
		t.Errorf("Expected and actual data is not equal")
	}
}

func TestCreateUserRepository(t *testing.T) {
	db.Reset()

	expected := model.User{
		Email:    _testFaker.Lorem().Word(),
		Password: _testFaker.Lorem().Word(),
		Key:      _testFaker.Lorem().Word(),
	}

	userRepo := UserRepository{
		db: _testDB,
	}

	err := userRepo.Create(&expected)
	if err != nil {
		t.Errorf("Fail to create data on UserRepository")
	}
}
