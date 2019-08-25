package mock

import "github.com/stretchr/testify/mock"

// RedisServiceMock represent mocking of RedisService.
type RedisServiceMock struct {
	mock.Mock
}

// HGet is mocking representation of HGet() on CacheService.
func (m *RedisServiceMock) HGet(hash string, field string) (string, error) {
	args := m.Called(hash, field)

	return args.String(0), args.Error(1)
}

// HSet is mocking representation of HSet() on CacheService.
func (m *RedisServiceMock) HSet(hash string, field string, value string) error {
	args := m.Called(hash, field, value)

	return args.Error(0)
}
