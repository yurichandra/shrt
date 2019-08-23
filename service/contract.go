package service

// RedisServiceContract represent contract of
// redis service.
type RedisServiceContract interface {
	HSet(hash string, field string, value string) error
	HGet(hash string, field string) (string, error)
}
