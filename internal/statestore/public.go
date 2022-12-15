package statestore

import (
	"context"
	"time"
)

type Service interface {
	// Refresh Token whitelist handling

	// Whitelist a new token for a given user
	StoreRefreshToken(ctx context.Context, userID int32, hashedRefreshToken string) error

	// Retrieve a whitelisted token for a given user
	GetToken(ctx context.Context, userID string) (string, error)

	// Alt version of StoreRefreshToken
	StoreRefreshTokenStringID(ctx context.Context, userID string, hashedRefreshToken string) error
}

func New() Service {
	// TODO: Put actual config init here
	redis := newRedisClient("localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", 48*time.Hour)

	return redis
}
