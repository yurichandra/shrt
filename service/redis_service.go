package service

import (
	"errors"

	"github.com/go-redis/redis"
)

// RedisService represent service layer of Redis.
type RedisService struct {
	client *redis.Client
}

// HGet fetch data in Redis cluster based on hash and field.
func (r *RedisService) HGet(hash string, field string) (string, error) {
	value, err := r.client.HGet(hash, field).Result()
	if err != nil {
		return "", errors.New("Hash or field not exist")
	}

	return value, nil
}

// HSet save field data in redis cluster based on hash and field
func (r *RedisService) HSet(hash string, field string, value string) error {
	err := r.client.HSet(hash, field, value).Err()
	if err != nil {
		return errors.New("HSet was failed")
	}

	return nil
}
