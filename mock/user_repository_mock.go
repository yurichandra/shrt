package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yurichandra/shrt/model"
)

// UserRepositoryMock mocks user repository.
type UserRepositoryMock struct {
	mock.Mock
}

// Find mocks find.
func (m *UserRepositoryMock) Find(id uint) model.User {
	args := m.Called(id)

	return args.Get(0).(model.User)
}

// FindByKey mocks FindByKey.
func (m *UserRepositoryMock) FindByKey(key string) model.User {
	args := m.Called(key)

	return args.Get(0).(model.User)
}

// FindByEmail mocks FindByEmail.
func (m *UserRepositoryMock) FindByEmail(email uint) model.User {
	args := m.Called(email)

	return args.Get(0).(model.User)
}

// Create mocks Create.
func (m *UserRepositoryMock) Create(data *model.User) error {
	args := m.Called(&data)

	return args.Error(0)
}
