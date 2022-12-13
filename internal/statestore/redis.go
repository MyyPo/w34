package statestore

import (
	"context"
	"strconv"
	"time"

	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"github.com/go-redis/redis/v9"
)

type RedisClient struct {
	db            redis.Client
	redSync       *redsync.Redsync
	mutexName     string
	tokenLifetime time.Duration
}

func NewRedisClient(
	address, password string,
	tokenLifetime time.Duration,
) *RedisClient {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	pool := goredis.NewPool(redisDB)
	redSync := redsync.New(pool)

	return &RedisClient{
		db:            *redisDB,
		redSync:       redSync,
		mutexName:     "global",
		tokenLifetime: tokenLifetime,
	}
}

func (c RedisClient) StoreRefreshToken(
	ctx context.Context,
	userID int32,
	hashedRefreshToken string,
) error {
	strUserID := strconv.FormatInt(int64(userID), 10)

	mutex := c.redSync.NewMutex(c.mutexName, redsync.WithExpiry(1*time.Minute))
	if err := mutex.LockContext(ctx); err != nil {
		return err
	}
	defer mutex.UnlockContext(ctx)

	err := c.db.Set(ctx, strUserID, hashedRefreshToken, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c RedisClient) GetToken(
	ctx context.Context,
	userID string,
) (string, error) {
	mutex := c.redSync.NewMutex(c.mutexName, redsync.WithExpiry(1*time.Minute))

	if err := mutex.LockContext(ctx); err != nil {
		return "", err
	}
	defer mutex.UnlockContext(ctx)

	token, err := c.db.Get(ctx, userID).Result()
	if err != nil {
		return "", err
	}

	return token, nil
}

// func handleConnectionClose(conn *redis.Conn) {
// 	// err := (*conn).Close()
// 	(*conn).Close()
// }

// func (c RedisClient) StoreRefreshTokenStringID(
// 	ctx context.Context,
// 	userID string,
// 	hashedRefreshToken string,
// ) error {
// 	// mutexName := userID
// 	mutex := c.redSync.NewMutex(c.mutexName)
// 	if err := mutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer mutex.Unlock()

// 	err := c.db.Set(ctx, userID, hashedRefreshToken, time.Hour*48).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
// func (c RedisClient) DeleteRefreshToken(
// 	ctx context.Context,
// 	userID int32,
// ) error {
// 	strUserID := strconv.FormatInt(int64(userID), 10)
// 	// mutexName := strUserID
// 	mutex := c.redSync.NewMutex(c.mutexName)
// 	if err := mutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer mutex.Unlock()

// 	err := c.db.Del(ctx, strUserID).Err()
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
// func (c RedisClient) DeleteRefreshTokenStringID(
// 	ctx context.Context,
// 	userID string,
// ) error {
// 	// mutexName := userID
// 	mutex := c.redSync.NewMutex(c.mutexName)
// 	if err := mutex.Lock(); err != nil {
// 		return err
// 	}
// 	defer mutex.Unlock()

// 	err := c.db.Del(ctx, userID).Err()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
