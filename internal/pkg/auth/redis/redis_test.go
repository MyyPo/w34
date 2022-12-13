package auth_redis

import (
	"context"
	"testing"
	"time"
)

func TestRedisClient(t *testing.T) {
	ctx := context.Background()
	redisClient := NewRedisClient("localhost:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81")

	const userId int32 = 9999

	t.Run("basic test", func(t *testing.T) {
		value := "hiyaaa"

		err := redisClient.db.Set(ctx, "hellokitty", value, time.Second*1).Err()
		if err != nil {
			t.Errorf("error setting value: %q", err)
		}

		val, err := redisClient.db.Get(ctx, "hellokitty").Result()
		if err != nil {
			t.Errorf("error retrieveing value: %q", err)
		}
		if val != value {
			t.Errorf("got: %s, want %s", val, value)
		}
	})
	t.Run("store token", func(t *testing.T) {
		token := "eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJ0a25fdHlwZSI6ImFjY2VzcyIsInN1YiI6IjQ1IiwiZXhwIjoxNjcwNjE3Njk4LCJuYmYiOjE2NzA2MTcwOTgsImlhdCI6MTY3MDYxNzA5OH0.YcoqC9w91F2pa_BU3fXr41CLrw4AGDhKKHVQF3dfsrG4gH2E8xnIvLHZcODCuih8eT0EPwohF9FMK2JgLMX7-K6PMRo37e2mGnJ5cLZIAYLgKAcIWEq7F6sGkSFu63Ng_AoTwZIzGC-J6uA-AoQDuvEMLe2bybTusQl5ixlWJn3U6THAwhwp3ok78m25jAW4452fzfLUj4LLkbpePM7NLXIFFjrPdyBggS7VKCCaDlALTgzTnAzkeqSrFlpB690fmOPWKHatRhq-PSLV6nsGNMWbBVZwpMqri5gbk-qz2EOeakMZrFNdGk2Lh0WDP1FWpZY8-45SGM4o34NG8KRpR79c6kEFzkEMmrh9VAJw6lDqg0t-F0at95veb4vjf-q60nHQoEV6n5HVx2g3ILyVckQ1yzghIvTsaoHye29p3-rbLgK7FCMXWIVcOq1qUL493d-IXXOzJdfeLkNoDWCOJN6Tvf7Nd5xnXzFQv-qItlKqGpL6vCbnrUw0MWeP9WiXeh0Igew15K3N-hZOXCHLB_riszcRgTXdy3wL4dg1dDVJV-eIAtCqMFKM-D-SLNeQB3d9KaDLKoFcCepuWxwJvrabrcrWbXCN9QkhiB0Ob90MQJfatKC0OWgRoPaKm-kohEVS_MA0ddnWJkHiX0-Jvt4-7w4Cv4nfKFaJY3mm3Lg"

		err := redisClient.StoreRefreshToken(ctx, userId, token)
		if err != nil {
			t.Errorf("error setting value: %q", err)
		}
	})
	// t.Run("get token, delete token then try to retrieve it", func(t *testing.T) {
	// 	_, err := redisClient.GetToken(ctx, "9999")
	// 	if err != nil {
	// 		t.Errorf("failed to retrieve an existing token: %q", err)
	// 	}

	// 	err = redisClient.DeleteRefreshToken(ctx, userId)
	// 	if err != nil {
	// 		t.Errorf("error deletting key: %q", err)
	// 	}

	// 	_, err = redisClient.GetToken(ctx, "9999")
	// 	if err == nil {
	// 		t.Errorf("error deleting: %q", err)
	// 	}
	// })
	// t.Run("try to delete a token that no longer exists", func(t *testing.T) {
	// 	err := redisClient.DeleteRefreshToken(ctx, userId)
	// 	if err != nil {
	// 		t.Logf("error: %q", err)
	// 	}

	// })
}
