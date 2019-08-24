package service

import "testing"

func TestHGet(t *testing.T) {
	hashKey := _mockFaker.Lorem().Word()
	_mockField := _mockFaker.Lorem().Word()
	_mockValue := _mockFaker.Lorem().Word()

	_mockRedis.HSet(hashKey, _mockField, _mockValue)

	mockRedisService := &RedisService{
		client: _mockRedis,
	}

	expected, err := mockRedisService.HGet(hashKey, _mockField)
	if err != nil {
		t.Errorf("Method HGet() was failed during testing")
	}

	if expected != _mockValue {
		t.Errorf("Expected and Actual result is not equal")
	}
}

func TestHSet(t *testing.T) {
	hashKey := _mockFaker.Lorem().Word()
	_mockField := _mockFaker.Lorem().Word()
	_mockValue := _mockFaker.Lorem().Word()

	mockRedisService := &RedisService{
		client: _mockRedis,
	}

	err := mockRedisService.HSet(
		hashKey,
		_mockField,
		_mockValue,
	)
	if err != nil {
		t.Errorf("Method HSet() was failed during testing")
	}
}
