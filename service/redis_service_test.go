package service

import (
	"testing"

	"github.com/yurichandra/shrt/model"
)

func TestInitRedisService(t *testing.T) {
	_mockRedis.FlushAll()

	redisService := RedisService{
		client: _mockRedis,
	}

	err := redisService.Init()
	if err != nil {
		t.Errorf("TestRedisService init was failed")
		t.Errorf(err.Error())
	}
}

func TestGenerateRedisService(t *testing.T) {
	_mockRedis.FlushAll()

	redisService := RedisService{
		client: _mockRedis,
	}

	redisService.Init()

	_, err := redisService.Generate()
	if err != nil {
		t.Errorf("TestRedisService generate was failed")
		t.Errorf(err.Error())
	}
}

func TestMapRedisService(t *testing.T) {
	_mockRedis.FlushAll()

	redisService := RedisService{
		client: _mockRedis,
	}

	redisService.Init()
	val, err := redisService.Generate()
	if err != nil {
		t.Errorf(err.Error())
	}

	mockURL := model.URL{
		OriginalURL: _mockFaker.Lorem().Word(),
		Keys:        val,
	}

	err = redisService.Map(val, &mockURL)
	if err != nil {
		t.Errorf("TestRedisService map was failed")
		t.Errorf(err.Error())
	}
}
