package auth

import (
	"crypto/rsa"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

type JWTManager struct {
	pathToAccessPrivateSignature  string
	pathToAccessPublicSignature   string
	pathToRefreshPrivateSignature string
	pathToRefreshPublicSignature  string
	accessTokenDuraion            time.Duration
	refreshTokenDuration          time.Duration
}

func NewJWTManager(
	pathToAccessPrivateSignature, pathToAccessPublicSignature string,
	pathToRefreshPrivateSignature, pathToRefreshPublicSignature string,
	accessTokenDuration, refreshTokenDuration time.Duration,
) *JWTManager {
	return &JWTManager{
		pathToAccessPrivateSignature:  pathToAccessPrivateSignature,
		pathToAccessPublicSignature:   pathToAccessPublicSignature,
		pathToRefreshPrivateSignature: pathToRefreshPrivateSignature,
		pathToRefreshPublicSignature:  pathToRefreshPublicSignature,
		accessTokenDuraion:            accessTokenDuration,
		refreshTokenDuration:          refreshTokenDuration,
	}
}

func (m JWTManager) GenerateAccessToken(userUUID uuid.UUID) (string, error) {
	rsaPrivateAccessSignature, err := m.LoadRSAPrivateKeyFromDisk(m.pathToAccessPrivateSignature)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["exp"] = now.Add(m.accessTokenDuraion).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["sub"] = userUUID
	claims["tkn_type"] = "access"

	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateAccessSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to access token: %w", err)
	}

	return accessToken, nil
}

func (m JWTManager) GenerateRefreshToken(userUUID uuid.UUID) (string, error) {
	rsaPrivateRefreshSignature, err := m.LoadRSAPrivateKeyFromDisk(m.pathToRefreshPrivateSignature)
	if err != nil {
		return "", err
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)

	claims["exp"] = now.Add(m.refreshTokenDuration).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()
	claims["sub"] = userUUID
	claims["tkn_type"] = "refresh"

	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(rsaPrivateRefreshSignature)

	if err != nil {
		return "", fmt.Errorf("failed to add claims to refresh token: %w", err)
	}

	return refreshToken, nil
}

func (m JWTManager) ValidateJwtExtractClaims(jwtTokenString, publicSignaturePath string) (jwt.MapClaims, error) {
	rsaPublicSignature, err := m.LoadRSAPublicKeyFromDisk(publicSignaturePath)
	if err != nil {
		return nil, fmt.Errorf("failed to load the signature: %q", err)
	}

	jwtToken, err := jwt.Parse(jwtTokenString, func(token *jwt.Token) (interface{}, error) {
		// check if the signing algorithm is correct
		if _, ok := token.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// check the signature of the token
		return rsaPublicSignature, nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to verify: %q", err)
	}

	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("claims error: %q", err)
	}

	return claims, nil
}

func (m JWTManager) LoadRSAPrivateKeyFromDisk(location string) (*rsa.PrivateKey, error) {
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

func (m JWTManager) LoadRSAPublicKeyFromDisk(location string) (*rsa.PublicKey, error) {
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
