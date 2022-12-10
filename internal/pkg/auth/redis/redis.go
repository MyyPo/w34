package auth_redis

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	db        redis.Client
	hasher    hasher.Hasher
	redSync   *redsync.Redsync
	mutexName string
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

	pool := goredis.NewPool(redisDB)
	redSync := redsync.New(pool)

	return &RedisClient{
		db:        *redisDB,
		hasher:    hasher,
		redSync:   redSync,
		mutexName: "cock",
	}
}

func (c RedisClient) StoreRefreshToken(
	ctx context.Context,
	userID int32,
	refreshToken string,
) error {
	strUserID := strconv.FormatInt(int64(userID), 10)

	// mutexName := strUserID
	mutex := c.redSync.NewMutex(c.mutexName)
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer mutex.Unlock()

	err := c.db.Set(ctx, strUserID, refreshToken, time.Hour*48).Err()
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
	// mutexName := userID
	mutex := c.redSync.NewMutex(c.mutexName)
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer mutex.Unlock()

	err := c.db.Set(ctx, userID, refreshToken, time.Hour*48).Err()
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
	// mutexName := strUserID
	mutex := c.redSync.NewMutex(c.mutexName)
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer mutex.Unlock()

	err := c.db.Del(ctx, strUserID).Err()
	if err != nil {
		return err
	}
	return nil
}
func (c RedisClient) DeleteRefreshTokenStringID(
	ctx context.Context,
	userID string,
) error {
	// mutexName := userID
	mutex := c.redSync.NewMutex(c.mutexName)
	if err := mutex.Lock(); err != nil {
		return err
	}
	defer mutex.Unlock()

	err := c.db.Del(ctx, userID).Err()
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
	// mutexName := userID
	mutex := c.redSync.NewMutex(c.mutexName)
	if err := mutex.Lock(); err != nil {
		return "", err
	}
	defer mutex.Unlock()

	token, err := c.db.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}
