package interceptor

import (
	"context"
	"testing"
	"time"

	"github.com/MyyPo/w34.Go/internal/jwt"
	"github.com/MyyPo/w34.Go/internal/statestore"
	"google.golang.org/grpc/metadata"
)

func TestInterceptor(t *testing.T) {
	jwtManager, err := jwt.NewJWTManager("../../../../configs/rsa", "../../../../configs/rsa.pub",
		"../../../../configs/refresh_rsa", "../../../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)
	if err != nil {
		t.Errorf("failed to initialize jwtManager: %v", err)
	}

	redisClient := statestore.New()

	roles := map[string][]string{
		"201": {"admin"},
	}

	interceptor := NewAuthInterceptor(*jwtManager, redisClient, roles)

	t.Run("Basic test authorize", func(t *testing.T) {
		_, err := interceptor.authorize(context.Background(), "200")
		if err != nil {
			t.Errorf("authorize error: %v", err)
		}

	})
	t.Run("Try to get protected route without metadata", func(t *testing.T) {
		_, err := interceptor.authorize(context.Background(), "201")
		if err == nil {
			t.Errorf("accessed admin route")
		}
	})
	t.Run("Provide invalid metadata, no access token", func(t *testing.T) {
		md := metadata.New(map[string]string{
			"invalid_key": "invalid_value",
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		_, err = interceptor.authorize(contextWithMD, "201")
		if err == nil {
			t.Errorf("accessed admin route")
		}
	})
	t.Run("Provide empty metadata", func(t *testing.T) {
		md := metadata.New(map[string]string{})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		_, err = interceptor.authorize(contextWithMD, "201")
		if err == nil {
			t.Errorf("accessed admin route")
		}
	})
	t.Run("Provide valid access token", func(t *testing.T) {
		accessToken, _ := jwtManager.GenerateAccessToken(42)

		md := metadata.New(map[string]string{
			"access_token": accessToken,
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		newCtx, err := interceptor.authorize(contextWithMD, "201")
		if err != nil {
			t.Errorf("failed to authorize: %v", err)
		}

		t.Logf("New context: %s", newCtx.Value(userIdKey))
	})
	t.Run("Provide no access token, but valid refresh token", func(t *testing.T) {
		// accessToken, err := jwtManager.GenerateAccessToken(42)
		refreshToken, _ := jwtManager.GenerateRefreshToken(42)
		redisClient.StoreRefreshToken(context.Background(), 42, refreshToken)

		md := metadata.New(map[string]string{
			"access_token":  "",
			"refresh_token": refreshToken,
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		newCtx, err := interceptor.authorize(contextWithMD, "201")
		if err != nil {
			t.Errorf("failed to authorize: %v", err)
		}

		t.Logf("New context: %s", newCtx.Value(userIdKey))
	})
	time.Sleep(1 * time.Second)
	t.Run("Provide invalid tokens", func(t *testing.T) {
		refreshToken, _ := jwtManager.GenerateRefreshToken(42)

		md := metadata.New(map[string]string{
			"access_token":  "",
			"refresh_token": refreshToken,
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		_, err := interceptor.authorize(contextWithMD, "201")
		if err == nil {
			t.Errorf("authorized the invalid request")
		}
	})
	t.Run("Try to pass fake user id to context with access token", func(t *testing.T) {
		accessToken, _ := jwtManager.GenerateAccessToken(42)

		md := metadata.New(map[string]string{
			"access_token":  accessToken,
			"refresh_token": "",
			"user_id":       "999",
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		newCtx, err := interceptor.authorize(contextWithMD, "201")
		if err != nil {
			t.Errorf("unexpected error with valid access token: %v", err)
		}
		if newCtx.Value(userIdKey) != "42" {
			t.Errorf("unexpected change of user id in md")
		}

	})
	t.Run("Try to pass fake user id to context with refresh token", func(t *testing.T) {
		refreshToken, _ := jwtManager.GenerateRefreshToken(42)
		redisClient.StoreRefreshToken(context.Background(), 42, refreshToken)

		md := metadata.New(map[string]string{
			"access_token":  "",
			"refresh_token": refreshToken,
			"user_id":       "999",
		})

		contextWithMD := metadata.NewIncomingContext(context.Background(), md)

		newCtx, err := interceptor.authorize(contextWithMD, "201")
		if err != nil {
			t.Errorf("unexpected error with valid refresh token: %v", err)
		}
		if newCtx.Value(userIdKey) != "42" {
			t.Errorf("unexpected change of user id in md")
		}

	})
}
