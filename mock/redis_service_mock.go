package mock

import (
	"github.com/stretchr/testify/mock"
	"github.com/yurichandra/shrt/model"
)

// RedisServiceMock represent mocking of RedisService.
type RedisServiceMock struct {
	mock.Mock
}

// Init mocks init.
func (m *RedisServiceMock) Init() error {
	args := m.Called()

	return args.Error(0)
}

// Generate mocks generate.
func (m *RedisServiceMock) Generate() (string, error) {
	args := m.Called()

	return args.String(0), args.Error(1)
}

// Map mocks map.
func (m *RedisServiceMock) Map(key string, url *model.URL) error {
	args := m.Called(key, url)

	return args.Error(0)
}

// Find mocks find.
func (m *RedisServiceMock) Find(key string) (string, error) {
	args := m.Called(key)

	return args.String(0), args.Error(1)
}
