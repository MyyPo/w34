package auth_redis

import (
	"context"
	"strconv"
	"time"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	db     redis.Client
	hasher hasher.Hasher
}

func NewRedisClient(
	address, password string,
	hasher hasher.Hasher,
) *RedisClient {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &RedisClient{
		db:     *redisDB,
		hasher: hasher,
	}
}

func (c RedisClient) StoreRefreshToken(
	ctx context.Context,
	userID int32,
	refreshToken string,
) error {
	strUserID := strconv.FormatInt(int64(userID), 10)
	err := c.db.Set(ctx, strUserID, refreshToken, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}
