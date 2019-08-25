package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yurichandra/shrt/model"
)

// URLRepositoryMock represent mock of URLRepository.
type URLRepositoryMock struct {
	mock.Mock
}

// Get is mocking representation of Get() on URLRepository.
func (m *URLRepositoryMock) Get() []model.URL {
	args := m.Called()

	return args.Get(0).([]model.URL)
}

// Find is mocking representation of Find() on URLRepository.
func (m *URLRepositoryMock) Find(id uint) model.URL {
	args := m.Called(id)

	return args.Get(0).(model.URL)
}

// Create is mocking representation of Create() on URLRepository.
func (m *URLRepositoryMock) Create(data *model.URL) error {
	args := m.Called(data)

	return args.Error(0)
}

// Update is mocking representation of Update() on URLRepository.
func (m *URLRepositoryMock) Update(data *model.URL) error {
	args := m.Called(data)

	return args.Error(0)
}

// Delete is mocking representation of Delete() on URLRepository.
func (m *URLRepositoryMock) Delete(id uint) error {
	args := m.Called(id)

	return args.Error(0)
}
