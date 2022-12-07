package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestJWT(t *testing.T) {
	jwtManager := NewJWTManager(time.Minute*10, time.Hour*48)

	t.Run("Create an access token", func(t *testing.T) {
		userUUID := "16e33d5a-bd6c-4a03-8416-73e89dff2a8a"

		got, err := jwtManager.GenerateJWT(userUUID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}
		testClaims, err := parseTestToken(t, got)
		if err != nil {
			t.Errorf("failed to parse the test token: %q", err)
		}

		sub := testClaims["sub"]
		if sub.(string) != userUUID {
			t.Errorf("issued sub: %s, want id: %s", sub.(string), userUUID)
		}

	})
}

func parseTestToken(t *testing.T, tokenString string) (jwt.MapClaims, error) {
	t.Helper()
	rsaPublicSignature, err := LoadRSAPublicKeyFromDisk("../../../configs/rsa.pub")
	if err != nil {
		return nil, fmt.Errorf("failed to load the signature: %q", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// check the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		// on success return our secret to satisfy the parse function
		// if the signature on token != our returned signature, returns error
		return rsaPublicSignature, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify: %q", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims error: %v", err)
	}

	return claims, nil
}
