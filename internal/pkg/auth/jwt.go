package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
	accessTokenDuraion   time.Duration
	refreshTokenDuration time.Duration
}

func NewJWTManager(
	accessTokenDuration, refreshTokenDuration time.Duration,
) *JWTManager {
	return &JWTManager{
		accessTokenDuraion:   accessTokenDuration,
		refreshTokenDuration: refreshTokenDuration,
	}
}

func (m JWTManager) GenerateJWT(userUUID string) (string, error) {
	rsaPrivateSignature, err := LoadRSAPrivateKeyFromDisk("../../../configs/rsa")
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(m.accessTokenDuraion).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["sub"] = userUUID

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to token: %w", err)
	}

	return token, nil
}

func LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
	keyData, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return privateKey, nil
}

func LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
	keyData, err := os.ReadFile(location)
	if err != nil {
		return nil, err
	}
	publicKey, err := jwt.ParseRSAPublicKeyFromPEM(keyData)
	if err != nil {
		return nil, err
	}
	return publicKey, nil
}
