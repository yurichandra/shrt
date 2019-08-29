package service

import (
	"crypto/rand"
	"encoding/base64"
	"errors"

	"golang.org/x/crypto/bcrypt"

	"github.com/yurichandra/shrt/model"
	"github.com/yurichandra/shrt/repository"
)

const fixedLength = 10

// AuthService represent service of Auth.
type AuthService struct {
	repo repository.UserRepositoryContract
}

// Authenticate is authenticate or recoginizing user.
func (auth *AuthService) Authenticate(email string, password string) (model.User, error) {
	user := auth.repo.FindByEmail(email)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("Unauthenticated")
	}

	return user, nil
}

// Authorize is authorize access to user.
func (auth *AuthService) Authorize(email string, password string) (model.User, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), 8)
	if err != nil {
		return model.User{}, err
	}

	key, err := auth.generateKey()
	if err != nil {
		return model.User{}, err
	}

	user := model.User{
		Email:    email,
		Password: string(hashedPassword),
		Key:      key,
	}

	err = auth.repo.Create(&user)
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}

func (auth *AuthService) generateKey() (string, error) {
	rb := make([]byte, fixedLength)
	_, err := rand.Read(rb)

	if err != nil {
		return "", err
	}

	key := base64.StdEncoding.EncodeToString(rb)

	return key, nil
}
