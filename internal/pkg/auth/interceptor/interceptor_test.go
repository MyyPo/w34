package interceptor

import (
	"context"
	"testing"
	"time"

	"github.com/MyyPo/w34.Go/internal/pkg/auth/jwt"
	"github.com/MyyPo/w34.Go/internal/statestore"
)

func TestInterceptor(t *testing.T) {
	jwtManager, err := jwt.NewJWTManager("../../../../configs/rsa", "../../../../configs/rsa.pub",
		"../../../../configs/refresh_rsa", "../../../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)
	if err != nil {
		t.Errorf("failed to initialize jwtManager: %v", err)
	}

	redisClient := statestore.NewRedisClient("host.docker.internal:6379", "eYVX7EwVmmxKPCDmwMtyKVge8oLd2t81", 48*time.Hour)

	roles := map[string][]string{
		"201": {"admin"},
	}

	interceptor := NewAuthInterceptor(*jwtManager, *redisClient, roles)

	t.Run("Basic test authorize", func(t *testing.T) {
		err := interceptor.authorize(context.Background(), "200")
		if err != nil {
			t.Errorf("authorize error: %v", err)
		}
	})
	// t.Run("Try to get protected route without metadata", func(t *testing.T) {
	// 	err := interceptor.authorize(context.Background(), "201")
	// 	if err == nil {
	// 		t.Errorf("accessed admin route")
	// 	}
	// })
	// t.Run("Provide invalid metadata, no access token", func(t *testing.T) {
	// 	// md := metadata.New(map[string]string{
	// 	// 	"invalid_key": "invalid_value",
	// 	// })
	// 	contextWithMD := metadata.AppendToOutgoingContext(context.Background(), "key", "val")
	// 	fmt.Println(metadata.FromOutgoingContext(contextWithMD))

	// 	err := interceptor.authorize(contextWithMD, "201")
	// 	if err == nil {
	// 		t.Errorf("accessed admin route")
	// 	}
	// 	t.Logf("err: %v", err)
	// })
}
