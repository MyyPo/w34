package auth_redis

import "github.com/go-redis/redis/v9"

type RedisClient struct {
	db redis.Client
}

func NewRedisClient(address, password string) *RedisClient {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &RedisClient{
		db: *redisDB,
	}
}
