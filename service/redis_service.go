package service

import (
	"crypto/rand"
	"encoding/base64"

	"github.com/go-redis/redis"
)

const (
	activeKey = "active_key"
	usedKey   = "used_key"
	mappedKey = "mapped_key"
	keyLength = 4
	keyTotal  = 200
)

// RedisService represent service layer of Redis.
type RedisService struct {
	client *redis.Client
}

// Init initialize random string in active_key redis.
func (service *RedisService) Init() error {
	var temp []string             // Temporary variable to store all encoded string
	keys := make(map[string]bool) // Variable to store is item of array exist or not.

	for i := 0; i < keyTotal; i++ {
		rb := make([]byte, 3)
		_, err := rand.Read(rb)

		if err != nil {
			return err
		}

		key := base64.RawURLEncoding.EncodeToString(rb)
		temp = append(temp, key)
	}

	for _, item := range temp {
		if _, value := keys[item]; !value {
			keys[item] = true
			service.client.RPush(activeKey, item)
		}
	}

	return nil
}
