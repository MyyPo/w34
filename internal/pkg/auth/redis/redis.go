package auth_redis

import (
	"context"
	"github.com/bsm/redislock"
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
	backoff := redislock.LinearBackoff(100 * time.Millisecond)
	lock, err := redislock.Obtain(ctx, c.db, strUserID, 100*time.Millisecond, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	err = c.db.Set(ctx, strUserID, refreshToken, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c RedisClient) StoreRefreshTokenStringID(
	ctx context.Context,
	userID string,
	refreshToken string,
) error {
	backoff := redislock.LinearBackoff(100 * time.Millisecond)
	lock, err := redislock.Obtain(ctx, c.db, userID, 100*time.Millisecond, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	err = c.db.Set(ctx, userID, refreshToken, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c RedisClient) DeleteRefreshToken(
	ctx context.Context,
	userID int32,
) error {
	strUserID := strconv.FormatInt(int64(userID), 10)
	backoff := redislock.LinearBackoff(100 * time.Millisecond)
	lock, err := redislock.Obtain(ctx, c.db, strUserID, 100*time.Millisecond, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	err = c.db.Del(ctx, strUserID).Err()
	if err != nil {
		return err
	}
	return nil
}
func (c RedisClient) DeleteRefreshTokenStringID(
	ctx context.Context,
	userID string,
) error {
	backoff := redislock.LinearBackoff(100 * time.Millisecond)
	lock, err := redislock.Obtain(ctx, c.db, userID, 100*time.Millisecond, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return err
	}
	defer lock.Release(ctx)

	err = c.db.Del(ctx, userID).Err()
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}
	return nil
}

func (c RedisClient) GetToken(
	ctx context.Context,
	userID string,
) (string, error) {
	backoff := redislock.LinearBackoff(100 * time.Millisecond)
	lock, err := redislock.Obtain(ctx, c.db, userID, 100*time.Millisecond, &redislock.Options{
		RetryStrategy: backoff,
	})
	if err != nil {
		return "", err
	}
	defer lock.Release(ctx)

	token, err := c.db.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}
