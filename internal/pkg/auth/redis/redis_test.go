package auth_redis

import (
	"context"
	"testing"
	"time"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/hasher"
)

func TestRedisClient(t *testing.T) {
	ctx := context.Background()

	hasher := hasher.NewHasher()
	redis := NewRedisClient("localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", *hasher)

	value := "hiyaaa"

	err := redis.db.Set(ctx, "hellokitty", value, time.Second*1).Err()
	if err != nil {
		t.Errorf("error setting value: %q", err)
	}

	val, err := redis.db.Get(ctx, "hellokitty").Result()
	if err != nil {
		t.Errorf("error retrieveing value: %q", err)
	}
	if val != value {
		t.Errorf("got: %s, want %s", val, value)
	}
}
