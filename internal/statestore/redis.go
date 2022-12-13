package statestore

import (
	"context"
	"strconv"
	"time"

	// "github.com/go-redsync/redsync/v4"
	// "github.com/go-redsync/redsync/v4/redis/goredis/v9"

	"github.com/go-redis/redis/v9"
)

type redisClient struct {
	db            redis.Client
	tokenLifetime time.Duration
}

func newRedisClient(
	address, password string,
	tokenLifetime time.Duration,
) *redisClient {
	redisDB := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       0,
	})

	return &redisClient{
		db:            *redisDB,
		tokenLifetime: tokenLifetime,
	}
}

func (c redisClient) StoreRefreshToken(
	ctx context.Context,
	userID int32,
	hashedRefreshToken string,
) error {
	strUserID := strconv.FormatInt(int64(userID), 10)

	err := c.db.Set(ctx, strUserID, hashedRefreshToken, time.Hour*48).Err()
	if err != nil {
		return err
	}

	return nil
}

func (c redisClient) GetToken(
	ctx context.Context,
	userID string,
) (string, error) {

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
