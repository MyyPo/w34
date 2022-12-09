package auth

import (
	"fmt"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestJWT(t *testing.T) {
	jwtManager := NewJWTManager("../../../configs/rsa", "../../../configs/rsa.pub",
		"../../../configs/refresh_rsa", "../../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)

	t.Run("Create an access token, then validate it", func(t *testing.T) {
		var userID int64 = 11

		gotJWT, err := jwtManager.GenerateAccessToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}
		// test the token with a test function independent from jwt manager implementation
		testClaims, err := parseTestToken(t, gotJWT, *jwtManager)
		if err != nil {
			t.Errorf("failed to parse the test token: %q", err)
		}
		sub := testClaims["sub"]
		sub, _ = sub.(int64)
		if sub != userID {
			t.Errorf("issued sub: %v, want id: %v", sub, userID)
		}

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.pathToAccessPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		sub = gotClaims["sub"]
		// sub, _ = sub.(int64)
		if sub != userID {
			t.Errorf("issued sub: %v, want id: %v", sub, userID)
		}
		if tknType := gotClaims["tkn_type"]; tknType != "access" {
			t.Errorf("issued tkn_type: %s, want tkn_type: %s", tknType, "access")
		}

	})
	t.Run("Try to validate a token with invalid signing method", func(t *testing.T) {
		hs256Token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiaWF0IjoxNTE2MjM5MDIyfQ.SflKxwRJSMeKKF2QT4fwpMeJf36POk6yJV_adQssw5c"

		_, err := jwtManager.ValidateJwtExtractClaims(hs256Token, jwtManager.pathToAccessPublicSignature)
		if err == nil {
			t.Errorf("invalid token validated")
		}
	})
	t.Run("Create and validate a refresh token", func(t *testing.T) {
		var userID int64 = 5

		gotJWT, err := jwtManager.GenerateRefreshToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.pathToRefreshPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		if tknType := gotClaims["tkn_type"]; tknType != "refresh" {
			t.Errorf("issued tkn_type: %v, want tkn_type: %s", tknType, "access")
		}
		sub := gotClaims["sub"]
		// sub, ok := sub.(int64)
		// if !ok {
		// 	t.Errorf("failed to convert")
		// }

		if sub != userID {
			t.Errorf("issued sub: %v, want id: %v", sub, userID)
		}

	})
}

func parseTestToken(t *testing.T, tokenString string, jwtManager JWTManager) (jwt.MapClaims, error) {
	t.Helper()
	rsaPublicSignature, err := jwtManager.LoadRSAPublicKeyFromDisk("../../../configs/rsa.pub")
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
