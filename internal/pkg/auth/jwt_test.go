package auth

import (
	"testing"
	"time"
)

func TestJWT(t *testing.T) {
	jwtManager := NewJWTManager("../../../configs/rsa", "../../../configs/rsa.pub",
		"../../../configs/refresh_rsa", "../../../configs/refresh_rsa.pub",
		time.Minute*10, time.Hour*48)

	t.Run("Create an access token, then validate it", func(t *testing.T) {
		var userID int32 = 11

		gotJWT, err := jwtManager.GenerateAccessToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}

		strUserID := "11"

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.pathToAccessPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		sub := gotClaims.Subject
		if sub != strUserID {
			t.Errorf("issued sub: %v, want id: %v", sub, strUserID)
		}
		if tknType := gotClaims.TknType; tknType != "access" {
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
		var userID int32 = 5

		gotJWT, err := jwtManager.GenerateRefreshToken(userID)
		if err != nil {
			t.Errorf("jwt error: %q", err)
		}

		gotClaims, err := jwtManager.ValidateJwtExtractClaims(gotJWT, jwtManager.pathToRefreshPublicSignature)
		if err != nil {
			t.Errorf("failed to validate valid token: %q", err)
		}

		if tknType := gotClaims.TknType; tknType != "refresh" {
			t.Errorf("issued tkn_type: %v, want tkn_type: %s", tknType, "refresh")
		}
		sub := gotClaims.Subject

		strUserID := "5"

		if sub != strUserID {
			t.Errorf("issued sub: %s, want id: %s", sub, strUserID)
		}

	})
}
